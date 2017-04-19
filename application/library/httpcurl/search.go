package httpcurl

import (
	"encoding/json"
	"fmt"
	"strings"

	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
)

//SolrServerURL Solr server服务器
var SolrServerURL string

//SOLRUser 用户
type SOLRUser struct {
	UserID      uint64   `json:"user_id"`
	DisplayName string   `json:"display_name"`
	Avatar      string   `json:"icon_url"`
	NodeCode    string   `json:"-"`
	NodeCodeArr []uint64 `json:"organization_path"`
	OrgID       uint64   `json:"organization_id"`
	OrgName     string   `json:"name"`
}

//SOLROrg 用户
type SOLROrg struct {
	OrgID          uint64 `json:"organization_id"`
	OrgName        string `json:"organization_name"`
	ParentID       uint64 `json:"organization_parent_id"`
	Type           uint   `json:"organization_type"`
	ChildNodeCount uint64 `json:"organization_child_node_count"`
	NodeCode       string `json:"organization_node_code"`
	CustomerCode   string `json:"customercode"`
}

//SOLR 搜索服务
type SOLR struct {
}

//httpCurl
func (S *SOLR) httpCurl(method string, url string, postData interface{}, resData interface{}) error {
	var (
		statusCode int
		res        []byte
		err        error
	)
	reqID := string(utils.RandomCreateBytes(8))
	body, _ := json.Marshal(postData)
	logs.Debug("%s->solr httpCurl url:%s body:%s", reqID, url, string(body))
	statusCode, res, err = Request(method, url, strings.NewReader(string(body)), "json")
	if statusCode != 200 {
		err = fmt.Errorf("%s->solr httpcurl status code: %d", reqID, statusCode)
	}
	logs.Debug("%s->solr httpcurl response:%s", reqID, string(res))
	if err = json.Unmarshal(res, resData); err != nil {
		return err
	}
	return err
}

//orgIDs2string orgids转字符串
func (S *SOLR) orgIDs2string(orgIDs []uint64) (orgStr string) {
	for _, v := range orgIDs {
		if orgStr != "" {
			orgStr += strconv.FormatUint(v, 10)
		} else {
			orgStr = strconv.FormatUint(v, 10)
		}
	}
	return
}

//SearchUser 搜索用户
func (S *SOLR) SearchUser(customerCode string, siteID uint64, userID uint64, orgIDs []uint64, keyword string, count uint) ([]SOLRUser, error) {
	url := fmt.Sprintf("%s?scope=[7]&indent=true&wt=json&start=0&customer_code=%s&keyword=%s&site_id=%d&rows=%d&shard.keys=%s!&user_id=%d&org_ids=%s", SolrServerURL, customerCode, keyword, siteID, count, customerCode, userID, S.orgIDs2string(orgIDs))
	var resData struct {
		Docs []SOLRUser `json:"docs"`
	}
	if err := S.httpCurl("GET", url, "", &resData); err != nil {
		return nil, err
	}
	return resData.Docs, nil
}

//SearchOrg 搜索组织
func (S *SOLR) SearchOrg(customerCode string, siteID uint64, userID uint64, orgIDs []uint64, keyword string, count uint) ([]SOLROrg, error) {
	url := fmt.Sprintf("%s?scope=[72]&indent=true&wt=json&start=0&customer_code=%s&keyword=%s&site_id=%d&rows=%d&shard.keys=%s!&user_id=%d&org_ids=%s", SolrServerURL, customerCode, keyword, siteID, count, customerCode, userID, S.orgIDs2string(orgIDs))
	var resData struct {
		Docs []SOLROrg `json:"docs"`
	}
	if err := S.httpCurl("GET", url, "", &resData); err != nil {
		return nil, err
	}
	return resData.Docs, nil
}
