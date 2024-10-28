package model

import (
	"context"
)

type UserCache interface {
	GetUserID(ctx context.Context) (int64, error)
	PutVerifyCode(ctx context.Context, code string, email string) error
	GetVerifyCode(ctx context.Context, email string) (string, error)
	GetPasswordVerifyCode(ctx context.Context, email string) (string, error)
	PutPasswordVerifyCode(ctx context.Context, code string, email string) error
}
