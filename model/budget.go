package model

import "time"

type Budget struct {
	BudgetId  int `json:"budget_id"`
	Fee       int `json:"fee"`
	UserID    int `json:"user_id"`
	Year      int `json:"year"`
	Mouth     int `json:"mouth"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
