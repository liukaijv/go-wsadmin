package handlers

import (
	"go-wsadmin/domain/req"
	"go-wsadmin/domain/resp"
	"go-wsadmin/pkg/win"
	"go-wsadmin/utils"
)

func Login(ctx win.Context) {
	var loginReq req.LoginReq
	if err := ctx.BindJson(&loginReq); err != nil {
		ctx.ReplyError(utils.StatusCodeParamError, utils.GetStatusDisplay(utils.StatusCodeParamError))
	}

	var loginResp = resp.LoginResp{}
	ctx.Reply(loginResp)
}

func Logout(ctx win.Context) {

}

func GetUserInfo(ctx win.Context) {

}
