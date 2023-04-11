package main

import (
	"financial_management/handlers"
	"financial_management/middleware"
)

func initRouter() {
	h.Use(middleware.Cors())
	userGroup()
	adminGroup()
	categoryGroup()
	billGroup()
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
	group.Use(middleware.Auth())
	{
		group.POST("/login", handlers.AdminLogin)
		group.GET("/user/get", handlers.AdminGetUser)
		group.POST("/user/delete", handlers.AdminDeleteUser)
		group.POST("/user/add", handlers.AdminAddUser)
	}
}

func categoryGroup() {
	group := h.Group("/category")
	group.Use(middleware.Auth())
	{
		group.POST("/add", handlers.CategoryAdd)
		group.GET("/list", handlers.CategoryGetList)
		group.POST("/update", handlers.CategoryUpdate)
		group.POST("/delete", handlers.CategoryDelete)
	}
}

func billGroup() {
	group := h.Group("/bill")
	group.Use(middleware.Auth())
	{
		group.GET("/get", handlers.BillGet)
		group.POST("/add", handlers.BillRecord)
		group.POST("/delete", handlers.BillDelete)
		group.POST("/update", handlers.BillUpdate)
	}
}
