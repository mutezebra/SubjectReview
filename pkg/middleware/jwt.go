package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/mutezebra/subject-review/pkg/constants"
	"github.com/mutezebra/subject-review/pkg/errno"
	"github.com/mutezebra/subject-review/pkg/jwt"
	"github.com/mutezebra/subject-review/pkg/pack"
)

func JWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader(constants.TokenHeader))
		if token == "" {
			pack.SendErrResp(c, errno.New(errno.LackToken, "lack of token"), consts.StatusUnauthorized)
			c.Abort()
			return
		}

		uid, _, err := jwt.CheckToken(token)
		if uid == 0 {
			pack.SendErrResp(c, errno.New(errno.WrongToken, err.Error()), consts.StatusUnauthorized)
			c.Abort()
			return
		}
		if err != nil {
			pack.SendErrResp(c, errno.New(errno.TokenExpire, "token is expired, please log in again"), consts.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next(context.WithValue(ctx, constants.TokenKey, uid))
	}
}
