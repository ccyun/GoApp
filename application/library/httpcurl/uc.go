package httpcurl

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
)

var (
	//UcOpenAPIURL 登录服务器
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
	ToUsers        []uint64                `json:"toUsers,omitempty"`
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
	SiteID      uint64   `json:"siteId"`
	AppID       uint64   `json:"appId"`
	WebPushData string   `json:"webPushData,omitempty"`
	ToUsers     []string `json:"toUsers,omitempty"`
	ToPartyIds  []string `json:"toPartyIds,omitempty"`
}

//ResponseData response 结构体
type ResponseData struct {
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
	statusCode, res, err = Request(method, url, strings.NewReader(string(body)))
	logs.Debug("GetToken url:", url, "body:", string(body), "code:", statusCode)
	if statusCode != 200 {
		err = fmt.Errorf("uc httpcurl status code: %d", statusCode)
	}
	if err = json.Unmarshal(res, resData); err != nil {
		return err
	}
	rv := reflect.ValueOf(resData).Elem()
	requestID := rv.FieldByName("RequestID").String()
	errorCode := rv.FieldByName("ErrorCode").Uint()
	errorMessage := rv.FieldByName("ErrorMessage").String()
	logs.Debug("uc httpcurl errorCode:%d,requestID:%s,errorMessage:%s", errorCode, requestID, errorMessage)
	if errorCode != 0 {
		err = fmt.Errorf("uc httpcurl errorCode:%d,requestID:%s,errorMessage:%s", errorCode, requestID, errorMessage)
	}
	if err != nil {
		logs.Error("uc httpcurl error:", err, "response:", string(res))
		return err
	}
	return nil
}

//GetToken 获取token
func (U *UC) GetToken() string {
	var tokenData struct {
		ResponseData
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
	var data struct {
		Token string   `json:"token"`
		Data  OASender `json:"data"`
	}
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
	var resData ResponseData
	err := U.httpCurl("POST", url, data, &resData)
	if err != nil {
		logs.Error("OASend error:", err)
	}
	return err
}

//CustomizedSend 定制消息
func (U *UC) CustomizedSend() {

}
