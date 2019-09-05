package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"created_at" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
}
