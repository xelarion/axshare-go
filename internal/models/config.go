package models

import (
	"axshare_go/internal/db"
	"axshare_go/internal/utils"
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"sync"
	"time"
)

type Config struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsValid   bool `json:"is_valid"`

	PublicConfig

	FileReleaseDir string `json:"file_release_dir"`
	WebDomain      string `json:"web_domain"`

	QiniuAccessKey    string `json:"qiniu_access_key"`
	QiniuSecretKey    string `json:"qiniu_secret_key"`
	QiniuBucket       string `json:"qiniu_bucket"`
	QiniuBucketDomain string `json:"qiniu_bucket_domain"`
	QiniuUploadUrl    string `json:"qiniu_upload_url"`
}

type PublicConfig struct {
	SiteName      string `json:"site_name"`
	ICPRecordNo   string `json:"icp_record_no"`
	ICPRecordLink string `json:"icp_record_link"`
	Copyright     string `json:"copyright"`
}

var CacheConfig = Config{}
var CacheConfigLock = sync.RWMutex{}
var CheckReleaseDirLock = sync.RWMutex{}

func InitCacheConfig() {
	CacheConfigLock.Lock()
	defer CacheConfigLock.Unlock()

	db.AxshareDb.Model(&Config{}).
		Where("is_valid = TRUE").
		Limit(1).
		Find(&CacheConfig)

	if !CacheConfig.IsValid {
		initDefaultConfig()
	}
}

func GenerateUploadToken() string {
	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: CacheConfig.QiniuBucket,
	}
	mac := qbox.NewMac(CacheConfig.QiniuAccessKey, CacheConfig.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	putPolicy.Expires = 7200
	return upToken
}

func initDefaultConfig() {
	CacheConfig.SiteName = "Axshare"
	CacheConfig.FileReleaseDir = "~/web/axure/"
	CacheConfig.WebDomain = "https://axshare.you-domain.com"
	CacheConfig.QiniuUploadUrl = "https://up-z2.qiniup.com"
}

// MkdirFileReleaseDir 生成 原型解压文件夹
func (c *Config) MkdirFileReleaseDir() (fileReleaseDir string, err error) {
	CheckReleaseDirLock.Lock()
	defer CheckReleaseDirLock.Unlock()

	if fileReleaseDir, err = utils.ExpandPath(CacheConfig.FileReleaseDir); err != nil {
		msg := fmt.Sprintf("请检查原型解压文件夹配置是否正确 '%s', error: , %s", CacheConfig.FileReleaseDir, err.Error())
		return fileReleaseDir, errors.New(msg)
	}

	return fileReleaseDir, utils.MkdirPath(fileReleaseDir)
}

// FileReleaseAbsDir 获取 原型解压文件夹
func (c *Config) FileReleaseAbsDir() (absDir string, err error) {
	if absDir, err = utils.ExpandPath(CacheConfig.FileReleaseDir); err != nil {
		msg := fmt.Sprintf("请检查原型解压文件夹配置是否正确 '%s', error: , %s", CacheConfig.FileReleaseDir, err.Error())
		return absDir, errors.New(msg)
	}

	return absDir, nil
}
