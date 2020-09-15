package task

import (
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
)

var axurePath string
var checkDeployDirLock = sync.RWMutex{}

func deployAxure(url string, fileName string) (webLink string, err error) {
	if err := checkDirExist(); err != nil {
		return "", err
	}

	return download(url, fileName)
}

func checkDirExist() error {
	checkDeployDirLock.Lock()
	defer checkDeployDirLock.Unlock()

	var err error
	if axurePath, err = utils.ExpandPath(models.CacheConfig.FileReleaseDir); err != nil {
		errMsg := fmt.Sprintf("请检查原型解压文件夹配置是否正确 '%s', error: , %s", models.CacheConfig.FileReleaseDir, err.Error())
		return errors.New(errMsg)
	}

	return utils.MkdirPath(axurePath)
}

func download(url string, fileName string) (webLink string, err error) {
	axureFileDir := genAxureFileDir(fileName)
	axureZipPath := filepath.Join(axureFileDir, fileName)

	if err = utils.RunCommand("wget", url, "-O", axureZipPath); err != nil {
		errMsg := fmt.Sprintf("执行 wget 命令发生异常, error: , %s", err.Error())
		return "", errors.New(errMsg)
	}

	if err = utils.RunCommand("unar", axureZipPath, "-o", axureFileDir); err != nil {
		errMsg := fmt.Sprintf("执行 unar 解压命令发生异常, error: , %s", err.Error())
		return "", errors.New(errMsg)
	}

	if err = utils.RunCommand("rm", axureZipPath); err != nil {
		errMsg := fmt.Sprintf("执行 rm 命令发生异常, error: , %s", err.Error())
		return "", errors.New(errMsg)
	}

	indexHtmlPath, err := filepath.Glob(filepath.Join(axureFileDir, "/*/index.html"))
	if err != nil {
		errMsg := fmt.Sprintf("请检查上传的文件是否包含 index.html 文件, error: , %s", err.Error())
		return "", errors.New(errMsg)
	}

	if indexHtmlPath == nil {
		errMsg := fmt.Sprintf("请检查上传的文件是否正确！需包含 index.html 文件")
		return "", errors.New(errMsg)
	}

	webLink = strings.ReplaceAll(indexHtmlPath[0], axurePath, "")
	return webLink, nil
}

func genAxureFileDir(fileName string) string {
	fileDir := filepath.Join(axurePath, fileName)
	_ = utils.MkdirPath(fileDir)
	return fileDir
}
