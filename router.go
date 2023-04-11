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
		group.POST("/register", handlers.UserRegister)
		group.POST("/login", handlers.UserLogin)
		group.POST("/update/password").Use(middleware.Auth())
		group.POST("/update/info", handlers.UserUpdateInfo).Use(middleware.Auth())
	}
}

func adminGroup() {
	group := h.Group("/admin")
	{
		group.POST("/login", handlers.AdminLogin)
		group.GET("/user/get", handlers.AdminGetUser).Use(middleware.Auth())
		group.POST("/user/delete", handlers.AdminDeleteUser).Use(middleware.Auth())
		group.POST("/user/add", handlers.AdminAddUser).Use(middleware.Auth())
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
		group.POST("/get", handlers.BillGet)
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
		group.POST("/get", handlers.NotesGet)
		group.POST("/add", handlers.NotesAdd)
		group.POST("/update", handlers.NotesUpdate)
		group.POST("/delete", handlers.NotesDelete)
	}
}

func newsGroup() {
	group := h.Group("/news")
	group.Use(middleware.Auth())
	{
		group.POST("/get", handlers.NewsGet)
		group.POST("/add", handlers.NewsAdd)
		group.POST("/update", handlers.NewsUpdate)
		group.POST("/delete", handlers.NewsDelete)
	}
}
