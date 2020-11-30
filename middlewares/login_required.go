package middlewares

import (
	"go-wsadmin/common"
	"go-wsadmin/pkg/win"
)

func LoginRequired(next win.HandlerFunc) win.HandlerFunc {
	return func(ctx win.Context) {
		token := ctx.Request.GetHeader(common.AUTH_TOKEN_KEY)
		if token == nil {
			ctx.ReplyError(common.StatusCodeMustLogin, "")
			return
		}
		next(ctx)
	}
}
