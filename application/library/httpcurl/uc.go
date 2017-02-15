package httpcurl

import (
	"encoding/json"
	"fmt"
	"strings"

	"reflect"

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
	AppID          uint64                  `json:"appId"`
	Title          string                  `json:"title"`
	Color          string                  `json:"color"`
	TitleStyle     string                  `json:"titleStyle"`
	TitleElements  []OASendTitleElementser `json:"titleElements"`
	Status         uint8                   `json:"status"`
	DetailURL      string                  `json:"detailURL"`
	DetailAuth     uint8                   `json:"detailAuth"`
	CustomizedType string                  `json:"customizedType"`
	CustomizedData string                  `json:"customizedData"`
	ToUsers        []string                `json:"toUsers"`
	ToPartyIds     []string                `json:"toPartyIds"`
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
	Type      uint8  `json:"type"`
	Status    string `json:"status"`
	ImageType string `json:"ImageType"`
	ImageID   string `json:"imageId"`
	Content   string `json:"content"`
}

//CustomizedSender 定制消息
type CustomizedSender struct {
	SiteID      uint64   `json:"siteId"`
	AppID       uint64   `json:"appId"`
	WebPushData string   `json:"webPushData,omitempty"`
	ToUsers     []string `json:"toUsers"`
	ToPartyIds  []string `json:"toPartyIds"`
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
	if err = json.Unmarshal(res, resData); err == nil {
		t := reflect.ValueOf(resData)
		requestID := t.FieldByName("RequestID").String()
		errorCode := t.FieldByName("ErrorCode").Uint()
		errorMessage := t.FieldByName("ErrorMessage").String()
		if errorCode != 0 {
			err = fmt.Errorf("uc httpcurl errorCode: %d,requestID:%s,errorMessage:%s", errorCode, requestID, errorMessage)
		}
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
	data := map[string]string{"role": "3", "appId": UcAPPID, "password": UcPaddword}
	if err := U.httpCurl("POST", url, data, tokenData); err != nil {
		logs.Error("GetToken error:", err)
	}
	return tokenData.Data.Token
}
