package service

import (
	"context"
	"time"

	"github.com/mutezebra/subject-review/biz/model/subject"
	"github.com/mutezebra/subject-review/pkg/errno"
	"github.com/mutezebra/subject-review/pkg/utils"
)

func (svc *SubjectService) AddForgetSubject(ctx context.Context, userID int64, subjectID int64, subjectType int16) error {
	num, err := svc.SubjectCache.GetSubjectNumber(ctx, subjectType)
	if err != nil {
		return err
	}

	if subjectID > num {
		return errno.New(errno.SubjectNotExist, "subject not exist")
	}

	return svc.SubjectDB.AddForgetSubject(ctx, userID, subjectID, subjectType)
}

func (svc *SubjectService) AddSuccessSubject(ctx context.Context, userID int64, subjectID int64, subjectType int16) error {
	num, err := svc.SubjectCache.GetSubjectNumber(ctx, subjectType)
	if err != nil {
		return err
	}

	if subjectID > num {
		return errno.New(errno.SubjectNotExist, "subject not exist")
	}

	return svc.SubjectDB.AddSuccessSubject(ctx, userID, subjectID, subjectType)
}

func (svc *SubjectService) AddSubject(ctx context.Context, userID int64, name, answer *string, subjectType int16) error {
	is, err := svc.SubjectCache.IsManager(ctx, userID)
	if err != nil {
		return err
	}
	if !is {
		return errno.New(errno.NotManager, "you are not the manager, can`t add new subject")
	}

	id, err := svc.SubjectCache.IncrSubjectNumber(ctx, subjectType)
	if err != nil {
		return err
	}
	s := subject.Subject{
		ID:          utils.Ptr(id),
		Name:        name,
		Answer:      answer,
		SubjectType: utils.Ptr(subjectType),
		CreatorID:   utils.Ptr(userID),
		CreatedAt:   utils.Ptr(time.Now().Unix()),
	}
	if err = svc.SubjectDB.CreateSubject(ctx, &s); err != nil {
		_, _ = svc.SubjectCache.DecrSubjectNumber(ctx)
		return err
	}
	return nil
}
