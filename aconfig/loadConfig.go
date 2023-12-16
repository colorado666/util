package aconfig

import (
	"encoding/xml"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/configor"
	"path/filepath"
)

//加载配置文件
func LoadConfig(config interface{}, filePath string) (workPath string, configPath string, err error) {
	if config == nil {
		return workPath, configPath, fmt.Errorf("config不能为空")
	}
	if filePath == "" {
		return workPath, configPath, fmt.Errorf("filePath不能为空")
	}

	workPath, configPath, err = GetConfigPath(filePath)
	if err != nil {
		return workPath, configPath, err
	}
	ext := filepath.Ext(configPath)
	switch ext {
	case ".yml", ".yaml", ".json", "toml":

		err = configor.Load(config, configPath)
		if err != nil {
			return workPath, configPath, err
		}

	case ".conf":

		err = LoadIni(config, configPath)
		if err != nil {
			return workPath, configPath, err
		}
		err = configor.Load(config)
		if err != nil {
			return workPath, configPath, err
		}

	case ".xml":

		configBytes, err := GetConfigBytes(filePath)
		if err != nil {
			return workPath, configPath, err
		}
		err = xml.Unmarshal(configBytes, config)
		if err != nil {
			return workPath, configPath, err
		}

	default:
		return workPath, configPath, fmt.Errorf("暂不支持[%s]文件", ext)
	}

	return workPath, configPath, nil
}

//加载ini标签文件（.conf）
func LoadIni(obj interface{}, path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}
	err = cfg.MapTo(obj)
	if err != nil {
		return err
	}
	return nil
}
