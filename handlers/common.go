package handlers

import (
	"go-wsadmin/pkg/win"
	"go-wsadmin/utils"
)

func ErrorCode(ctx win.Context) {
	ctx.Reply(utils.StatusCode)
}
