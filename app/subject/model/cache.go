package model

import "context"

type SubjectCache interface {
	IncrSubjectNumber(ctx context.Context, sbType int16) (int64, error)
	DecrSubjectNumber(ctx context.Context) (int64, error)
	GetSubjectNumber(ctx context.Context, sbType int16) (int64, error)
	IsManager(ctx context.Context, uid int64) (bool, error)
	GetManagers(ctx context.Context) ([]int64, error)
}
