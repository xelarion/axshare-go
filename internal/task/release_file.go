package task

import (
	"axshare_go/internal/utils"
	"errors"
	"os"
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
	if axurePath, err = utils.ExpandPath(os.Getenv("AXURE_PATH")); err != nil {
		return err
	}

	return utils.MkdirPath(axurePath)
}

func download(url string, fileName string) (webLink string, err error) {
	axureFileDir := genAxureFileDir(fileName)
	axureZipPath := filepath.Join(axureFileDir, fileName)

	if err = utils.RunCommand("wget", url, "-O", axureZipPath); err != nil {
		return "", err
	}

	if err = utils.RunCommand("unar", axureZipPath, "-o", axureFileDir); err != nil {
		return "", err
	}

	if err = utils.RunCommand("rm", axureZipPath); err != nil {
		return "", err
	}

	indexHtmlPath, err := filepath.Glob(filepath.Join(axureFileDir, "/*/index.html"))
	if err != nil {
		return "", err
	}

	if indexHtmlPath == nil {
		return "", errors.New("release failed")
	}

	webLink = strings.ReplaceAll(indexHtmlPath[0], axurePath, "")
	return webLink, nil
}

func genAxureFileDir(fileName string) string {
	fileDir := filepath.Join(axurePath, fileName)
	utils.MkdirPath(fileDir)
	return fileDir
}
