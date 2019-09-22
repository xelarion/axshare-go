package tasks

import (
	"axshare_go/internal/utils"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var AxurePath string

func DeployAxure(url string, fileName string) (webLink string) {
	checkDirExist()
	webLink = download(url, fileName)
	return webLink
}

func checkDirExist() {
	AxurePath, _ = utils.ExpandPath(viper.GetString("axure_path"))
	utils.MkdirPath(AxurePath)
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
	webLink = strings.ReplaceAll(indexHtmlPath[0], AxurePath, "")
	return webLink
}

func genAxureFileDir(fileName string) string {
	fileDir := filepath.Join(AxurePath, fileName)
	utils.MkdirPath(fileDir)
	return fileDir
}
