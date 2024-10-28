package persistence

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/mutezebra/subject-review/app/subject/model"
	"github.com/mutezebra/subject-review/biz/model/subject"
	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/utils"
)

func initConfig(t *testing.T) {
	t.Helper()
	config.InitConfig()
}

func modelDB(t *testing.T, db *SubjectDB) model.SubjectDB {
	t.Helper()
	return db
}

func buildSb() *subject.Subject {
	return &subject.Subject{
		ID:          utils.Ptr(id),
		Name:        utils.Ptr(name),
		Answer:      utils.Ptr(answer),
		SubjectType: utils.Ptr(sbTYpe),
		CreatorID:   utils.Ptr(uid),
		CreatedAt:   utils.Ptr(now),
	}
}

var (
	id     = int64(123)
	name   = "Go语言从入门到放弃"
	answer = "没有放弃可言"
	sbTYpe = int16(constants.STGo)
	uid    = int64(1)
	now    = time.Now().Unix()
)

func TestCreateSubject(t *testing.T) {
	initConfig(t)
	db := modelDB(t, NewSubjectDB())

	sb := buildSb()
	ctx := context.Background()

	Convey("Test Create Subject", t, func() {
		err := db.CreateSubject(ctx, sb)
		So(err, ShouldBeNil)

		var s *subject.Subject
		s, err = db.GetSubject(ctx, sb.GetID())
		So(err, ShouldBeNil)
		So(s.GetID(), ShouldEqual, sb.GetID())
		So(s.GetName(), ShouldEqual, sb.GetName())
		So(s.GetAnswer(), ShouldEqual, sb.GetAnswer())
		So(s.GetSubjectType(), ShouldEqual, sb.GetSubjectType())
		So(s.GetCreatorID(), ShouldEqual, sb.GetCreatorID())
		So(s.GetCreatedAt(), ShouldEqual, sb.GetCreatedAt())
	})

	Convey("Test Get Subjects", t, func() {
		results, err := db.GetSubjects(ctx, sbTYpe, 1, 10)
		So(err, ShouldBeNil)
		So(len(results), ShouldEqual, 1)
		s := results[0]
		So(s.GetID(), ShouldEqual, sb.GetID())
		So(s.GetName(), ShouldEqual, sb.GetName())
		So(s.GetAnswer(), ShouldEqual, sb.GetAnswer())
		So(s.GetSubjectType(), ShouldEqual, sb.GetSubjectType())
	})
}

func TestSubjectDBAddSubject(t *testing.T) {
	initConfig(t)
	db := modelDB(t, NewSubjectDB())

	ctx := context.Background()

	Convey("Test add forget subject", t, func() {
		err := db.AddForgetSubject(ctx, uid, id, sbTYpe)
		So(err, ShouldBeNil)
	})

	Convey("Test add forget subject", t, func() {
		err := db.AddSuccessSubject(ctx, uid, id, sbTYpe)
		So(err, ShouldBeNil)
	})
}

func TestReviewRecords(t *testing.T) {
	initConfig(t)
	db := modelDB(t, NewSubjectDB())

	sb := buildSb()
	ctx := context.Background()

	Convey("Test review records", t, func() {
		records, err := db.GetReviewRecords(ctx, uid, 1, 10)
		So(err, ShouldBeNil)
		So(len(records), ShouldEqual, 1)
		s := records[0]
		So(s.GetID(), ShouldEqual, sb.GetID())
		So(s.GetName(), ShouldEqual, sb.GetName())
		So(s.GetAnswer(), ShouldEqual, sb.GetAnswer())
		So(s.GetSubjectType(), ShouldEqual, sb.GetSubjectType())
		So(s.GetPhase(), ShouldEqual, 1)
		So(s.GetLearnTimes(), ShouldEqual, 2)
	})
}
