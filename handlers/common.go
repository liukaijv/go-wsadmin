package handlers

import (
	"go-wsadmin/common"
	"go-wsadmin/pkg/win"
)

func ErrorCode(ctx win.Context) {
	ctx.Reply(common.StatusCode)
}

func Download(ctx win.Context) {

}
