package persistence

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/mutezebra/subject-review/app/user/model"
)

func modelCache(cache *UserCache) model.UserCache {
	return cache
}

func TestCache(t *testing.T) {
	initConfig(t)
	client := modelCache(NewUserCache())
	code := "123456"
	email := "email"
	ctx := context.Background()

	Convey("Test Put Verify code", t, func() {
		err := client.PutVerifyCode(ctx, code, email)
		So(err, ShouldBeNil)
	})

	Convey("Test Get Verify code", t, func() {
		result, err := client.GetVerifyCode(ctx, email)
		So(result, ShouldEqual, code)
		So(err, ShouldBeNil)
	})
}
