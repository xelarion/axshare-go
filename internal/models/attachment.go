package models

import (
	"axshare_go/internal/utils"
	"github.com/xandercheung/acct"
	"gorm.io/gorm"
	"strings"
)

type Attachment struct {
	gorm.Model
	Desc          string                  `json:"desc" xml:"desc" binding:"required"`
	Link          string                  `json:"link"`
	FileHash      string                  `json:"file_hash"`
	ReleaseStatus AttachmentReleaseStatus `json:"release_status"`
	ReleaseError  string                  `json:"release_error"`
	AxureId       uint                    `json:"axure_id"`
	Axure         Axure                   `json:"axure" gorm:"foreignKey:AxureId"`
	AccountId     uint                    `json:"account_id"`
	Account       acct.Account            `json:"user" gorm:"foreignKey:AccountId"`
}

func (c *Attachment) GenFileName() string {
	axure := FindAxure(c.AxureId)
	fileName := strings.Join([]string{
		utils.FormatUint(axure.AxureGroupId),
		utils.Strftime(c.CreatedAt, "20060102150405"),
		c.FileHash,
	}, "_")

	return fileName
}

// 文件是否解压
func (c *Attachment) IsReleased() bool {
	return c.ReleaseStatus == AttachmentReleaseStatusSuccessful
}

// 文件是否上传
func (c *Attachment) IsFileUploaded() bool {
	return len(c.FileHash) > 0
}

// 原型静态web链接
func (c *Attachment) WebLink() string {
	if !c.IsReleased() {
		return ""
	}

	return CacheConfig.WebDomain + c.Link
}

// 原型压缩包下载地址
func (c *Attachment) DownloadUrl() string {
	return CacheConfig.QiniuBucketDomain + "/" + c.FileHash
}
