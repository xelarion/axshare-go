package task

import (
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

func deployAxure(url string, fileName string) (webLink string, err error) {
	fileReleaseDir, err := models.CacheConfig.MkdirFileReleaseDir()
	if err != nil {
		return webLink, err
	}

	// 此 fileName 所在目录
	axureFileDir := ""
	if webLink, axureFileDir, err = download(url, fileReleaseDir, fileName); err != nil {
		// 下载解压失败，则删除该目录
		if len(axureFileDir) > 0 {
			logrus.Warning("下载解压附件失败，清除文件夹: ", axureFileDir)
			_ = utils.RunCommand("rm", "-rf", axureFileDir)
		}
	}

	return webLink, err
}

func download(url, fileReleaseDir, fileName string) (webLink, axureFileDir string, err error) {
	// 此 fileName 所在目录
	axureFileDir = genAxureFileDir(fileReleaseDir, fileName)
	// axureZipPath 此 fileName 下载下来的压缩包路径
	axureZipPath := filepath.Join(axureFileDir, fileName)

	if err = utils.RunCommand("wget", url, "-O", axureZipPath); err != nil {
		errMsg := fmt.Sprintf("执行 wget 命令发生异常, error: , %s", err.Error())
		return "", axureFileDir, errors.New(errMsg)
	}

	if err = utils.RunCommand("unar", axureZipPath, "-o", axureFileDir); err != nil {
		errMsg := fmt.Sprintf("执行 unar 解压命令发生异常, error: , %s", err.Error())
		return "", axureFileDir, errors.New(errMsg)
	}

	if err = utils.RunCommand("rm", axureZipPath); err != nil {
		errMsg := fmt.Sprintf("执行 rm 命令发生异常, error: , %s", err.Error())
		return "", axureFileDir, errors.New(errMsg)
	}

	indexHtmlPath, err := filepath.Glob(filepath.Join(axureFileDir, "/*/index.html"))
	if err != nil {
		errMsg := fmt.Sprintf("请检查上传的文件是否包含 index.html 文件, error: , %s", err.Error())
		return "", axureFileDir, errors.New(errMsg)
	}

	if indexHtmlPath == nil {
		errMsg := fmt.Sprintf("请检查上传的文件是否正确！需包含 index.html 文件")
		return "", axureFileDir, errors.New(errMsg)
	}

	webLink = strings.ReplaceAll(indexHtmlPath[0], fileReleaseDir, "")
	return webLink, axureFileDir, nil
}

func genAxureFileDir(fileReleaseDir, fileName string) string {
	fileDir := filepath.Join(fileReleaseDir, fileName)
	_ = utils.MkdirPath(fileDir)
	return fileDir
}
