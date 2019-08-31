package models

import (
	"time"
)

type AxureGroup struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
