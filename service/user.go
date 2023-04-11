package service

import (
	"financial_management/model"
	"financial_management/setting"
)

func UpdateUser(user *model.User) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(userIDs []int) error {
	var (
		db       = setting.GetMySQL()
		UserList []*model.User
	)
	if err := db.Delete(&UserList, userIDs).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUser() []*model.User {
	var (
		db       = setting.GetMySQL()
		userList []*model.User
	)
	if err := db.Find(&userList).Error; err != nil {
		return nil
	}
	return userList
}

func GetUserByTelephone(telephone int) *model.User {
	var (
		db   = setting.GetMySQL()
		user *model.User
	)
	if err := db.Where("telephone = ?", telephone).First(&user); err != nil {
		return nil
	}
	return user
}

func GetUserByID(userID int) *model.User {
	var (
		db   = setting.GetMySQL()
		user = &model.User{
			UserID: userID,
		}
	)
	if err := db.First(&user); err != nil {
		return nil
	}
	return user
}

func IsUserExist(telephone int) bool {
	var (
		db   = setting.GetMySQL()
		user *model.User
	)
	if err := db.Where("telephone = ?", telephone).First(&user); err != nil {
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
