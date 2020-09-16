package models

import (
	"axshare_go/internal/db"
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
