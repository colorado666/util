package aconfig

import (
	"fmt"
	"gitee.com/asktop_golib/util/afile"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	AppPath  string //项目根目录绝对路径
	ConfPath string //项目配置文件绝对路径
)

//加载项目配置文件
func InitAppConfig(config interface{}, paths ...string) (workPath string, configPath string, err error) {
	filePath := filepath.Clean(filepath.Join(paths...))

	workPath, configPath, err = LoadConfig(config, filePath)
	if err != nil {
		return workPath, configPath, err
	}

	if AppPath == "" {
		AppPath = workPath
		ConfPath = configPath
	}
	return workPath, configPath, nil
}

// 获取本地路径（不创建文件夹）
func GetAppPath(paths ...string) string {
	appPath := AppPath
	if appPath == "" {
		appPath, _ = os.Getwd()
	}
	if len(paths) == 0 {
		// 获取项目根目录
		return appPath
	}
	localPath := filepath.Clean(filepath.Join(paths...))
	if !filepath.IsAbs(localPath) {
		localPath = filepath.Join(appPath, localPath)
	}
	return localPath
}

// 获取本地路径（创建文件夹）
func GetAppLocalPath(paths ...string) string {
	localPath := GetAppPath(paths...)
	tempLocalPath := localPath
	if filepath.Ext(localPath) != "" {
		tempLocalPath, _ = filepath.Split(localPath)
	}
	err := afile.CreateDir(tempLocalPath)
	if err == nil {
		return localPath
	}
	return GetAppPath()
}

// 获取日志存放文件夹和文件绝对路径
func GetAppLogPath(configLogPath string, appname string, driver string) string {
	configLogPath = filepath.Clean(configLogPath)
	if configLogPath == "" || (!filepath.IsAbs(configLogPath) && !strings.HasSuffix(configLogPath, "/log") && !strings.HasSuffix(configLogPath, `\log`)) {
		configLogPath = GetAppLocalPath("log", configLogPath)
	}
	if filepath.Ext(configLogPath) != ".log" && filepath.Ext(configLogPath) != ".text" {
		names := []string{}
		if appname == "" {
			appname = "app"
		}
		if driver == "" {
			names = append(names, appname, "log")
		} else {
			names = append(names, appname, driver, "log")
		}
		logName := strings.ReplaceAll(strings.Join(names, "."), " ", "")
		configLogPath = filepath.Join(configLogPath, logName)
	}
	return configLogPath
}

// 获取文件存放绝对路径
func GetAppUploadPath(configUploadPath string, paths ...string) string {
	localPath := filepath.Clean(filepath.Join(paths...))
	if filepath.IsAbs(localPath) {
		return GetAppLocalPath(localPath)
	} else {
		if configUploadPath == "" || (!filepath.IsAbs(configUploadPath) && !strings.HasSuffix(configUploadPath, "/upload") && !strings.HasSuffix(configUploadPath, `\upload`)) {
			configUploadPath = GetAppLocalPath("upload")
		}
		return GetAppLocalPath(configUploadPath, localPath)
	}
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//通用配置
//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

//mysql配置
type MysqlConfig struct {
	Host              string                    `json:"host" yaml:"host" toml:"host" ini:"host" default:"127.0.0.1"`
	Port              int                       `json:"port" yaml:"port" toml:"port" ini:"port" default:"3306"`
	Username          string                    `json:"username" yaml:"username" toml:"username" ini:"username"`
	Password          string                    `json:"password" yaml:"password" toml:"password" ini:"password"`
	Database          string                    `json:"database" yaml:"database" toml:"database" ini:"database"`
	Charset           string                    `json:"charset" yaml:"charset" toml:"charset" ini:"charset"`
	Timezone          string                    `json:"timezone" yaml:"timezone" toml:"timezone" ini:"timezone"`
	AllowOldPasswords int                       `json:"allowoldpasswords" yaml:"allowoldpasswords" toml:"allowoldpasswords" ini:"allowoldpasswords"`
	MaxIdleConns      int                       `json:"maxidleconns" yaml:"maxidleconns" toml:"maxidleconns" ini:"maxidleconns"`
	MaxOpenConns      int                       `json:"maxopenconns" yaml:"maxopenconns" toml:"maxopenconns" ini:"maxopenconns"`
	ConnMaxLifetime   int                       `json:"connmaxlifetime" yaml:"connmaxlifetime" toml:"connmaxlifetime" ini:"connmaxlifetime" `
	SqlLogLevel       int                       `json:"sqlloglevel" yaml:"sqlloglevel" toml:"sqlloglevel" ini:"sqlloglevel" default:"1"` //SQL日志打印级别：0：不打印SQL；1：只打印err；2：打印全部
	LogFunc           func(args ...interface{}) `json:"-"`                                                                               //SQL日志打印方法
}

//获取mysql连接配置
func (c *MysqlConfig) GetDsn(withoutDb ...bool) string {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == 0 {
		c.Port = 3306
	}
	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}
	if c.AllowOldPasswords == 0 {
		c.AllowOldPasswords = 1
	}
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 300
	}
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 500
	}
	if c.ConnMaxLifetime == 0 {
		c.ConnMaxLifetime = 180
	}
	database := c.Database
	if len(withoutDb) > 0 && withoutDb[0] {
		database = ""
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&allowOldPasswords=%d",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		database,
		c.Charset,
		c.AllowOldPasswords,
	)
	if c.Timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(c.Timezone)
	}
	return dsn
}

//redis配置
type RedisConfig struct {
	Host        string `json:"host" yaml:"host" toml:"host" ini:"host" default:"127.0.0.1"`
	Port        string `json:"port" yaml:"port" toml:"port" ini:"port" default:"6379"`
	Password    string `json:"password" yaml:"password" toml:"password" ini:"password"`
	Select      int    `json:"select" yaml:"select" toml:"select" ini:"select"`
	MaxIdle     int    `json:"maxidle" yaml:"maxidle" toml:"maxidle" ini:"maxidle"`
	MaxActive   int    `json:"maxactive" yaml:"maxactive" toml:"maxactive" ini:"maxactive"`
	IdleTimeout int    `json:"idletimeout" yaml:"idletimeout" toml:"idletimeout" ini:"idletimeout"`
}

func (c *RedisConfig) GetBeegoSessionDsn() string {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "6379"
	}
	if c.MaxActive == 0 {
		c.MaxActive = 500
	}
	if c.MaxIdle == 0 {
		c.MaxIdle = 300
	}
	if c.IdleTimeout == 0 {
		c.IdleTimeout = 180
	}
	//host:port,MaxIdle,password,dbnum,IdleTimeout
	redisAddr := net.JoinHostPort(c.Host, c.Port)
	redisMaxidle := strconv.Itoa(c.MaxIdle)
	redisPassword := c.Password
	redisDb := strconv.Itoa(c.Select)
	return strings.Join([]string{redisAddr, redisMaxidle, redisPassword, redisDb}, ",")
}
