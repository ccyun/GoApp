package httpcurl

import (
	"encoding/json"
	"fmt"
	"strings"

	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
)

//BillServerURL Bill server服务器
var BillServerURL string

//BILL 计费
type BILL struct {
}

//ReqAccepter 计费接收者
type ReqAccepter struct {
	SiteID     uint64   `json:"siteId"`
	UserID     []uint64 `json:"userId"`
	AcceptTime uint64   `json:"acceptTime"`
}

func (B *BILL) httpCurl(method string, url string, body string, resData interface{}) error {
	var (
		statusCode int
		res        []byte
		err        error
	)
	reqID := string(utils.RandomCreateBytes(8))
	logs.Debug("%s->bill httpCurl url:%s body:%s", reqID, url, string(body))
	statusCode, res, err = Request(method, url, strings.NewReader(body), "json")
	if statusCode != 200 {
		err = fmt.Errorf("%s->bill httpcurl status code: %d", reqID, statusCode)
	}
	logs.Debug("%s->bill httpcurl response:%s", reqID, string(res))
	if err = json.Unmarshal(res, resData); err != nil {
		return err
	}
	return err
}

//Accepter 按照接收者计费
func (B *BILL) Accepter(siteID uint64, userIDs []uint64) error {
	url := fmt.Sprintf("%s/interface/broadcast/accepter", BillServerURL)
	startIndex := 0
	userCount := len(userIDs)
	nowTime := uint64(time.Now().UnixNano() / 1e6)
	for true {
		if startIndex >= userCount {
			break
		}
		endIndex := startIndex + 1000
		if endIndex > userCount {
			endIndex = userCount
		}
		body, _ := json.Marshal(ReqAccepter{
			SiteID:     siteID,
			UserID:     userIDs[startIndex:endIndex],
			AcceptTime: nowTime,
		})
		if err := B.httpCurl("POST", url, string(body), ""); err != nil {
			return err
		}
		startIndex += 1000
	}
	return nil
}
