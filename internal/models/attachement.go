package models

import (
	"github.com/spf13/viper"
	"time"
)

type Attachment struct {
	ID            uint `gorm:"primary_key"`
	ReferenceId   uint
	ReferenceType string
	key           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserId        uint
	User          User
	Link          string
	Desc          string
}

// 文件是否解压
func (c *Attachment) IsReleased() bool {
	isReleased := len(c.Link) > 0
	return isReleased
}

// 原型静态web链接
func (c *Attachment) WebLink() string {
	if !c.IsReleased() {
		return ""
	}
	if c.ReferenceType == "Axure" {
		webHost := viper.GetString("web_host")
		return webHost + c.Link
	} else {
		return ""
	}
}

// 原型压缩包下载地址
func (c *Attachment) DownloadUrl() string {
	if c.ReferenceType == "Axure" {
		return ""
	} else {
		return ""
	}
}
