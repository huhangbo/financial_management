package model

import "time"

type News struct {
	NewsID    int       `json:"news_id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
