package service

import (
	"github.com/mutezebra/subject-review/app/subject/model"
	"github.com/mutezebra/subject-review/biz/model/subject"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/errno"
	"github.com/mutezebra/subject-review/pkg/logger"
)

type SubjectService struct {
	model.SubjectDB
	model.SubjectCache
	Email model.Email
}

func NewSubjectService(svc *SubjectService) *SubjectService {
	if svc.SubjectCache == nil {
		logger.Fatal("subject`s cache should not be nil")
	}
	if svc.SubjectDB == nil {
		logger.Fatal("subject`s db should not be nil")
	}
	if svc.Email == nil {
		logger.Fatal("subject`s email should not be nil")
	}

	go svc.Remind()

	return svc
}

func (svc *SubjectService) VerifyReq(r any) error {
	switch req := r.(type) {
	case *subject.GetSubjectsReq:
		if err := svc.verifySubjectType(req.GetSubjectType()); err != nil {
			return err
		}
		if err := svc.verifyPages(req.GetPages()); err != nil {
			return err
		}
		if err := svc.verifySize(req.GetSize()); err != nil {
			return err
		}
	case *subject.AddForgetSubjectReq:
		if err := svc.verifySubjectID(req.GetSubjectID()); err != nil {
			return err
		}
		if err := svc.verifySubjectType(req.GetSubjectType()); err != nil {
			return err
		}
	case *subject.AddSuccessSubjectReq:
		if err := svc.verifySubjectID(req.GetSubjectID()); err != nil {
			return err
		}
		if err := svc.verifySubjectType(req.GetSubjectType()); err != nil {
			return err
		}
	case *subject.GetAnswerSubjectRecordReq:
		if err := svc.verifyPages(req.GetPages()); err != nil {
			return err
		}
		if err := svc.verifySize(req.GetSize()); err != nil {
			return err
		}
	case *subject.GetNeededReviewSubjectsReq:
		if err := svc.verifyPages(req.GetPages()); err != nil {
			return err
		}
		if err := svc.verifySize(req.GetSize()); err != nil {
			return err
		}
	case *subject.AddNewSubjectReq:
		if err := svc.verifySubjectType(req.GetSubjectType()); err != nil {
			return err
		}
		if err := svc.verifySubjectName(req.Name); err != nil {
			return err
		}

	default:
		return errno.New(errno.UnknownSubjectReqType, "unknown subject request type")
	}
	return nil
}

func (svc *SubjectService) verifySubjectID(id int64) error {
	if id <= 0 {
		return errno.New(errno.SubjectIDLessThanZero, "not subject`s id less than id")
	}
	return nil
}

func (svc *SubjectService) verifySubjectName(v *string) error {
	if len(*v) >= constants.SubjectNameLength || len(*v) == 0 {
		return errno.New(errno.LengthOFSubjectNameOverLimit, "name is to long or zero")
	}
	return nil
}

func (svc *SubjectService) verifySubjectType(v int16) error {
	for i := range constants.TypeSlice {
		if v == constants.TypeSlice[i] {
			return nil
		}
	}
	return errno.New(errno.UnknownSubjectType, "un supported subject type")
}

func (svc *SubjectService) verifyPages(v int64) error {
	if v <= 0 {
		return errno.New(errno.PagesLessThanZero, "pages should bigger than 0")
	}
	return nil
}

func (svc *SubjectService) verifySize(v int64) error {
	if v <= 0 {
		return errno.New(errno.SizeLessThanZero, "size should bigger than 0")
	}
	return nil
}
