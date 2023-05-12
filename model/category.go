package model

import "time"

type Category struct {
	CategoryID     int       `json:"category_id" gorm:"primaryKey"`
	BillType       BillType  `json:"bill_type"`
	CategoryDetail string    `json:"category_detail"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
