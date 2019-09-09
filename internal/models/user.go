package models

import (
	"axshare_go/internal/db"
	"time"
)

type User struct {
	ID        uint      `json:"created_at" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
}

// todo
func (c *User) DestroyToken() error {
	return nil
}

func FindUserbyToken(token string) User {
	user := User{}
	// todo
	db.AxshareDb.First(&user)
	return user
}
