package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/mutezebra/subject-review/biz/model/subject"
	"github.com/mutezebra/subject-review/pkg/client"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/pack"
	"github.com/mutezebra/subject-review/pkg/utils"
)

type SubjectDB struct {
	*gorm.DB
}

func NewSubjectDB() *SubjectDB {
	return &SubjectDB{DB: client.GetMysqlDB()}
}

var (
	SubjectTable = constants.SubjectTable
	UWSTable     = constants.SubjectWithUserTable
	RemindTable  = constants.RemindTable
)

func (db *SubjectDB) GetSubjects(ctx context.Context, sbType int16, pages, size int64) ([]*subject.BaseSubject, error) {
	offset := int((pages - 1) * size)

	result := make([]*subject.BaseSubject, 0)
	if err := db.WithContext(ctx).Table(SubjectTable).
		Offset(offset).Limit(int(size)).Where("subject_type=?", sbType).
		Select([]string{"id", "name", "answer", "subject_type"}).Find(&result).Error; err != nil {
		return nil, errors.WithMessage(err, "failed when get subjects,error: ")
	}

	return result, nil
}

func (db *SubjectDB) GetSubject(ctx context.Context, sbID int64) (*subject.Subject, error) {
	var s subject.Subject
	if err := db.WithContext(ctx).Table(SubjectTable).Where("id=?", sbID).First(&s).Error; err != nil {
		return nil, errors.WithMessage(err, "failed when get subject,error: ")
	}
	return &s, nil
}

func (db *SubjectDB) GetReviewRecords(ctx context.Context, uid int64, pages, size int64) ([]*subject.SubjectRecord, int64, error) {
	/*
		SELECT s.id, s.name, s.answer, s.subject_type, uws.phase, uws.learn_times, uws.last_review_at, r.remind
		    FROM user_with_subject AS uws
		        JOIN subject AS s ON s.id = uws.subject_id
		        JOIN remind AS r ON r.subject_id = uws.subject_id
		    WHERE r.user_id = ? LIMIT ? OFFSET ?
	*/
	sql, err := db.DB.DB()
	if err != nil {
		return nil, 0, errors.WithMessage(err, "get db from gorm.DB failed")
	}

	// Query to get total count
	countQuery := "SELECT COUNT(*) FROM user_with_subject AS uws JOIN subject AS s ON s.id = uws.subject_id JOIN remind AS r ON r.subject_id = uws.subject_id WHERE r.user_id = ?"
	var total int64
	err = sql.QueryRowContext(ctx, countQuery, uid).Scan(&total)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed to get total count of records")
	}

	// Query to get paginated results
	query := "SELECT s.id, s.name, s.answer, s.subject_type, uws.phase, uws.learn_times, uws.last_review_at, r.remind FROM user_with_subject AS uws JOIN subject AS s ON s.id = uws.subject_id JOIN remind AS r ON r.subject_id = uws.subject_id WHERE r.user_id = ? LIMIT ? OFFSET ?"
	rows, err := sql.QueryContext(ctx, query, uid, size, (pages-1)*size)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed when query answer records")
	}
	defer func() {
		pack.LogError(rows.Close())
	}()

	records := make([]*subject.SubjectRecord, 0)
	for rows.Next() {
		record := &subject.SubjectRecord{}
		if err = rows.Scan(&record.ID, &record.Name, &record.Answer, &record.SubjectType, &record.Phase,
			&record.LearnTimes, &record.LastReviewAt, &record.NextReviewAt); err != nil {
			return nil, 0, errors.WithMessage(err, "failed when scan variable to SubjectRecord")
		}
		records = append(records, record)
	}

	return records, total, nil
}

func (db *SubjectDB) GetNeedReviewedSubjects(ctx context.Context, uid int64, pages, size int64, endTime int64) ([]*subject.SubjectRecord, int64, error) {
	sql, err := db.DB.DB()
	if err != nil {
		return nil, 0, errors.WithMessage(err, "get db from gorm.DB failed")
	}

	// Query to get total count
	countQuery := "SELECT COUNT(*) FROM user_with_subject AS uws JOIN subject AS s ON s.id = uws.subject_id JOIN remind AS r ON r.subject_id = uws.subject_id WHERE (r.user_id = ? AND r.remind < ?)"
	var total int64
	err = sql.QueryRowContext(ctx, countQuery, uid, endTime).Scan(&total)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed to get total count of records")
	}

	// Query to get paginated results
	query := "SELECT s.id, s.name, s.answer, s.subject_type, uws.phase, uws.learn_times, uws.last_review_at, r.remind FROM user_with_subject AS uws JOIN subject AS s ON s.id = uws.subject_id JOIN remind AS r ON r.subject_id = uws.subject_id WHERE (r.user_id = ? AND r.remind < ?) LIMIT ? OFFSET ?"

	rows, err := sql.QueryContext(ctx, query, uid, endTime, size, (pages-1)*size)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed when query answer records")
	}
	defer func() {
		pack.LogError(rows.Close())
	}()

	records := make([]*subject.SubjectRecord, 0)
	for rows.Next() {
		record := &subject.SubjectRecord{}
		if err = rows.Scan(&record.ID, &record.Name, &record.Answer, &record.SubjectType, &record.Phase,
			&record.LearnTimes, &record.LastReviewAt, &record.NextReviewAt); err != nil {
			return nil, 0, errors.WithMessage(err, "failed when scan variable to SubjectRecord")
		}
		records = append(records, record)
	}

	return records, total, nil
}

