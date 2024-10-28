package usecase

import (
	"context"
	"time"

	"github.com/mutezebra/subject-review/app/subject/service"
	"github.com/mutezebra/subject-review/biz/model/subject"
	"github.com/mutezebra/subject-review/pkg/pack"
	"github.com/mutezebra/subject-review/pkg/utils"
)

var usecase *SubjectUsecase

type SubjectUsecase struct {
	svc *service.SubjectService
}

func InitUsecase(svc *service.SubjectService) {
	usecase = &SubjectUsecase{svc: svc}
}

func GetUsecase() *SubjectUsecase {
	return usecase
}

func (uc *SubjectUsecase) GetSubjects(ctx context.Context, req *subject.GetSubjectsReq) (resp *subject.GetSubjectsResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	var subjects []*subject.BaseSubject
	if subjects, err = uc.svc.GetSubjects(ctx, req.GetSubjectType(), req.GetPages(), req.GetSize()); err != nil {
		return nil, err
	}

	var total int64
	if total, err = uc.svc.GetSubjectNumber(ctx, req.GetSubjectType()); err != nil {
		return nil, err
	}

	resp = new(subject.GetSubjectsResp)
	resp.Total = utils.Ptr(total)
	resp.Subjects = subjects
	resp.Base = pack.Base
	return resp, nil
}

func (uc *SubjectUsecase) AddForgetSubject(ctx context.Context, req *subject.AddForgetSubjectReq) (resp *subject.AddForgetSubjectResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	if err = uc.svc.AddForgetSubject(ctx, req.GetUserID(), req.GetSubjectID(), req.GetSubjectType()); err != nil {
		return nil, err
	}

	resp = new(subject.AddForgetSubjectResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *SubjectUsecase) AddSuccessSubject(ctx context.Context, req *subject.AddSuccessSubjectReq) (resp *subject.AddSuccessSubjectResp, err error) {
	defer func() {
		pack.LogError(err)
	}()
	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	if err = uc.svc.AddSuccessSubject(ctx, req.GetUserID(), req.GetSubjectID(), req.GetSubjectType()); err != nil {
		return nil, err
	}

	resp = new(subject.AddSuccessSubjectResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *SubjectUsecase) GetNeededReviewSubjects(ctx context.Context, req *subject.GetNeededReviewSubjectsReq) (resp *subject.GetNeededReviewSubjectsResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	var subjects []*subject.SubjectRecord
	var total int64

	if subjects, total, err = uc.svc.GetNeedReviewedSubjects(ctx, req.GetUserID(), req.GetPages(), req.GetSize(), time.Now().Unix()); err != nil {
		return nil, err
	}

	resp = new(subject.GetNeededReviewSubjectsResp)
	resp.Total = utils.Ptr(total)
	resp.Subjects = utils.BuildSubjectRecordResp(subjects)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *SubjectUsecase) GetAnswerSubjectRecord(ctx context.Context, req *subject.GetAnswerSubjectRecordReq) (resp *subject.GetAnswerSubjectRecordResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	var subjects []*subject.SubjectRecord
	var total int64

	if subjects, total, err = uc.svc.GetReviewRecords(ctx, req.GetUserID(), req.GetPages(), req.GetSize()); err != nil {
		return nil, err
	}

	resp = new(subject.GetAnswerSubjectRecordResp)
	resp.Total = utils.Ptr(total)
	resp.Subjects = utils.BuildSubjectRecordResp(subjects)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *SubjectUsecase) AddNewSubject(ctx context.Context, req *subject.AddNewSubjectReq) (resp *subject.AddNewSubjectResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	if err = uc.svc.AddSubject(ctx, req.GetUserID(), req.Name, req.Answer, req.GetSubjectType()); err != nil {
		return nil, err
	}

	resp = new(subject.AddNewSubjectResp)
	resp.Base = pack.Base
	return resp, nil
}
