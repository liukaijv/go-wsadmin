package handlers

import (
	"go-wsadmin/common"
	"go-wsadmin/pkg/win"
)

func WithToken(c win.Context, f func(string)) {
	token := c.Request.GetHeader(common.AUTH_TOKEN_KEY)
	if token != nil {
		f(token.(string))
		return
	}
	c.ReplyError(common.StatusCodeForbidden, "认证凭证不存在")
}
