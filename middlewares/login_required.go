package middlewares

import (
	"go-wsadmin/pkg/win"
	"go-wsadmin/utils"
)

func LoginRequired(next win.HandlerFunc) win.HandlerFunc {
	return func(ctx win.Context) {
		token := ctx.Request.GetHeader("access_token")
		if token == nil {
			ctx.ReplyError(utils.StatusCodeMustLogin, "")
			return
		}
		next(ctx)
	}
}
