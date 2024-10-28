package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/logger"
	"html/template"
	"path"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/errno"
	"github.com/mutezebra/subject-review/pkg/jwt"
	"github.com/mutezebra/subject-review/pkg/utils"
)

func (svc *UserService) SendEmail(ctx context.Context, email string) error {
	code := utils.GenerateCode(6)

	if err := svc.Email.SendVerifyCode(email, parse(code)); err != nil {
		return err
	}
	if err := svc.UserCache.PutVerifyCode(ctx, code, email); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) SendPasswordVerifyCode(ctx context.Context, email string) error {
	code := utils.GenerateCode(6)

	if err := svc.Email.SendVerifyCode(email, parse(code)); err != nil {
		return err
	}
	if err := svc.UserCache.PutPasswordVerifyCode(ctx, code, email); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) VerifyVerifyCode(ctx context.Context, email string, verifyCode string) error {
	code, err := svc.UserCache.GetVerifyCode(ctx, email)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return errno.New(errno.ExpiredVerifyCode, "verify code have expired,please retry")
	}

	if code != verifyCode {
		return errno.New(errno.WrongVerifyCode, "wrong verify code")
	}
	return nil
}

func (svc *UserService) VerifyPasswordVerifyCode(ctx context.Context, email string, verifyCode string) error {
	code, err := svc.UserCache.GetPasswordVerifyCode(ctx, email)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return errno.New(errno.ExpiredVerifyCode, "verify code have expired,please retry")
	}

	if code != verifyCode {
		return errno.New(errno.WrongVerifyCode, "wrong verify code")
	}
	return nil
}

func (svc *UserService) EncryptPassword(pwd string) (string, error) {
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", errors.WithMessage(err, "failed when encrypt password")
	}
	return string(passwordDigest), nil
}

func (svc *UserService) CheckPassword(pwd, passwordDigest string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(pwd)) == nil
}

func (svc *UserService) GenerateToken(uid int64, email string) (token string, err error) {
	return jwt.GenerateToken(uid, email)
}

func (svc *UserService) UploadAvatar(ctx context.Context, uid int64, fileName string, data []byte) (string, error) {
	ext := path.Ext(fileName)
	name := fmt.Sprintf("%d%s", uid, ext)
	filepath := fmt.Sprintf("%s/%s", config.OSS.AvatarPrefix, name)
	return filepath, svc.OSS.UploadAvatar(ctx, filepath, bytes.NewReader(data))
}

func parse(code string) *string {
	tmpl := template.Must(template.New("verifyCodeTemplate").Parse(constants.VerifyCodeTemplate))

	// 将验证码填充到模板中
	var body bytes.Buffer
	if err := tmpl.Execute(&body, map[string]interface{}{
		"Code": code,
	}); err != nil {
		logger.Fatal(err)
	}
	s := body.String()
	return &s
}
