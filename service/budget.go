package service

import (
	"financial_management/model"
	"financial_management/setting"
)

func AddBudget(budget *model.Budget) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Create(&budget).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBudget(budget *model.Budget) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Save(&budget).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBudget(budgetID int) error {
	var (
		db        = setting.GetMySQL()
		tmpBudget = &model.Budget{
			BudgetId: budgetID,
		}
	)
	if err := db.Delete(tmpBudget).Error; err != nil {
		return err
	}
	return nil

}

func GetBudgetByID(budgetID int) *model.Budget {
	var (
		db        = setting.GetMySQL()
		tmpBudget = &model.Budget{
			BudgetId: budgetID,
		}
	)
	if err := db.First(&tmpBudget).Error; err != nil {
		return nil
	}
	return tmpBudget
}

func GetBudgetByUserID(userID, year, month int) *model.Budget {
	var (
		db        = setting.GetMySQL()
		tmpBudget = &model.Budget{}
	)
	if err := db.Where(&model.Budget{
		UserID: userID,
		Year:   year,
		Month:  month,
	},
	).First(&tmpBudget).Error; err != nil {
		return nil
	}
	return tmpBudget
}
