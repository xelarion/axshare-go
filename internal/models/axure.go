package models

import (
	"axshare_go/internal/db"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
)

type Axure struct {
	gorm.Model
	Name         string       `json:"name" xml:"name" binding:"required"`
	SecretKey    string       `json:"secret_key" binding:"required"`
	AxureGroupId uint         `gorm:"index"json:"axure_group_id"`
	Attachments  []Attachment `json:"attachments"`
	AxureGroup   AxureGroup
}

func (c *Axure) BeforeCreate() (err error) {
	c.genSecretKey()
	return
}

// 原型静态web链接
func (c *Axure) WebLink() string {
	attachment := c.LatestAttachment()
	return attachment.WebLink()
}

// 原型永久地址
func (c *Axure) PermanentLink() string {
	adminHost := os.Getenv("DASHBOARD_WEB_HOST")
	permanentLink := strings.Join([]string{
		adminHost, "/#/axures/", fmt.Sprint(c.ID), "?key=", c.SecretKey}, "")
	return permanentLink
}

// 文件是否解压
func (c *Axure) IsReleased() bool {
	attachment := c.LatestAttachment()
	return attachment.IsReleased()
}

func (c *Axure) LatestAttachment() Attachment {
	attachment := Attachment{}
	db.AxshareDb.Model(&attachment).Where("axure_id = ?", c.ID).Order("id desc").First(&attachment)
	return attachment
}

// generate secret key
func (c *Axure) genSecretKey() {
	c.SecretKey = uuid.NewV4().String()
}
