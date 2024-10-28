package persistence

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/mutezebra/subject-review/app/user/model"
	"github.com/mutezebra/subject-review/biz/model/user"
	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/utils"
)

func initConfig(t *testing.T) {
	t.Helper()
	config.InitConfig()
}

func modelDB(t *testing.T, db *UserDB) model.UserDB {
	t.Helper()
	return db
}

func TestDB(t *testing.T) {
	initConfig(t)
	db := modelDB(t, NewUserDB())

	uid := int64(22211)
	name := "mz111"
	email := "82711@qq.com"
	pwd := "password"
	avatar := "/avatar/default.jpg"
	ctx := context.Background()

	u := &user.User{
		ID:             utils.Ptr(uid),
		UserName:       utils.Ptr(name),
		Email:          utils.Ptr(email),
		PasswordDigest: utils.Ptr(pwd),
		Avatar:         utils.Ptr(avatar),
		CreatedAt:      utils.Ptr(time.Now().Unix()),
	}

	Convey("TestCreateUser", t, func() {
		err := db.CreateUser(ctx, u)
		So(err, ShouldBeNil)
	})

	Convey("Test Get", t, func() {
		Convey("Test Get user by id", func() {
			us, err := db.GetUserByID(ctx, uid)
			So(us.GetID(), ShouldEqual, uid)
			So(us.GetUserName(), ShouldEqual, name)
			So(us.GetEmail(), ShouldEqual, email)
			So(us.GetPasswordDigest(), ShouldEqual, pwd)
			So(us.GetAvatar(), ShouldEqual, avatar)
			So(err, ShouldBeNil)
		})
		Convey("Test Get user by email", func() {
			us, err := db.GetUserByEmail(ctx, email)
			So(us.GetID(), ShouldEqual, uid)
			So(us.GetUserName(), ShouldEqual, name)
			So(us.GetEmail(), ShouldEqual, email)
			So(us.GetPasswordDigest(), ShouldEqual, pwd)
			So(us.GetAvatar(), ShouldEqual, avatar)
			So(err, ShouldBeNil)
		})
		Convey("Test Get password by email", func() {
			pw, err := db.GetUserPasswordByEmail(ctx, email)
			So(pw, ShouldEqual, pwd)
			So(err, ShouldBeNil)
		})
		Convey("Test Get userid by email", func() {
			id, err := db.GetUserIDByEmail(ctx, email)
			So(id, ShouldEqual, uid)
			So(err, ShouldBeNil)
		})
	})

	Convey("Test Whether user exists", t, func() {
		exist, err := db.WhetherUserExists(ctx, uid)
		So(exist, ShouldBeTrue)
		So(err, ShouldBeNil)
	})

	Convey("Test Whether email exists", t, func() {
		exist, id, err := db.WhetherEmailExists(ctx, email)
		So(exist, ShouldBeTrue)
		So(id, ShouldEqual, uid)
		So(err, ShouldBeNil)

		Convey("Test fake mail ", func() {
			exist, id, err = db.WhetherEmailExists(ctx, "fake mail")
			So(exist, ShouldBeFalse)
			So(id, ShouldEqual, int64(0))
			So(err, ShouldBeNil)
		})
	})

	Convey("Test Whether user name exists", t, func() {
		exist, id, err := db.WhetherUserNameExists(ctx, name)
		So(exist, ShouldBeTrue)
		So(id, ShouldEqual, uid)
		So(err, ShouldBeNil)
	})

	Convey("Test Update info", t, func() {
		err := db.UpdateInfo(ctx, uid, "user_name", "new name")
		So(err, ShouldBeNil)
		Convey("Test get user name after update", func() {
			us, err := db.GetUserByID(ctx, uid)
			So(us.GetUserName(), ShouldEqual, "new name")
			So(err, ShouldBeNil)
		})
	})

	Convey("Test Update info with email", t, func() {
		err := db.UpdateInfoWithEmail(ctx, email, "user_name", "mz")
		So(err, ShouldBeNil)
		Convey("Test get user name after update", func() {
			us, err := db.GetUserByEmail(ctx, email)
			So(us.GetUserName(), ShouldEqual, "mz")
			So(err, ShouldBeNil)
		})
	})

	Convey("Test Delete user", t, func() {
		err := db.DeleteUser(ctx, uid)
		So(err, ShouldBeNil)
		Convey("Test whether user exists after delete", func() {
			exist, err := db.WhetherUserExists(ctx, uid)
			So(exist, ShouldBeFalse)
			So(err, ShouldBeNil)
		})
	})
}
