package httpcurl

import (
	"encoding/json"
	"fmt"
	"log"
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

//_afterAPI
func (U *UC) _afterAPI() {

}

//GetToken 获取token
func (U *UC) GetToken() string {
	url := fmt.Sprintf("%s/auth/token/create", UcOpenAPIURL)
	data := make(map[string]string)
	data["role"] = "3"
	data["appId"] = UcAPPID
	data["password"] = UcPaddword
	body, _ := json.Marshal(data)
	statusCode, res, err := Request("POST", url, strings.NewReader(string(body)))
	if err != nil {
		logs.Error("GetToken url:", url, "body:", string(body), "error:", err)
	}
	logs.Debug("GetToken url:", url, "body:", string(body), "code:", statusCode)

	log.Println(string(res))
	return ""
}
