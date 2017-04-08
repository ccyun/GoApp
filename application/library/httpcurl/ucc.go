package httpcurl

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
)

//UccServerURL ucc server服务器
var UccServerURL string

//UCC UccServer
type UCC struct {
}

//UCCMsgSender 结构
type UCCMsgSender struct {
	UserID           uint64 `json:"user_id"`
	SiteID           uint64 `json:"site_id"`
	ConversationType int64  `json:"conversation_type"`
	Message          struct {
		ContentType int64  `json:"content_type"`
		Content     string `json:"content"`
	} `json:"message"`
	Control struct {
		NotSave    int64 `json:"notsave"`
		NoSendself int64 `json:"nosendself"`
	} `json:"control"`
	To struct {
		ToID         uint64   `json:"to_id"`
		ToPrivateIDs []uint64 `json:"toprivate_ids"`
	} `json:"to"`
}

//UCCSessionData session数据对象
type UCCSessionData struct {
	UserID       uint64 `json:"user_id"`
	UserAccount  string `json:"user_account"`
	SiteID       uint64 `json:"site_id"`
	DisplayName  string `json:"display_name"`
	CustomerCode string `json:"customer_code"`
	DepartmentID uint64 `json:"department_id"`
	OrgNodeCode  string `json:"org_node_code"`
}

//UCCResponseData response 结构体
type UCCResponseData struct {
	ErrorCode    uint64 `json:"code"`
	ErrorMessage string `json:"msg"`
	RequestID    string `json:"request_id"`
}

//httpCurl
func (U *UCC) httpCurl(method string, url string, body string, resData interface{}) error {
	var (
		statusCode int
		res        []byte
		err        error
	)
	reqID := string(utils.RandomCreateBytes(8))
	logs.Debug("%s->ucc httpCurl url:%s body:%s", reqID, url, string(body))
	statusCode, res, err = Request(method, url, strings.NewReader(body), "form")
	if statusCode != 200 {
		err = fmt.Errorf("%s->ucc httpcurl status code: %d", reqID, statusCode)
	}
	logs.Debug("%s->ucc httpcurl response:%s", reqID, string(res))
	if err = json.Unmarshal(res, resData); err != nil {
		return err
	}
	rv := reflect.ValueOf(resData).Elem()
	requestID := rv.FieldByName("RequestID").String()
	errorCode := rv.FieldByName("ErrorCode").Uint()
	errorMessage := rv.FieldByName("ErrorMessage").String()
	logs.Debug("%s->ucc httpcurl errorCode:%d,requestID:%s,errorMessage:%s", reqID, errorCode, requestID, errorMessage)
	if errorCode != 0 {
		err = fmt.Errorf("%s->ucc httpcurl errorCode:%d,requestID:%s,errorMessage:%s", reqID, errorCode, requestID, errorMessage)
	}
	return err
}

//MsgSend OA消息 return msg_id
func (U *UCC) MsgSend(postData UCCMsgSender) (string, error) {
	var data struct {
		UCCResponseData
		Data struct {
			Seq uint64 `json:"seq"`
		} `json:"data"`
	}
	postData.Control.NotSave = 0
	postData.Message.ContentType = 4
	value := url.Values{}
	value.Set("site_id", strconv.FormatUint(postData.SiteID, 10))
	value.Set("user_id", strconv.FormatUint(postData.UserID, 10))
	value.Set("conversation_type", strconv.FormatInt(postData.ConversationType, 10))
	message, err := json.Marshal(postData.Message)
	if err != nil {
		return "", err
	}
	value.Set("message", string(message))
	control, err := json.Marshal(postData.Control)
	if err != nil {
		return "", err
	}
	value.Set("control", string(control))
	to, err := json.Marshal(postData.To)
	if err != nil {
		return "", err
	}
	value.Set("to", string(to))
	if err := U.httpCurl("POST", fmt.Sprintf("%s/message/msgsend", UccServerURL), value.Encode(), &data); err != nil {
		logs.Error("MsgSend error:", err)
	}
	return strconv.FormatUint(data.Data.Seq, 10), nil
}

//CheckSession 检测session
func (U *UCC) CheckSession(userID uint64, sessionID string) UCCSessionData {
	var data struct {
		UCCResponseData
		Data UCCSessionData `json:"data"`
	}
	value := url.Values{}
	value.Set("session_id", sessionID)
	value.Set("user_id", strconv.FormatUint(userID, 10))
	if err := U.httpCurl("POST", fmt.Sprintf("%s/user/check", UccServerURL), value.Encode(), &data); err != nil {
		logs.Error("CheckSession error:", err)
	}
	return data.Data
}
