package application

import (
	"bbs_server/application/controller"

	"github.com/astaxie/beego"
)

type name struct {
}

//InitRouter 初始化路由
func InitRouter() error {

	beego.Router("/*", new(controller.Index), "*:Index")
	//beego.AutoRouter(new(controller.Index))
	return nil
}
