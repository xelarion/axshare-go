package models

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Attachment struct {
	gorm.Model
	Desc   string `json:"desc"`
	Link   string `json:"link"`
	UserId uint   `gorm:"index"json:"user_id"`
	User   User   `json:"user"`
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
	webHost := viper.GetString("web_host")
	return webHost + c.Link
}

// 原型压缩包下载地址
func (c *Attachment) DownloadUrl() string {
	return ""
}
