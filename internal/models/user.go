package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Nickname  string
}
