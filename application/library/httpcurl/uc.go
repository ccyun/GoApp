package httpcurl

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
)

var (
	//UcOpenAPIURL ucopenapi服务器
	UcOpenAPIURL string
	//UcAPPID appid
	UcAPPID string
	//UcPaddword password
	UcPaddword string
)

//UC UCopenAPI
type UC struct {
}

//OASender OA消息
type OASender struct {
	SiteID         uint64                  `json:"siteId"`
	AppID          string                  `json:"appId"`
	Title          string                  `json:"title"`
	Color          string                  `json:"color"`
	TitleStyle     string                  `json:"titleStyle"`
	TitleElements  []OASendTitleElementser `json:"titleElements"`
	Status         uint8                   `json:"status"`
	DetailURL      string                  `json:"detailURL"`
	DetailAuth     uint8                   `json:"detailAuth"`
	CustomizedType string                  `json:"customizedType"`
	CustomizedData string                  `json:"customizedData"`
	ToUsers        []string                `json:"toUsers,omitempty"`
	ToPartyIds     []uint64                `json:"toPartyIds,omitempty"`
	Elements       []OASendElementser      `json:"elements"`
}

//OASendTitleElementser 标题元素
type OASendTitleElementser struct {
	Status uint8  `json:"status"`
	Title  string `json:"title"`
	Color  string `json:"Color"`
}

//OASendElementser 内容元素
type OASendElementser struct {
	Type      string `json:"type"`
	Status    uint8  `json:"status"`
	ImageType string `json:"ImageType,omitempty"`
	ImageID   string `json:"imageId,omitempty"`
	Content   string `json:"content,omitempty"`
}

//CustomizedSender 定制消息
type CustomizedSender struct {
	SiteID      string   `json:"siteId"`
	AppID       string   `json:"appId"`
	WebPushData string   `json:"webPushData,omitempty"`
	ToUsers     []string `json:"toUsers,omitempty"`
	ToPartyIds  []uint64 `json:"toPartyIds,omitempty"`
	Data1       string   `json:"data1"`
	Data2       string   `json:"data2,omitempty"`
	Data3       string   `json:"data3,omitempty"`
	Data4       string   `json:"data4,omitempty"`
}

//UCResponseData response 结构体
type UCResponseData struct {
	ErrorCode    uint64 `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	RequestID    string `json:"requestId"`
}

//httpCurl
func (U *UC) httpCurl(method string, url string, postData interface{}, resData interface{}) error {
	var (
		statusCode int
		res        []byte
		err        error
	)
	body, _ := json.Marshal(postData)
	reqID := string(utils.RandomCreateBytes(8))
	logs.Debug("%s->uc httpCurl url:%s body:%s", reqID, url, string(body))
	statusCode, res, err = Request(method, url, strings.NewReader(string(body)), "json")
	logs.Debug("%s->uc httpCurl url:%s body:%s code:%d", reqID, url, string(body), statusCode)
	if statusCode != 200 {
		err = fmt.Errorf("%s->uc httpcurl status code: %d", reqID, statusCode)
	}
	logs.Debug("%s->uc httpcurl response:%s", reqID, string(res))
	if err = json.Unmarshal(res, resData); err != nil {
		return err
	}
	rv := reflect.ValueOf(resData).Elem()
	requestID := rv.FieldByName("RequestID").String()
	errorCode := rv.FieldByName("ErrorCode").Uint()
	errorMessage := rv.FieldByName("ErrorMessage").String()
	logs.Debug("%s->uc httpcurl errorCode:%d,requestID:%s,errorMessage:%s", reqID, errorCode, requestID, errorMessage)

	if errorCode != 0 {
		err = fmt.Errorf("%s->uc httpcurl errorCode:%d,requestID:%s,errorMessage:%s", reqID, errorCode, requestID, errorMessage)
	}
	return err
}

//GetToken 获取token
func (U *UC) GetToken() string {
	var tokenData struct {
		UCResponseData
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	url := fmt.Sprintf("%s/auth/token/create", UcOpenAPIURL)
	postData := map[string]string{"role": "3", "appId": UcAPPID, "password": UcPaddword}
	if err := U.httpCurl("POST", url, postData, &tokenData); err != nil {
		logs.Error("GetToken error:", err)
	}
	return tokenData.Data.Token
}

//OASend OA消息
func (U *UC) OASend(postData OASender) error {
	var (
		resData UCResponseData
		data    struct {
			Token string   `json:"token"`
			Data  OASender `json:"data"`
		}
		err error
	)
	postData.AppID = UcAPPID
	postData.Color = "yellow"
	postData.TitleStyle = "simple"
	postData.Status = 1
	postData.DetailAuth = 1
	postData.CustomizedType = "application/json"
	postData.TitleElements[0].Status = 1
	postData.TitleElements[0].Color = "white"
	postData.Elements[0].Type = "image"
	postData.Elements[0].ImageType = "url"
	postData.Elements[0].Status = 1
	postData.Elements[1].Type = "text"
	postData.Elements[1].Status = 1
	data.Data = postData
	data.Token = U.GetToken()
	url := fmt.Sprintf("%s/appmsg/oa/send", UcOpenAPIURL)
	i := 0
	for true {
		tempData := data
		tempData.Data.ToUsers, tempData.Data.ToPartyIds = U.getPublishScope(data.Data.ToUsers, data.Data.ToPartyIds, i)
		if len(tempData.Data.ToUsers) == 0 && len(tempData.Data.ToPartyIds) == 0 {
			break
		} else {
			err = U.httpCurl("POST", url, tempData, &resData)
			if err != nil {
				logs.Error("OASend error:", err)
				return err
			}
		}
		i++
	}
	return nil
}

//CustomizedSend 定制消息
func (U *UC) CustomizedSend(postData CustomizedSender) error {
	var (
		resData UCResponseData
		data    struct {
			Token string           `json:"token"`
			Data  CustomizedSender `json:"data"`
		}
		err error
	)
	postData.AppID = UcAPPID
	data.Data = postData
	data.Token = U.GetToken()
	url := fmt.Sprintf("%s/appmsg/customized/send", UcOpenAPIURL)
	i := 0
	for true {
		tempData := data
		tempData.Data.ToUsers, tempData.Data.ToPartyIds = U.getPublishScope(data.Data.ToUsers, data.Data.ToPartyIds, i)
		if len(tempData.Data.ToUsers) == 0 && len(tempData.Data.ToPartyIds) == 0 {
			break
		} else {
			err = U.httpCurl("POST", url, tempData, &resData)
			if err != nil {
				logs.Error("CustomizedSend error:", err)
				return err
			}
		}
		i++
	}
	return nil
}

//getPublishScope 分批发送消息
func (U *UC) getPublishScope(users []string, partyIds []uint64, page int) ([]string, []uint64) {
	var (
		u []string
		p []uint64
	)
	um := len(users)
	pm := len(partyIds)
	start := page * 20
	end := start + 20
	if start < um {
		if end > um {
			end = um
		}
		u = users[start:end]
	}
	end = start + 20
	if start < pm {
		if end > pm {
			end = pm
		}
		p = partyIds[start:end]
	}
	return u, p
}
