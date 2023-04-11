package model

import "time"

type Notes struct {
	NotesID   int       `json:"notes_id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
