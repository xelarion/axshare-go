package app

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"os"
)

func GenUpToken() string {
	var accessKey = os.Getenv("QINIU_ACCESS_KEY")
	var secretKey = os.Getenv("QINIU_SECRET_KEY")
	var bucket = os.Getenv("QINIU_BUCKET")
	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	putPolicy.Expires = 7200
	return upToken
}
