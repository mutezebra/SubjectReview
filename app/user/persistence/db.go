package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/mutezebra/subject-review/biz/model/user"
	"github.com/mutezebra/subject-review/pkg/client"
	"github.com/mutezebra/subject-review/pkg/constants"
)

type UserDB struct {
	*gorm.DB
}

func NewUserDB() *UserDB {
	return &UserDB{client.GetMysqlDB(constants.TableNameOFUser)}
}

var nodel = "deleted_at IS NULL"

func (db *UserDB) GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	u := user.User{ID: &id}
	if err := db.WithContext(ctx).Where(nodel).First(&u).Error; err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("can`t find user by id: %d,err: ", id))
	}
	return &u, nil
}

func (db *UserDB) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	if err := db.WithContext(ctx).Where(nodel).
		Where("email=?", email).Limit(1).
		Find(&u).Error; err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("can`t find user by email: %s,err: ", email))
	}
	return &u, nil
}

func (db *UserDB) GetUserPasswordByEmail(ctx context.Context, email string) (string, error) {
	var pwd string
	if err := db.WithContext(ctx).Where(nodel).
		Select("password_digest").
		Find(&pwd).Error; err != nil {
		return "", errors.WithMessage(err, fmt.Sprintf("can`t find pwd by email: %s,err: ", email))
	}
	return pwd, nil
}

func (db *UserDB) GetUserIDByEmail(ctx context.Context, email string) (int64, error) {
	var uid int64
	if err := db.WithContext(ctx).Where(nodel).
		Select("id").Limit(1).Where("email=?", email).
		Find(&uid).Error; err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("can`t find uid by email: %s,err: ", email))
	}
	return uid, nil
}

func (db *UserDB) WhetherUserExists(ctx context.Context, id int64) (bool, error) {
	var findID int64
	if err := db.WithContext(ctx).Where(nodel).
		Select("id").Limit(1).Where("id=?", id).
		Find(&findID).Error; err != nil {
		return false, errors.WithMessage(err, fmt.Sprintf("failed when query user exist,id: %d,error: ", id))
	}
	return findID != 0, nil
}

func (db *UserDB) WhetherEmailExists(ctx context.Context, email string) (bool, int64, error) {
	var findID int64
	if err := db.WithContext(ctx).Where(nodel).
		Select("id").Limit(1).Where("email=?", email).
		Find(&findID).Error; err != nil {
		return false, findID, errors.WithMessage(err, fmt.Sprintf("failed when query email exist,email: %s,error: ", email))
	}
	return findID != 0, findID, nil
}

func (db *UserDB) WhetherUserNameExists(ctx context.Context, name string) (bool, int64, error) {
	var findID int64
	if err := db.WithContext(ctx).Where(nodel).
		Select("id").Limit(1).Where("user_name=?", name).
		Find(&findID).Error; err != nil {
		return false, findID, errors.WithMessage(err, fmt.Sprintf("failed when query email exist,name: %s,error: ", name))
	}
	return findID != 0, findID, nil
}

func (db *UserDB) CreateUser(ctx context.Context, u *user.User) error {
	if err := db.WithContext(ctx).Create(u).Error; err != nil {
		return errors.WithMessage(err, "create user failed")
	}
	return nil
}

func (db *UserDB) UpdateInfo(ctx context.Context, uid int64, filedName string, value string) error {
	if err := db.WithContext(ctx).Where(nodel).Where("id=?", uid).
		UpdateColumn(filedName, value).Error; err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed when update %d`s %s to %s", uid, filedName, value))
	}
	return nil
}

func (db *UserDB) UpdateInfoWithEmail(ctx context.Context, email string, filedName string, value string) error {
	if err := db.WithContext(ctx).Where(nodel).Where("email=?", email).
		UpdateColumn(filedName, value).Error; err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed when update %s`s %s to %s", email, filedName, value))
	}
	return nil
}

func (db *UserDB) DeleteUser(ctx context.Context, id int64) error {
	if err := db.WithContext(ctx).Where(nodel).
		Where("id=?", id).UpdateColumn("deleted_at", time.Now().Unix()).
		Error; err != nil {
		return errors.WithMessage(err, fmt.Sprintf("failed when delete %d,err: ", id))
	}
	return nil
}
