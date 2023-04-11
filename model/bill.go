package model

import "time"

type BillType int8

const (
	Expend BillType = 1
	Income BillType = 2
)

type Bill struct {
	BillID     int       `json:"bill_id" gorm:"primaryKey"`
	BillType   BillType  `json:"bill_type"`
	UserID     int       `json:"user_id"`
	Fee        int       `json:"fee"`
	CategoryID int       `json:"category_id"`
	Remark     string    `json:"remark"`
	Category   *Category `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
