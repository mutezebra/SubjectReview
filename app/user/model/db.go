package model

import (
	"context"

	"github.com/mutezebra/subject-review/biz/model/user"
)

type UserDB interface {
	GetUserByID(ctx context.Context, id int64) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetUserPasswordByEmail(ctx context.Context, email string) (string, error)
	GetUserIDByEmail(ctx context.Context, email string) (int64, error)

	WhetherUserExists(ctx context.Context, id int64) (bool, error)
	WhetherEmailExists(ctx context.Context, email string) (bool, int64, error)
	WhetherUserNameExists(ctx context.Context, name string) (bool, int64, error)

	CreateUser(ctx context.Context, user *user.User) error
	UpdateInfo(ctx context.Context, uid int64, filedName string, value string) error
	UpdateInfoWithEmail(ctx context.Context, email string, filedName string, value string) error
	DeleteUser(ctx context.Context, id int64) error
}
