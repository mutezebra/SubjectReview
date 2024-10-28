package service

import (
	"fmt"
	"path"
	"regexp"

	"github.com/pkg/errors"

	"github.com/mutezebra/subject-review/app/user/model"
	"github.com/mutezebra/subject-review/biz/model/user"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/errno"
	"github.com/mutezebra/subject-review/pkg/logger"
)

type UserService struct {
	model.UserDB
	model.UserCache
	Email   model.Email
	OSS     model.OSS
	PwdRe   *regexp.Regexp
	EmailRe *regexp.Regexp
}

func NewUserService(svc *UserService) *UserService {
	if svc.UserDB == nil {
		logger.Fatal("db is required")
	}
	if svc.UserCache == nil {
		logger.Fatal("cache is required")
	}
	if svc.Email == nil {
		logger.Fatal("email is required")
	}
	if svc.OSS == nil {
		logger.Fatal("oss is required")
	}
	if svc.PwdRe == nil {
		logger.Fatal("PwdRe is required")
	}
	if svc.EmailRe == nil {
		logger.Fatal("EmailRe is required")
	}

	return &UserService{
		UserDB:    svc.UserDB,
		UserCache: svc.UserCache,
		Email:     svc.Email,
		OSS:       svc.OSS,
		PwdRe:     svc.PwdRe,
		EmailRe:   svc.EmailRe,
	}
}

func (svc *UserService) VerifyReq(r interface{}) (err error) {
	switch req := r.(type) {
	case *user.GetVerifyCodeReq:
		if err = svc.verifyEmail(req.GetEmail()); err != nil {
			return err
		}
	case *user.RegisterReq:
		if err = svc.verifyUserName(req.GetUserName()); err != nil {
			return err
		}
		if err = svc.verifyVerifyCode(req.GetVerifyCode()); err != nil {
			return err
		}
		if err = svc.verifyPassword(req.GetPassword()); err != nil {
			return err
		}
		if err = svc.verifyEmail(req.GetEmail()); err != nil {
			return err
		}
	case *user.LoginReq:
		if err = svc.verifyPassword(req.GetPassword()); err != nil {
			return err
		}
		if err = svc.verifyEmail(req.GetEmail()); err != nil {
			return err
		}
	case *user.UserInfoReq:
	case *user.UpdateAvatarReq:
		if err = svc.verifyAvatarFile(req.GetAvatarData()); err != nil {
			return err
		}
		if err = svc.verifyAvatarName(req.GetAvatarName()); err != nil {
			return err
		}
	case *user.UpdateNameReq:
		if err = svc.verifyUserName(req.GetUserName()); err != nil {
			return err
		}
	case *user.UpdatePasswordReq:
		if err = svc.verifyEmail(req.GetEmail()); err != nil {
			return err
		}
		if err = svc.verifyPassword(req.GetNewPassword()); err != nil {
			return err
		}

	default:
		return errors.New(fmt.Sprintf("Unknown req type, type: %v", req))
	}
	return nil
}

func (svc *UserService) verifyEmail(email string) error {
	if !svc.EmailRe.MatchString(email) {
		return errno.New(errno.InvalidEmailFormat, "invalid Email format")
	}
	return nil
}

func (svc *UserService) verifyPassword(pw string) error {
	if !svc.PwdRe.MatchString(pw) {
		return errno.New(errno.InvalidPasswordFormat, "invalid password format")
	}
	return nil
}

func (svc *UserService) verifyUserName(name string) error {
	if len(name) == 0 || len(name) >= 20 {
		return errno.New(errno.InvalidUserNameFormat, "user name length should be not small than 0,and bigger than 20")
	}

	for i := range name {
		if name[i] == ' ' || name[i] == '\n' || name[i] == '\t' {
			return errno.New(errno.SpaceUserName, "there should be no white space or line breaks in the user name")
		}
	}
	return nil
}

func (svc *UserService) verifyVerifyCode(code string) error {
	if len(code) != 6 {
		return errno.New(errno.InvalidVerifyCodeFormat, "the length of verify code shoule be 6")
	}
	return nil
}

func (svc *UserService) verifyAvatarName(name string) error {
	ext := path.Ext(name)
	if ext != "jpg" && ext != "png" && ext != "jpeg" && ext != "JPG" && ext != "JPEG" && ext != "PNG" {
		return errno.New(errno.UnsupportedAvatarFormat, "unsupported avatar format")
	}
	return nil
}

func (svc *UserService) verifyAvatarFile(data []byte) error {
	if len(data) == 0 {
		return errno.New(errno.AvatarFileTooSmall, "the size of avatar size is too small")
	}
	if len(data) >= constants.MaxAvatarSize {
		return errno.New(errno.AvatarFileOverLimit, "the size of avatar size have over the limit")
	}
	return nil
}
