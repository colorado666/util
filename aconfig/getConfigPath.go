package aconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//通过查找config文件，获取工作路径和配置文件路径
//@param filePath 相对路径：自动从workDir往上级目录逐级查找；绝对路径：不会往上级目录逐级查找
func GetConfigPath(filePath string) (workPath string, configPath string, err error) {
	var isExist bool
	if _, err := os.Stat(filePath); err == nil {
		isExist = true
	} else {
		isExist = os.IsExist(err)
	}
	if isExist {
		workPath, _ = os.Getwd() //workDir
		if filepath.IsAbs(filePath) {
			configPath = filePath
		} else {
			configPath = filepath.Join(workPath, filePath)
		}
		return
	} else {
		workPath, _ = os.Getwd() //workDir
		configBasePath, ok := findConfigPath(workPath, filePath)
		if ok {
			workPath = configBasePath
			configPath = filepath.Join(configBasePath, filePath)
			return
		} else {
			workPath, _ = filepath.Abs(filepath.Dir(os.Args[0])) //outputDir
			configBasePath, ok = findConfigPath(workPath, filePath)
			if ok {
				workPath = configBasePath
				configPath = filepath.Join(configBasePath, filePath)
				return
			} else {
				err = fmt.Errorf("找不到文件[%s]", filePath)
				return "", "", err
			}
		}
	}
}

//查找配置文件
func findConfigPath(configBasePath string, configFilePath string) (string, bool) {
	var okSpace, notSpace string
	if runtime.GOOS == "windows" {
		okSpace = `\`
		notSpace = `/`
	} else {
		okSpace = `/`
		notSpace = `\`
	}
	var configFile string
	if configBasePath == `` || strings.HasSuffix(configBasePath, `:`) {
		configBasePath = configBasePath + okSpace
		configFile = configBasePath + configFilePath
	} else {
		configFile = filepath.Join(configBasePath, configFilePath)
	}
	configFile = strings.Replace(configFile, notSpace, okSpace, -1)
	//fmt.Println(configFile)
	var isExist bool
	if _, err := os.Stat(configFile); err == nil {
		isExist = true
	} else {
		isExist = os.IsExist(err)
	}
	if isExist {
		return configBasePath, true
	} else {
		if configBasePath == `` || configBasePath == okSpace || configBasePath == notSpace || strings.HasSuffix(configBasePath, `:\`) || strings.HasSuffix(configBasePath, `:/`) {
			return configBasePath, false
		} else {
			configBasePath = configBasePath[:strings.LastIndex(configBasePath, okSpace)]
			return findConfigPath(configBasePath, configFilePath)
		}
	}
}
