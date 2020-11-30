package router

import (
	"go-wsadmin/handlers"
	"go-wsadmin/middlewares"
	"go-wsadmin/pkg/win"
)

func InitRouter(s *win.Server) {

	api := s.Group("api/v1")

	api.AddHandler("login", handlers.Login)

	// must login
	g := api.Group("", middlewares.LoginRequired)

	g.AddHandler("errorCode", handlers.ErrorCode)
	g.AddHandler("logout", handlers.Logout)
	g.AddHandler("getAuthInfo", handlers.GetAuthInfo)

	g.AddHandler("admin/list", handlers.AdminList)
	g.AddHandler("admin/create", handlers.AdminCreate)
	g.AddHandler("admin/update", handlers.AdminCreate)
	g.AddHandler("admin/delete", handlers.AdminCreate)

	g.Group("/user")

}
