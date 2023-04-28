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
		tmpBill = &model.Bill{
			BillID: billID,
		}
	)
	if err := db.First(&billID).Error; err != nil {
		return nil
	}
	return tmpBill
}

func GetBillByTime(billTypeList []model.BillType, beginTime, endTime time.Time, page int, limit int) ([]*model.Bill, error) {
	var (
		db       = setting.GetMySQL()
		billList []*model.Bill
	)
	offset := (page - 1) * limit
	if err := db.Where("bill_type IN ? AND creat_at Between ? AND ?", billTypeList, beginTime, endTime).Find(&billList).Limit(limit).Offset(offset).Error; err != nil {
		return nil, err
	}
	return billList, nil
}

func GetBillCount(billTypeList []model.BillType, beginTime, endTime time.Time) int {
	var (
		db    = setting.GetMySQL()
		count int64
	)
	if err := db.Model(&model.Bill{}).Where("bill_type IN ? AND creat_at Between ? AND ?", billTypeList, beginTime, endTime).Count(&count).Error; err != nil {
	}
	return int(count)
}
