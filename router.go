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

	budgetGroup()

	notesGroup()

	newsGroup()
}

func userGroup() {
	group := h.Group("/user")
	{
		group.GET("/info", handlers.UserInfoGet)
		group.POST("/register", handlers.UserRegister)
		group.POST("/login", handlers.UserLogin)

	}

	update := group.Group("/update").Use(middleware.Auth())
	{
		update.POST("/password", handlers.UpdatePassword)
		update.POST("/info", handlers.UserUpdateInfo)
	}
}

func adminGroup() {
	group := h.Group("/admin")
	{
		group.POST("/login", handlers.AdminLogin)
	}

	user := group.Group("/user").Use(middleware.Auth())
	{
		user.GET("/get", handlers.AdminGetUser)
		user.POST("/delete", handlers.AdminDeleteUser)
		user.POST("/add", handlers.AdminAddUser)
		user.POST("/update", handlers.UserUpdateInfo)
		user.POST("/password", handlers.UpdatePassword)
	}
}

func categoryGroup() {
	group := h.Group("/category")
	group.Use(middleware.Auth())
	{
		group.GET("/list", handlers.CategoryGetList)
		group.POST("/add", handlers.CategoryAdd)
		group.POST("/update", handlers.CategoryUpdate)
		group.POST("/delete", handlers.CategoryDelete)
	}
}

func billGroup() {
	group := h.Group("/bill")
	group.Use(middleware.Auth())
	{
		group.POST("/list", handlers.BillGet)
		group.POST("/add", handlers.BillRecord)
		group.POST("/delete", handlers.BillDelete)
		group.POST("/update", handlers.BillUpdate)
	}
}

func budgetGroup() {
	group := h.Group("/budget")
	group.Use(middleware.Auth())
	{
		group.POST("/get", handlers.BudgetGet)
		group.POST("/add", handlers.BudgetAdd)
		group.POST("/update", handlers.BudgetUpdate)
		group.POST("/delete", handlers.BudgetDelete)
	}
}

func notesGroup() {
	group := h.Group("/notes")
	group.Use(middleware.Auth())
	{
		group.GET("/list", handlers.NotesGet)
		group.POST("/add", handlers.NotesAdd)
		group.POST("/update", handlers.NotesUpdate)
		group.POST("/delete", handlers.NotesDelete)
	}
}

func newsGroup() {
	group := h.Group("/news")
	group.Use(middleware.Auth())
	{
		group.GET("/list", handlers.NewsGet)
		group.POST("/add", handlers.NewsAdd)
		group.POST("/update", handlers.NewsUpdate)
		group.POST("/delete", handlers.NewsDelete)
	}
}
