package task

import (
	"axshare_go/internal/utils"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var axurePath string

func deployAxure(url string, fileName string) (webLink string) {
	checkDirExist()
	webLink = download(url, fileName)
	return webLink
}

func checkDirExist() {
	axurePath, _ = utils.ExpandPath(viper.GetString("axure_path"))
	utils.MkdirPath(axurePath)
}

func download(url string, fileName string) (webLink string) {
	axureFileDir := genAxureFileDir(fileName)
	axureZipPath := filepath.Join(axureFileDir, fileName)

	_ = utils.RunCommand("wget", url, "-O", axureZipPath)
	_ = utils.RunCommand("unar", axureZipPath, "-o", axureFileDir)
	_ = utils.RunCommand("rm", axureZipPath)

	indexHtmlPath, err := filepath.Glob(filepath.Join(axureFileDir, "/*/index.html"))
	if err != nil || indexHtmlPath == nil {
		return ""
	}
	webLink = strings.ReplaceAll(indexHtmlPath[0], axurePath, "")
	return webLink
}

func genAxureFileDir(fileName string) string {
	fileDir := filepath.Join(axurePath, fileName)
	utils.MkdirPath(fileDir)
	return fileDir
}
