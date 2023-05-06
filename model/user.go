package model

import (
	"time"
)

type User struct {
	UserID    int       `json:"user_id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Gender    int       `json:"gender"`
	Email     string    `json:"email"`
	Telephone int       `json:"telephone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
