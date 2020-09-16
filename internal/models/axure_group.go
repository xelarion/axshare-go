package models

import (
	"gorm.io/gorm"
	"time"
)

type AxureGroup struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name      string         `json:"name"`
	Desc      string         `json:"desc"`
}
