package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/mutezebra/subject-review/app/user/service"
	"github.com/mutezebra/subject-review/biz/model/user"
	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/errno"
	"github.com/mutezebra/subject-review/pkg/pack"
	"github.com/mutezebra/subject-review/pkg/utils"
)

var usecase *UserUsecase

type UserUsecase struct {
	svc *service.UserService
}

func InitUsecase(svc *service.UserService) {
	usecase = &UserUsecase{svc: svc}
}

func GetUsecase() *UserUsecase {
	return usecase
}

func (uc *UserUsecase) GetVerifyCode(ctx context.Context, req *user.GetVerifyCodeReq) (resp *user.GetVerifyCodeResp, err error) {
	defer func() {
		pack.LogError(err)
	}()
	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	var exist bool
	if exist, _, err = uc.svc.WhetherEmailExists(ctx, req.GetEmail()); err != nil {
		return nil, err
	}
	if exist {
		return nil, errno.New(errno.EmailHaveExisted, fmt.Sprintf("email: %v have exist", req.GetEmail()))
	}
	if err = uc.svc.SendEmail(ctx, req.GetEmail()); err != nil {
		return nil, err
	}

	resp = new(user.GetVerifyCodeResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) GetPasswordVerifyCode(ctx context.Context, req *user.GetVerifyCodeReq) (resp *user.GetVerifyCodeResp, err error) {
	defer func() {
		pack.LogError(err)
	}()
	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	var exist bool
	if exist, _, err = uc.svc.WhetherEmailExists(ctx, req.GetEmail()); err != nil {
		return nil, err
	}

	if !exist {
		return nil, errno.New(errno.EmailNotExisted, fmt.Sprintf("email: %v not exist", req.GetEmail()))
	}

	if err = uc.svc.SendPasswordVerifyCode(ctx, req.GetEmail()); err != nil {
		return nil, err
	}

	resp = new(user.GetVerifyCodeResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	defer func() {
		pack.LogError(err)
	}()
	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	if err = uc.svc.VerifyVerifyCode(ctx, req.GetEmail(), req.GetVerifyCode()); err != nil {
		return nil, err
	}
	var exist bool
	if exist, _, err = uc.svc.WhetherUserNameExists(ctx, req.GetUserName()); err != nil {
		return nil, err
	}

	if exist {
		return nil, errno.New(errno.UserNameExisted, fmt.Sprintf("user name: %s have existed", req.GetUserName()))
	}
	if exist, _, err = uc.svc.WhetherEmailExists(ctx, req.GetEmail()); err != nil {
		return nil, err
	}
	if exist {
		return nil, errno.New(errno.EmailHaveExisted, fmt.Sprintf("email: %s have exist", req.GetEmail()))
	}

	var pwd string
	if pwd, err = uc.svc.EncryptPassword(req.GetPassword()); err != nil {
		return nil, err
	}

	var u *user.User
	if u, err = uc.buildUser(ctx, req.GetUserName(), req.GetEmail(), pwd); err != nil {
		return nil, err
	}

	if err = uc.svc.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	resp = new(user.RegisterResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}

	var pwd string
	if pwd, err = uc.svc.GetUserPasswordByEmail(ctx, req.GetEmail()); err != nil {
		return nil, err
	}
	if pwd == "" {
		return nil, errno.New(errno.EmailNotExisted, fmt.Sprintf("email: %s not exist", req.GetEmail()))
	}

	if !uc.svc.CheckPassword(req.GetPassword(), pwd) {
		return nil, errno.New(errno.WrongPassword, "wrong password")
	}

	var uid int64
	if uid, err = uc.svc.GetUserIDByEmail(ctx, req.GetEmail()); err != nil {
		return nil, err
	}

	var token string
	if token, err = uc.svc.GenerateToken(uid, req.GetEmail()); err != nil {
		return nil, err
	}

	resp = new(user.LoginResp)
	resp.Token = &token
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}

	var u *user.User
	if u, err = uc.svc.GetUserByID(ctx, req.GetUserID()); err != nil {
		return nil, err
	}

	resp = new(user.UserInfoResp)
	resp.Base = pack.Base
	resp.User = uc.user2BaseUser(u)
	return resp, nil
}

func (uc *UserUsecase) UpdateAvatar(ctx context.Context, req *user.UpdateAvatarReq) (resp *user.UpdateAvatarResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}

	var newpath string
	if newpath, err = uc.svc.UploadAvatar(ctx, req.GetUserID(), req.GetAvatarName(), req.GetAvatarData()); err != nil {
		return nil, err
	}
	if err = uc.svc.UpdateInfo(ctx, req.GetUserID(), "avatar", newpath); err != nil {
		return nil, err
	}

	resp = new(user.UpdateAvatarResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) UpdateName(ctx context.Context, req *user.UpdateNameReq) (resp *user.UpdateNameResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}
	if err = uc.svc.UpdateInfo(ctx, req.GetUserID(), "user_name", req.GetUserName()); err != nil {
		return nil, err
	}

	resp = new(user.UpdateNameResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = uc.svc.VerifyReq(req); err != nil {
		return nil, err
	}

	var pwd string
	if req.GetVerifyCode() != "" {
		if err = uc.svc.VerifyPasswordVerifyCode(ctx, req.GetEmail(), req.GetVerifyCode()); err != nil {
			return nil, err
		}

		if pwd, err = uc.svc.EncryptPassword(req.GetNewPassword()); err != nil {
			return nil, err
		}

		if err = uc.svc.UpdateInfoWithEmail(ctx, req.GetEmail(), "password_digest", pwd); err != nil {
			return nil, err
		}

		resp = new(user.UpdatePasswordResp)
		resp.Base = pack.Base
		return resp, nil
	}

	if pwd, err = uc.svc.GetUserPasswordByEmail(ctx, req.GetEmail()); err != nil {
		return nil, err
	}
	if pwd == "" {
		return nil, errno.New(errno.EmailNotExisted, fmt.Sprintf("email: %s not exist", req.GetEmail()))
	}

	if !uc.svc.CheckPassword(req.GetOldPassword(), pwd) {
		return nil, errno.New(errno.WrongPassword, "wrong password")
	}

	if pwd, err = uc.svc.EncryptPassword(req.GetNewPassword()); err != nil {
		return nil, err
	}

	if err = uc.svc.UpdateInfoWithEmail(ctx, req.GetEmail(), "password_digest", pwd); err != nil {
		return nil, err
	}

	resp = new(user.UpdatePasswordResp)
	resp.Base = pack.Base
	return resp, nil
}

func (uc *UserUsecase) buildUser(ctx context.Context, username string, email string, pwd string) (*user.User, error) {
	id, err := uc.svc.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:             &id,
		UserName:       &username,
		Email:          &email,
		PasswordDigest: &pwd,
		Avatar:         &config.OSS.DefaultAvatarPath,
		CreatedAt:      utils.Ptr(time.Now().Unix()),
	}, nil
}

func (uc *UserUsecase) user2BaseUser(u *user.User) *user.BaseUser {
	return &user.BaseUser{
		UserName: u.UserName,
		Email:    u.Email,
		Avatar:   u.Avatar,
	}
}
