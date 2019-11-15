package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"json:"email"`
	Nickname string `gorm:"type:varchar(100);unique_index"json:"nickname"`
	Username string `gorm:"type:varchar(100);unique_index"json:"username"`
	Password string `gorm:"column:encrypted_password" json:"password"`
	Avatar   string `json:"avatar"`
}
