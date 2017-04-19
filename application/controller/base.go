package controller

import (
	"bbs_server/application/library/httpcurl"
	"regexp"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
	"github.com/ccyun/form2json"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// 通用错误码定义
const (
	ErrorCodeSuccess               = 0      //处理成功
	ErrorCodeForm2jsonError        = 100001 //form2json错误
	ErrorCodeParamsValidationError = 100002 //参数错误
	ErrorSessionIsNotValid         = 100003 //session无效
)

//outData json输出
type outData struct {
	Code      int64       `json:"code"`
	RequestID string      `json:"request_id"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

//Base 公共控制器
type Base struct {
	beego.Controller
	UserID       uint64
	SessionID    string
	SiteID       uint64
	CustomerCode string
	OrgID        uint64
	NodeCode     []uint64
	V            uint
	ClientType   string
	Lang         string
	RequestID    string
	PostData     string
}

//Init 构造函数
func (B *Base) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	B.Controller.Init(ctx, controllerName, actionName, app)
	//检测参数，检测session，检测权限
	func(funcs ...func() bool) {
		for _, f := range funcs {
			if !f() {
				B.StopRun()
				return
			}
		}
	}(B.Params, B.CheckSession)
	B.RequestID = string(utils.RandomCreateBytes(32))
}

//Params 参数处理
func (B *Base) Params() bool {
	var err error
	B.Ctx.Input.Bind(&B.UserID, "user_id")
	B.Ctx.Input.Bind(&B.SessionID, "session_id")
	B.Ctx.Input.Bind(&B.V, "v")
	B.Ctx.Input.Bind(&B.ClientType, "client_type")
	B.Ctx.Input.Bind(&B.Lang, "lang")
	B.RequestID = string(utils.RandomCreateBytes(32))
	if regexp.MustCompile(`(application/json)(?:,|$)`).MatchString(B.Ctx.Input.Header("Content-Type")) {
		B.PostData = string(B.Ctx.Input.RequestBody)
	} else {
		if B.PostData, err = form2json.Unmarshal(string(B.Ctx.Input.RequestBody), nil); err != nil {
			B.Error(ErrorCodeForm2jsonError, "form2json Unmarshal error:", err)
			return false
		}
	}
	v := validation.Validation{}
	v.Numeric(B.UserID, "user_id")
	v.Required(B.SessionID, "session_id")
	v.Numeric(B.V, "v")
	v.Required(B.ClientType, "client_type")
	v.Required(B.Lang, "lang")
	if v.HasErrors() {
		for _, e := range v.Errors {
			logs.Error(B.L("Params validation error:"), e.Key, e.Message)
		}
		B.Error(ErrorCodeParamsValidationError, "Params validation error:", fmt.Errorf("Params [user_id,session_id,v,client_type,lang] is not valid"))
		return false
	}
	return true
}

//CheckSession 检测用户session是否有效
func (B *Base) CheckSession() bool {
	data, err := new(httpcurl.UCC).CheckSession(B.UserID, B.SessionID)
	if err != nil {
		logs.Error(B.L("CheckSession error:%v"), err)
		B.Error(ErrorSessionIsNotValid, "session not valid", nil)
		return false
	}
	B.SiteID = data.SiteID
	B.CustomerCode = data.CustomerCode
	B.OrgID = data.DepartmentID
	B.NodeCode = data.OrgNodeCodeArr

	return true
}

//L 语言log
func (B *Base) L(l string) string {
	return B.RequestID + "  " + l
}

//Err 错误输出
func (B *Base) Error(errCode int64, msg string, err error) {
	B.Data["json"] = outData{
		Code:      errCode,
		RequestID: B.RequestID,
		Message:   fmt.Sprint(msg, err),
	}
	logs.Error(B.L(msg+"errcode:%d,error:%v"), errCode, err)
	B.ServeJSON()
}

//Success 处理成功
func (B *Base) Success(data interface{}) {
	B.Data["json"] = outData{
		Code:      ErrorCodeSuccess,
		RequestID: B.RequestID,
		Message:   "Successful!",
		Data:      data,
	}
	B.ServeJSON()
}
