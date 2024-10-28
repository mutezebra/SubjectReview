// Code generated by hertz generator.

package subject

import (
	"context"
	"github.com/mutezebra/subject-review/app/subject/usecase"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/pack"
	"github.com/mutezebra/subject-review/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	subject "github.com/mutezebra/subject-review/biz/model/subject"
)

// GetSubjects .
// @router /api/subject/get [POST]
func GetSubjects(ctx context.Context, c *app.RequestContext) {
	var err error
	var req subject.GetSubjectsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.UserID = utils.Ptr(ctx.Value(constants.TokenKey).(int64))

	var resp *subject.GetSubjectsResp
	resp, err = usecase.GetUsecase().GetSubjects(ctx, &req)
	if err != nil {
		pack.SendErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// AddSuccessSubject .
// @router /api/subject/add-success [POST]
func AddSuccessSubject(ctx context.Context, c *app.RequestContext) {
	var err error
	var req subject.AddSuccessSubjectReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.UserID = utils.Ptr(ctx.Value(constants.TokenKey).(int64))

	var resp *subject.AddSuccessSubjectResp
	resp, err = usecase.GetUsecase().AddSuccessSubject(ctx, &req)
	if err != nil {
		pack.SendErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetNeededReviewSubjects .
// @router /api/subject/get-needed-review [POST]
func GetNeededReviewSubjects(ctx context.Context, c *app.RequestContext) {
	var err error
	var req subject.GetNeededReviewSubjectsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.UserID = utils.Ptr(ctx.Value(constants.TokenKey).(int64))

	var resp *subject.GetNeededReviewSubjectsResp
	resp, err = usecase.GetUsecase().GetNeededReviewSubjects(ctx, &req)
	if err != nil {
		pack.SendErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetAnswerSubjectRecord .
// @router /api/subject/get-answer-record [POST]
func GetAnswerSubjectRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req subject.GetAnswerSubjectRecordReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.UserID = utils.Ptr(ctx.Value(constants.TokenKey).(int64))

	var resp *subject.GetAnswerSubjectRecordResp
	resp, err = usecase.GetUsecase().GetAnswerSubjectRecord(ctx, &req)
	if err != nil {
		pack.SendErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// AddNewSubject .
// @router /api/subject/add-new-subject [POST]
func AddNewSubject(ctx context.Context, c *app.RequestContext) {
	var err error
	var req subject.AddNewSubjectReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.UserID = utils.Ptr(ctx.Value(constants.TokenKey).(int64))

	var resp *subject.AddNewSubjectResp
	resp, err = usecase.GetUsecase().AddNewSubject(ctx, &req)
	if err != nil {
		pack.SendErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// AddForgetSubject .
// @router /api/subject/add-forget [POST]
func AddForgetSubject(ctx context.Context, c *app.RequestContext) {
	var err error
	var req subject.AddForgetSubjectReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.UserID = utils.Ptr(ctx.Value(constants.TokenKey).(int64))

	var resp *subject.AddForgetSubjectResp
	resp, err = usecase.GetUsecase().AddForgetSubject(ctx, &req)
	if err != nil {
		pack.SendErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}