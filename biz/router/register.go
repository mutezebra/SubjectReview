// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	base "github.com/mutezebra/subject-review/biz/router/base"
	subject "github.com/mutezebra/subject-review/biz/router/subject"
	user "github.com/mutezebra/subject-review/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user.Register(r)

	subject.Register(r)

	base.Register(r)

}
