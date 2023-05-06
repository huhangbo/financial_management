package model

import "time"

type Budget struct {
	BudgetId  int       `json:"budget_id" gorm:"primaryKey"`
	Fee       int       `json:"fee"`
	UserID    int       `json:"user_id"`
	Year      int       `json:"year"`
	Mouth     int       `json:"mouth"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
