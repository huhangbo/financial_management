package financial_management

import (
	"financial_management/handlers"
	"financial_management/middleware"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
)

func initRouter() {
	h.Use(recovery.Recovery(), middleware.Cors)
	userGroup()
	adminGroup()
}

func userGroup() {
	group := h.Group("/user")
	{
		group.POST("/register", handlers.UserRegister)
		group.POST("/login", handlers.UserLogin)
		group.POST("/update/password")
		group.POST("/update/info", handlers.UserUpdateInfo)
	}
}

func adminGroup() {
	group := h.Group("/admin")
	{
		group.POST("/login", handlers.AdminLogin)
		group.GET("/user/get", handlers.AdminGetUser)
		group.POST("/user/delete", handlers.AdminDeleteUser)
		group.POST("/user/add", handlers.AdminAddUser)
	}
}
