package models

import "time"

type AxureGroup struct {
	ID        uint `json:"id"gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
}
