package model

import (
	"context"

	"github.com/mutezebra/subject-review/biz/model/subject"
)

type SubjectDB interface {
	GetSubjects(ctx context.Context, sbType int16, pages, size int64) ([]*subject.BaseSubject, error)
	GetSubject(ctx context.Context, sbID int64) (*subject.Subject, error)
	GetReviewRecords(ctx context.Context, uid int64, pages, size int64) ([]*subject.SubjectRecord, int64, error)
	GetNeedReviewedSubjects(ctx context.Context, uid int64, pages, size int64, endTime int64) ([]*subject.SubjectRecord, int64, error)

	CreateSubject(ctx context.Context, subject *subject.Subject) error
	AddForgetSubject(ctx context.Context, userID, sbID int64, sbType int16) error
	AddSuccessSubject(ctx context.Context, userID, sbID int64, sbType int16) error
	GetEmailByID(ctx context.Context, uid int64) (string, error)
}
