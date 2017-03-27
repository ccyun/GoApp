package conf

import (
	"encoding/json"

	"os"

	"github.com/astaxie/beego/config"
)

//Conf 配置
var Conf config.Configer

//InitConfig 初始化配置
func InitConfig() error {
	confile := "conf.ini"
	if len(os.Args) > 2 && os.Args[1] == "-c" {
		confile = os.Args[2]
	}
	conf, err := config.NewConfig("ini", confile)
	if err != nil {
		return err
	}
	Conf = conf
	return nil
}

//String 读取配置
func String(name string) string {
	return Conf.String(name)
}

//Bool 读取配置
func Bool(name string) (bool, error) {
	return Conf.Bool(name)
}

//Int 读取配置
func Int(name string) (int, error) {
	return Conf.Int(name)
}

//JSON 处理json配置
func JSON(name string, v interface{}) error {
	return json.Unmarshal([]byte(Conf.String(name)), v)
}
