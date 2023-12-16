package aconfig

import (
	"fmt"
	"io/ioutil"
)

//获取配置文件内容
func GetConfig(filePath string) (content string, err error) {
	configBytes, err := GetConfigBytes(filePath)
	if err != nil {
		return "", err
	}
	return string(configBytes), nil
}

//获取配置文件内容
func GetConfigBytes(filePath string) (configBytes []byte, err error) {
	if filePath == "" {
		return configBytes, fmt.Errorf("文件路径不能为空")
	}

	_, configPath, err := GetConfigPath(filePath)
	if err != nil {
		return configBytes, err
	}

	configBytes, err = ioutil.ReadFile(configPath)
	if err != nil {
		return configBytes, err
	}
	return configBytes, nil
}
