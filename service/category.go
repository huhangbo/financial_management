package service

import (
	"financial_management/model"
	"financial_management/setting"
)

func GetCategoryByID(categoryID int) *model.Category {
	var (
		db          = setting.GetMySQL()
		tmpCategory = &model.Category{
			CategoryID: categoryID,
		}
	)
	if err := db.First(&tmpCategory); err != nil {
		return nil
	}
	return tmpCategory
}

func MGetCategory(categoryIDs []int) (map[int]*model.Category, error) {
	var (
		db           = setting.GetMySQL()
		categoryList []*model.Category
		categoryMap  = make(map[int]*model.Category)
	)
	if err := db.Find(&categoryList, categoryIDs).Error; err != nil {
		return nil, err
	}
	for _, category := range categoryList {
		categoryMap[category.CategoryID] = category
	}
	return categoryMap, nil
}

func GetCategoryList() ([]*model.Category, error) {
	var (
		db           = setting.GetMySQL()
		categoryList []*model.Category
	)
	result := db.Find(&categoryList)
	if result.Error != nil {
		return nil, result.Error
	}
	return categoryList, nil
}

func AddCategory(category *model.Category) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategoryByID(categoryID int) error {
	var (
		db          = setting.GetMySQL()
		tmpCategory = &model.Category{
			CategoryID: categoryID,
		}
	)
	if err := db.Delete(&tmpCategory).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCategoryInfo(category *model.Category) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Save(&category).Error; err != nil {
		return err
	}
	return nil
}
