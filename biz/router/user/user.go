// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	user "github.com/mutezebra/subject-review/biz/handler/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_user := _api.Group("/user", _userMw()...)
			_user.POST("/change-avatar", append(_updateavatarMw(), user.UpdateAvatar)...)
			_user.POST("/change-name", append(_updatenameMw(), user.UpdateName)...)
			_user.POST("/change-password", append(_updatepasswordMw(), user.UpdatePassword)...)
			_user.GET("/get-avatar", append(_getuseravatarMw(), user.GetUserAvatar)...)
			_user.GET("/info", append(_userinfoMw(), user.UserInfo)...)
			_user.POST("/login", append(_loginMw(), user.Login)...)
			_user.POST("/password-verify-code", append(_getpasswordverifycodeMw(), user.GetPasswordVerifyCode)...)
			_user.POST("/register", append(_registerMw(), user.Register)...)
			_user.POST("/verify-code", append(_getverifycodeMw(), user.GetVerifyCode)...)
		}
	}
}
