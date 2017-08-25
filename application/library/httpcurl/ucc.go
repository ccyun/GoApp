package httpcurl

import (
	"bytes"
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
	UserID         uint64   `json:"user_id"`
	UserAccount    string   `json:"user_account"`
	SiteID         uint64   `json:"site_id"`
	DisplayName    string   `json:"display_name"`
	CustomerCode   string   `json:"customer_code"`
	DepartmentID   uint64   `json:"department_id"`
	OrgNodeCode    string   `json:"org_node_code"`
	OrgNodeCodeArr []uint64 `json:"-"`
}

//UCCDiscussData 讨论组数据
type UCCDiscussData struct {
	DiscussType    int      `json:"group_type"`
	DiscussStatus  int      `json:"group_status"`
	DiscussID      uint64   `json:"group_id"`
	DiscussName    string   `json:"group_name"`
	AdminID        uint64   `json:"admin_id"`
	AdminList      []uint64 `json:"admin_list"`
	Conversation   uint64   `json:"conversation"`
	ValidMemberIDs []uint64 `json:"valid_member_ids"`
}

//UCCResponseData response 结构体
type UCCResponseData struct {
	ErrorCode    uint64 `json:"code"`
	ErrorMessage string `json:"msg"`
	RequestID    string `json:"request_id"`
}

//BroadcastRange 广播消息发送范围
type BroadcastRange struct {
	SiteID  uint64   `json:"siteid"`
	UserIDs []uint64 `json:"userid"`
	OrgIDs  []uint64 `json:"orgid"`
}

//httpCurl
func (U *UCC) httpCurl(method string, url string, body string, resData interface{}) error {
	var (
		statusCode int
		res        []byte
		err        error
	)
	reqID := string(utils.RandomCreateBytes(8))
	logs.Debug("%s->ucc httpCurl url:%s body:%s", reqID, url, body)
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
		return "", err
	}
	return strconv.FormatUint(data.Data.Seq, 10), nil
}

//CheckSession 检测session
func (U *UCC) CheckSession(userID uint64, sessionID string) (UCCSessionData, error) {
	var data struct {
		UCCResponseData
		Data UCCSessionData `json:"data"`
	}
	value := url.Values{}
	value.Set("session_id", sessionID)
	value.Set("user_id", strconv.FormatUint(userID, 10))
	if err := U.httpCurl("POST", fmt.Sprintf("%s/user/check", UccServerURL), value.Encode(), &data); err != nil {
		return data.Data, err
	}
	for _, v := range strings.Split("-", data.Data.OrgNodeCode) {
		n, err := strconv.ParseUint(v, 10, 0)
		if err != nil {
			return data.Data, err
		}
		data.Data.OrgNodeCodeArr = append(data.Data.OrgNodeCodeArr, n)
	}

	return data.Data, nil
}

//GetDiscussInfo 获取讨论组信息
func (U *UCC) GetDiscussInfo(userID uint64, discussID uint64) (UCCDiscussData, error) {
	var data struct {
		UCCResponseData
		Data []UCCDiscussData `json:"data"`
	}
	if err := U.httpCurl("GET", fmt.Sprintf("%s/group/info?user_id=%d&group_id=%d", UccServerURL, userID, discussID), "", &data); err != nil {
		return UCCDiscussData{}, err
	}
	if len(data.Data) > 0 && data.Data[0].DiscussType == 2 {
		return data.Data[0], nil
	}
	return UCCDiscussData{}, nil
}

//SendBroadcastMsg 发送广播消息
func (U *UCC) SendBroadcastMsg(data []byte, siteID uint64, userIDs []uint64) error {
	reqID := string(utils.RandomCreateBytes(8))
	broadcastRange := BroadcastRange{SiteID: siteID, OrgIDs: []uint64{}}
	i := 0
	for {
		broadcastRange.UserIDs = U.getPublishScope(userIDs, i)
		if len(broadcastRange.UserIDs) == 0 {
			break
		}
		i++
		rangebyte, _ := json.Marshal(broadcastRange)
		v := url.Values{
			"range": []string{string(rangebyte)},
			"type":  []string{"openapi"},
		}
		url := fmt.Sprintf("%s/message/broadcast?%s", UccServerURL, v.Encode())
		statusCode, res, err := Request("POST", url, bytes.NewReader(data), "octet")
		if statusCode != 200 || err != nil {
			err = fmt.Errorf("%s->ucc httpcurl status code: %d,err：%s", reqID, statusCode, err.Error())
			return err
		}
		logs.Debug("%s->ucc httpcurl url：%s", reqID, url)
		logs.Debug("%s->ucc httpcurl response:%s", reqID, string(res))
	}
	return nil
}

//getPublishScope 分批发送消息
func (U *UCC) getPublishScope(users []uint64, page int) []uint64 {
	var (
		u     []uint64
		start int
		end   int
	)
	um := len(users)
	start = page * 200
	end = start + 200
	if start < um {
		if end > um {
			end = um
		}
		u = users[start:end]
	}
	return u
}
