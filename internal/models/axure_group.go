package models

import "github.com/jinzhu/gorm"

type AxureGroup struct {
	gorm.Model
	Name string
	Desc string
}
