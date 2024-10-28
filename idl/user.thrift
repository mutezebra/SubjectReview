namespace go user

include "base.thrift"

struct User {
    1: optional i64 ID
    2: optional string UserName
    3: optional string Email
    4: optional string PasswordDigest
    5: optional string Avatar
    6: optional i64 CreatedAt
    7: optional i64 DeletedAt
}

struct BaseUser {
    1: optional string UserName (api.body="user_name")
    2: optional string Email (api.body="email")
    3: optional string Avatar (api.body="avatar")
}

struct GetVerifyCodeReq {
    1: optional string Email (api.form="email")
}

struct GetVerifyCodeResp {
    1: optional base.Base Base (api.body="base")
}

struct RegisterReq {
    1: optional string UserName (api.form="user_name")
    2: optional string Password (api.form="password")
    3: optional string Email (api.body="email")
    4: optional string VerifyCode (api.body="verify_code")
}

struct RegisterResp {
    1: optional base.Base Base (api.body="base")
}

struct LoginReq {
    1: optional string Email (api.form="email")
    2: optional string Password (api.form="password")
}

struct LoginResp {
    1: optional base.Base Base (api.body="base")
    2: optional string Token (api.body="token")
}

struct UserInfoReq {
    1: optional i64 UserID
}

struct UserInfoResp {
    1: optional base.Base Base (api.body="base")
    2: optional BaseUser User (api.body="user")
}

struct UpdateAvatarReq {
    1: optional i64 UserID
    2: optional string AvatarName
    3: optional binary AvatarData
}

struct UpdateAvatarResp {
    1: optional base.Base Base (api.body="base")
}

struct UpdateNameReq {
    1: optional i64 UserID
    2: optional string UserName (api.form="user_name")
}

struct UpdateNameResp {
    1: optional base.Base Base (api.body="base")
}

struct UpdatePasswordReq {
    1: optional string Email (api.form="email")
    2: optional string OldPassword (api.form="old_password")
    3: optional string NewPassword (api.form="new_password")
    4: optional string VerifyCode  (api.form="verify_code")
}

struct UpdatePasswordResp {
    1: optional base.Base Base (api.body="base")
}

struct GetUserAvatarReq {
    1: optional i64 UserID
}

struct GetUserAvatarResp {
    1: optional base.Base Base (api.body="base")
    2: optional string AvatarName (api.body="avatar_name")
    3: optional string AvatarUrl (api.body="avatar_url")
}

service UserService {
    GetVerifyCodeResp GetVerifyCode(1: GetVerifyCodeReq Req) (api.post="/api/user/verify-code")
    GetVerifyCodeResp GetPasswordVerifyCode(1: GetVerifyCodeReq Req) (api.post="/api/user/password-verify-code")
    RegisterResp Register(1: RegisterReq Req) (api.post="/api/user/register")
    LoginResp Login(1: LoginReq Req) (api.post="/api/user/login")
    UserInfoResp UserInfo(1: UserInfoReq Req)   (api.get="/api/user/info")
    UpdateAvatarResp UpdateAvatar(1: UpdateAvatarReq Req) (api.post="/api/user/change-avatar")
    UpdateNameResp UpdateName(1: UpdateNameReq Req) (api.post="/api/user/change-name")
    UpdatePasswordResp UpdatePassword(1: UpdatePasswordReq Req) (api.post="/api/user/change-password")
    GetUserAvatarResp GetUserAvatar(1: GetUserAvatarReq Req) (api.get="/api/user/get-avatar")
}