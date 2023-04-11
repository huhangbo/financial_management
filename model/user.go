package model

import (
	"time"
)

type User struct {
	UserID    int       `json:"user_id" gorm:"primaryKey"`
	UserName  string    `json:"user_name"`
	Telephone int       `json:"telephone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
