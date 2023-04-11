package service

import (
	"financial_management/model"
	"financial_management/setting"
)

func MGetUserByID(userIDs []int) []*model.User {
	var (
		db       = setting.GetMySQL()
		userList []*model.User
	)
	db.Find(&userList)
	return userList
}

func GetUserByID(userID int) *model.User {
	var (
		db   = setting.GetMySQL()
		user *model.User
	)
	if err := db.First(&user, userID).Error; err != nil {
		return user
	}
	return nil
}

func IsUserExist(userID int) bool {
	var (
		db   = setting.GetMySQL()
		user *model.User
	)
	if err := db.First(&user, userID).Error; err != nil {
		return false
	}
	return true
}

func InsertUser(user *model.User) error {
	var (
		db = setting.GetMySQL()
	)
	return db.Create(user).Error
}
