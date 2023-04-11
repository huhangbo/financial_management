package model

import "time"

type News struct {
	NewsID    int       `json:"news_id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
