package model

import "time"

type Notes struct {
	NotesID   int       `json:"notes_id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Place     string    `json:"place"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
