package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/aconfig"
	"gitee.com/asktop_golib/util/ajson"
	"path/filepath"
	"testing"
)

func TestGo(t *testing.T) {
	fmt.Println(filepath.Ext("app.conf"))
}

func TestGetConfigPath(t *testing.T) {
	fmt.Println(aconfig.GetConfigPath("conf/app.conf"))
}

func TestGetConfigContent(t *testing.T) {
	fmt.Println(aconfig.GetConfig("conf/app.conf"))
}

func TestLoadConfig(t *testing.T) {
	config := new(app)
	_, _, err := aconfig.LoadConfig(config, "conf/app.conf")
	//_, _, err := aconfig.LoadConfig(config, "conf/app.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ajson.Encode(config))
}

//项目配置文件
type app struct {
	AppName  string              `json:"appname" yaml:"appname" ini:"appname"`                        //项目名
	RunMode  string              `json:"runmode" yaml:"runmode" ini:"runmode"`                        //运行模式：[prod:(默认)生产模式，对模板进行编译缓存； dev:开发模式，对模板不做编译缓存，每次修改自动编译]
	HttpPort int                 `json:"httpport" yaml:"httpport" ini:"httpport"`                     //beego服务端口
	Username []string            `json:"username" yaml:"username" ini:"username" default:"[003,004]"` //
	Mysql    aconfig.MysqlConfig `json:"mysql" yaml:"mysql" ini:"mysql"`                              //mysql配置
}
