package controller

import (
	"log"
	"regexp"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"

	"github.com/astaxie/beego"
)

//Base 首页
type Base struct {
	beego.Controller
	UserID     uint64
	SessionID  string
	V          uint
	ClientType string
	Lang       string
	RequestID  string
}

//Init 构造函数
func (B *Base) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	B.Controller.Init(ctx, controllerName, actionName, app)
	B.Ctx.Input.Bind(&B.UserID, "user_id")
	B.Ctx.Input.Bind(&B.SessionID, "session_id")
	B.Ctx.Input.Bind(&B.V, "v")
	B.Ctx.Input.Bind(&B.ClientType, "client_type")
	B.Ctx.Input.Bind(&B.Lang, "lang")
	B.RequestID = string(utils.RandomCreateBytes(32))
	if regexp.MustCompile(`(application/json)(?:,|$)`).MatchString(B.Ctx.Input.Header("Content-Type")) {
		log.Println(string(B.Ctx.Input.RequestBody))
	} else {
		log.Println(string(B.Ctx.Input.RequestBody))
	}

	for k, v := range B.Input() {

		log.Println(k)
		log.Println(v)
	}

}
