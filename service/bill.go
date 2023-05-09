package service

import (
	"financial_management/model"
	"financial_management/setting"
	"time"
)

func AddBill(bill *model.Bill) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Create(&bill).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBill(billID []int) error {
	var (
		db       = setting.GetMySQL()
		billLIst []*model.Bill
	)
	if err := db.Delete(&billLIst, billID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBill(bill *model.Bill) error {
	var (
		db = setting.GetMySQL()
	)
	if err := db.Save(&bill).Error; err != nil {
		return err
	}
	return nil
}

func GetBillByIDs(billIDs []int) []*model.Bill {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
	)
	if err := db.Find(&billList, billIDs).Error; err != nil {
		return nil
	}
	return billList
}

func GetBillByID(billID int) *model.Bill {
	var (
		db      = setting.GetMySQL()
		tmpBill *model.Bill
	)
	if err := db.First(&tmpBill, billID).Error; err != nil {
		return nil
	}
	return tmpBill
}

func GetBillByMonth(billTypeList []model.BillType, userID int, year int, month int) ([]*model.Bill, error) {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
	)
	if err := db.Where("bill_type IN ? AND user_id = ? AND year = ? AND month = ?", billTypeList, userID, year, month).Find(&billList).Error; err != nil {
		return nil, err
	}
	return billList, nil
}

func GetBillByYear(billTypeList []model.BillType, userID int, year int) ([]*model.Bill, error) {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
	)
	if err := db.Where("bill_type IN ? AND user_id = ? AND year = ?", billTypeList, userID, year).Find(&billList).Error; err != nil {
		return nil, err
	}
	return billList, nil
}

func GetBillByTime(billTypeList []model.BillType, userID int, beginTime time.Time, endTime time.Time) ([]*model.Bill, error) {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
	)
	if err := db.Where("bill_type IN ? AND user_id = ? AND created_at BETWEEN ? AND ?", billTypeList, userID, beginTime, endTime).Find(&billList).Error; err != nil {
		return nil, err
	}
	return billList, nil
}

func GetBillAnalysis(billTypeList []model.BillType, userID int, year int, month int) ([]*model.Bill, error) {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
		lastList []*model.Bill
	)
	if err := db.Where("bill_type IN ? AND user_id = ? AND year = ? AND month = ?", billTypeList, userID, year, month).Find(&billList).Error; err != nil {
		return nil, err
	}
	if month == 0 {
		if err := db.Where("bill_type IN ? AND user_id = ? AND year = ? AND month = ?", billTypeList, userID, year-1, 11).Find(&billList).Error; err != nil {
			return nil, err
		}

	} else {
		if err := db.Where("bill_type IN ? AND user_id = ? AND year = ? AND month = ?", billTypeList, userID, year, month-1).Find(&lastList).Error; err != nil {
			return nil, err
		}
	}
	billList = append(billList, lastList...)
	return billList, nil
}

func GetUserBill(billTypeList []model.BillType, userID int) ([]*model.Bill, error) {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
	)
	if err := db.Where("bill_type IN ? AND user_id = ?", billTypeList, userID).Find(&billList).Error; err != nil {
		return nil, err
	}
	return billList, nil
}

func GetBillCount(billTypeList []model.BillType, userID int, year int, month int) int {
	var (
		db    = setting.GetMySQL()
		count int64
	)
	if err := db.Model(&model.Bill{}).Where("bill_type IN ? AND user_id = ? AND year = ? AND month = ?", billTypeList, userID, year, month).Count(&count).Error; err != nil {
	}
	return int(count)
}
