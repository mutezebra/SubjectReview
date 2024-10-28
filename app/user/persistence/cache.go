package persistence

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/redis/go-redis/v9"

	"github.com/mutezebra/subject-review/pkg/client"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/utils"
)

type UserCache struct {
	*redis.Client
}

func NewUserCache() *UserCache {
	return &UserCache{client.NewRedisClient()}
}

func (cache *UserCache) GetUserID(ctx context.Context) (int64, error) {
	id, err := cache.Incr(ctx, constants.UserIDKey).Result()
	if err != nil {
		return 0, errors.WithMessage(err, "get user id failed")
	}
	return id, nil
}

func (cache *UserCache) PutVerifyCode(ctx context.Context, code string, email string) error {
	if err := cache.Set(ctx, *buildVCKey(&email), code, constants.VerifyCodeDDL).Err(); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("put verify code failed, email: %s", email))
	}
	return nil
}

func (cache *UserCache) GetVerifyCode(ctx context.Context, email string) (string, error) {
	result, err := cache.Get(ctx, *buildVCKey(&email)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", errors.WithMessage(err, fmt.Sprintf("get verify code failed, email: %s", email))
	}
	return result, nil
}

func (cache *UserCache) PutPasswordVerifyCode(ctx context.Context, code string, email string) error {
	if err := cache.Set(ctx, *buildPVCKey(&email), code, constants.VerifyCodeDDL).Err(); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("put password verify code failed, email: %s", email))
	}
	return nil
}

func (cache *UserCache) GetPasswordVerifyCode(ctx context.Context, email string) (string, error) {
	result, err := cache.Get(ctx, *buildPVCKey(&email)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", errors.WithMessage(err, fmt.Sprintf("get password verify code failed, email: %s", email))
	}
	return result, nil
}

func buildVCKey(email *string) *string {
	return utils.Ptr(fmt.Sprintf("%s/%s", constants.VerifyCodePrefix, *email))
}

func buildPVCKey(email *string) *string {
	return utils.Ptr(fmt.Sprintf("%s/%s", constants.PasswordVerifyCodePrefix, *email))
}
