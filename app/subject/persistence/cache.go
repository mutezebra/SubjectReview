package persistence

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/mutezebra/subject-review/pkg/client"
	"github.com/mutezebra/subject-review/pkg/constants"
)

type SubjectCache struct {
	*redis.Client
}

func NewSubjectCache() *SubjectCache {
	return &SubjectCache{client.NewRedisClient()}
}

func (cache *SubjectCache) IncrSubjectNumber(ctx context.Context, sbType int16) (int64, error) {
	id, err := cache.Incr(ctx, buildSubjectType(sbType)).Result()
	if err != nil {
		return 0, errors.WithMessage(err, "IncrSubjectNumber failed")
	}
	return id, nil
}

func (cache *SubjectCache) DecrSubjectNumber(ctx context.Context) (int64, error) {
	id, err := cache.Decr(ctx, constants.SubjectIDKey).Result()
	if err != nil {
		return 0, errors.WithMessage(err, "DecrSubjectNumber failed")
	}
	return id, nil
}

func (cache *SubjectCache) GetSubjectNumber(ctx context.Context, sbType int16) (int64, error) {
	num, err := cache.Get(ctx, buildSubjectType(sbType)).Int64()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, errors.WithMessage(err, "GetSubjectNumber failed")
	}
	return num, nil
}

func (cache *SubjectCache) IsManager(ctx context.Context, uid int64) (bool, error) {
	is, err := cache.SIsMember(ctx, constants.ManagerListKey, uid).Result()
	if err != nil {
		return false, errors.WithMessage(err, "IsManager failed")
	}
	return is, nil
}

func (cache *SubjectCache) GetManagers(ctx context.Context) ([]int64, error) {
	managers, err := cache.SMembers(ctx, constants.ManagerListKey).Result()
	if err != nil {
		return nil, errors.WithMessage(err, "get managers failed")
	}
	results := make([]int64, len(managers))
	for i, m := range managers {
		result, _ := strconv.ParseInt(m, 10, 64)
		results[i] = result
	}
	return results, nil
}

func buildSubjectType(sbType int16) string {
	return fmt.Sprintf("%s:%d", constants.SubjectIDKey, sbType)
}
