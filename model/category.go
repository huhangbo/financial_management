package model

import "time"

type Category struct {
	CategoryID     int       `json:"category_id" gorm:"primaryKey"`
	UserID         int       `json:"user_id"`
	CategoryDetail string    `json:"category_detail"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