func (db *SubjectDB) CreateSubject(ctx context.Context, subject *subject.Subject) error {
	if err := db.WithContext(ctx).Table(SubjectTable).Create(&subject).Error; err != nil {
		return errors.WithMessage(err, "create subject failed")
	}
	return nil
}

type times struct {
	Phase      int16 `json:"phase" gorm:"phase"`
	LearnTimes int16 `json:"learn_times" gorm:"learn_times"`
}

func (db *SubjectDB) AddForgetSubject(ctx context.Context, userID, sbID int64, sbType int16) error {
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count times
		now := time.Now().Unix()
		// 先尝试从`user_with_subject`表中查找相应记录, 看用户先前是否有学习过这个subject
		if err := tx.Table(UWSTable).Select("phase,learn_times").Where("user_id=? AND subject_id=?", userID, sbID).Find(&count).Error; err != nil {
			return errors.WithMessage(err, fmt.Sprintf("find user_id=%d and sbid=%d in `user_with_subject` failed", userID, sbID))
		}
		if count.LearnTimes == 0 {
			// 如果没学习过，那就创建一条记录
			if err := tx.Table(UWSTable).Create(subject.UserWithSubject{
				UserID:       utils.Ptr(userID),
				SubjectID:    utils.Ptr(sbID),
				SubjectType:  utils.Ptr(sbType),
				Phase:        utils.Ptr(int16(0)),
				LearnTimes:   utils.Ptr(int16(1)),
				LastReviewAt: utils.Ptr(now),
			}).Error; err != nil {
				return errors.WithMessage(err, "failed when create user_with_subject record,error: ")
			}

			// 没学习过要创建信息
			if err := tx.Table(RemindTable).Create(&subject.Remind{
				UserID:    utils.Ptr(userID),
				SubjectID: utils.Ptr(sbID),
				Remind:    utils.Ptr(utils.GetRemind(0)),
			}).Error; err != nil {
				return errors.WithMessage(err, "failed create remind record,error: ")
			}

			return nil
		}

		// 学习过
		if count.Phase > 0 { // 如果phase大于0，那就倒退一个阶段
			count.Phase -= 1
		}

		if err := tx.Table(UWSTable).Where("user_id=? AND subject_id=?", userID, sbID).
			UpdateColumns(map[string]interface{}{
				"phase":          count.Phase,
				"learn_times":    count.LearnTimes + 1,
				"last_review_at": now,
			}).Error; err != nil {
			return errors.WithMessage(err, "failed when update phase,error: ")
		}

		if err := tx.Table(RemindTable).Where("user_id=? AND subject_id=?", userID, sbID).
			UpdateColumn("remind", utils.GetRemind(count.Phase)).Error; err != nil {
			return errors.WithMessage(err, "failed update remind,error: ")
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *SubjectDB) AddSuccessSubject(ctx context.Context, userID, sbID int64, sbType int16) error {
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count times
		now := time.Now().Unix()
		// 先尝试从`user_with_subject`表中查找相应记录, 看用户先前是否有学习过这个subject
		if err := tx.Table(UWSTable).Select("phase,learn_times").Where("user_id=? AND subject_id=?", userID, sbID).Find(&count).Error; err != nil {
			return errors.WithMessage(err, fmt.Sprintf("find user_id=%d and sbid=%d in `user_with_subject` failed", userID, sbID))
		}

		if count.LearnTimes == 0 {
			// 如果没学习过，那就创建一条记录
			if err := tx.Table(UWSTable).Create(subject.UserWithSubject{
				UserID:       utils.Ptr(userID),
				SubjectID:    utils.Ptr(sbID),
				SubjectType:  utils.Ptr(sbType),
				Phase:        utils.Ptr(int16(0)),
				LearnTimes:   utils.Ptr(int16(1)),
				LastReviewAt: utils.Ptr(now),
			}).Error; err != nil {
				return errors.WithMessage(err, "failed when create user_with_subject record,error: ")
			}

			if err := tx.Table(RemindTable).Create(&subject.Remind{
				UserID:    utils.Ptr(userID),
				SubjectID: utils.Ptr(sbID),
				Remind:    utils.Ptr(utils.GetRemind(0)),
			}).Error; err != nil {
				return errors.WithMessage(err, "failed create remind record,error: ")
			}

			return nil
		}

		if count.Phase < utils.BiggestPhase {
			count.Phase += 1
		}

		// 先前学习过这个题目的
		if err := tx.Table(UWSTable).Where("user_id=? AND subject_id=?", userID, sbID).
			UpdateColumns(map[string]interface{}{
				"phase":          count.Phase,
				"learn_times":    count.LearnTimes + 1,
				"last_review_at": now,
			}).Error; err != nil {
			return errors.WithMessage(err, "failed when update phase,error: ")
		}
		if err := tx.Table(RemindTable).Where("user_id=? AND subject_id=?", userID, sbID).
			UpdateColumn("remind", utils.GetRemind(count.Phase)).Error; err != nil {
			return errors.WithMessage(err, "failed update remind,error: ")
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *SubjectDB) GetEmailByID(ctx context.Context, uid int64) (string, error) {
	var email string
	if err := db.WithContext(ctx).Table(constants.TableNameOFUser).Select("email").Where("id=?", uid).Find(&email).Error; err != nil {
		return "", errors.WithMessage(err, "get email by id failed")
	}
	return email, nil
}
