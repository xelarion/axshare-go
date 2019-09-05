package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type Axure struct {
	gorm.Association
	ID           uint `gorm:"primary_key"`
	Name         string
	Link         string
	Desc         string
	AxureGroupId uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Uuid         string
	Attachments  []Attachment `gorm:"polymorphic:Reference;polymorphic_value:Axure"`
}

// 原型静态web链接
func (c *Axure) WebLink() string {
	if !c.IsReleased() {
		return ""
	}
	webHost := viper.GetString("web_host")
	return webHost + c.Link
}

// 原型永久地址
func (c *Axure) PermanentLink() string {
	adminHost := viper.GetString("admin_host")
	permanentLink := strings.Join([]string{
		adminHost, "/axures/", fmt.Sprint(c.ID), "?key=", c.Uuid}, "")
	return permanentLink
}

// 文件是否解压
func (c *Axure) IsReleased() bool {
	isReleased := len(c.Link) > 0
	return isReleased
}
