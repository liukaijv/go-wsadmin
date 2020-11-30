package handlers

import (
	"go-wsadmin/common"
	"go-wsadmin/domain/req"
	"go-wsadmin/domain/resp"
	"go-wsadmin/models"
	"go-wsadmin/pkg/win"
	"go-wsadmin/utils"
	"gorm.io/gorm"
	"time"
)

func Login(ctx win.Context) {

	var req req.LoginReq
	if err := ctx.BindJson(&req); err != nil {
		ctx.ReplyError(common.StatusCodeParamError, common.GetStatusDisplay(common.StatusCodeParamError))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.ReplyError(common.StatusCodeParamError, err.Error())
		return
	}

	if req.Username == "" {
		ctx.ReplyError(common.StatusCodeParamError, "用户名不能为空")
		return
	}

	var admin models.Admin
	err := models.DB().Where("username = ?", req.Username).Find(&admin).Error
	if err == gorm.ErrRecordNotFound {
		ctx.ReplyError(common.StatusCodeDataNotExist, "用户不存在")
		return
	}

	if !utils.VerifyPassword(admin.Password, req.Password) {
		ctx.ReplyError(common.StatusCodeForbidden, "账号或密码错误")
		return
	}

	var loginResp = resp.LoginResp{
		Token: generateToken(admin),
	}

	ctx.Reply(loginResp)
}

func generateToken(admin models.Admin) string {
	token := utils.RandomString(32)
	expire := time.Now().Add(30 * time.Minute)
	admin.Token = token
	admin.TokenExpire = &expire
	models.DB().Save(&admin)
	return token
}

func Logout(ctx win.Context) {
	WithToken(ctx, func(token string) {

		var admin models.Admin
		err := models.DB().Where("token = ?", token).Find(&admin).Error
		if err == gorm.ErrRecordNotFound {
			ctx.ReplyError(common.StatusCodeDataNotExist, "信息不存在")
			return
		}
		admin.Token = ""
		admin.TokenExpire = nil
		models.DB().Save(&admin)
		ctx.Reply(nil)

	})
}

func GetAuthInfo(ctx win.Context) {
	WithToken(ctx, func(token string) {

		var admin models.Admin
		err := models.DB().Where("token = ?", token).Find(&admin).Error
		if err == gorm.ErrRecordNotFound {
			ctx.ReplyError(common.StatusCodeDataNotExist, "信息不存在")
			return
		}

		ctx.Reply(admin)
	})
}
