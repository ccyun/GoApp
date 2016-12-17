package httpcurl

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"sync"

	"math"

	"github.com/astaxie/beego/logs"
)

var (
	//UMSLoginURL 登录服务器
	UMSLoginURL string
	//UMSBusinessURL 非登录服务器
	UMSBusinessURL string
)

//UMS UMS
type UMS struct {
}

//UMSUser 用户
type UMSUser struct {
	UserID           uint64 `json:"id"`
	LoginName        string `json:"loginName"`
	UserStatus       uint   `json:"userstatus"`
	ProductStatus    uint   `json:"productStatus"`
	DisplayName      string `json:"displayName"`
	IconURL          string `json:"iconUrl"`
	OrganizationID   uint64 `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
}

//UMSOrg 组织
type UMSOrg struct {
	OrgID        uint64 `json:"id"`
	OrgName      string `json:"name"`
	ParentID     uint64 `json:"parentId"`
	Type         string `json:"type"`
	NodeCode     string `json:"nodeCode"`
	NodeCodeArr  []uint64
	CustomerCode string `json:"customercode"`
}

//GetAllUserByOrgIDs 批量获取组织下所有用户
func (U *UMS) GetAllUserByOrgIDs(orgIDs []uint64) ([]UMSUser, error) {
	var pageSize uint64 = 500
	data, totalCount, err := U._getAllUserByOrgIDs(orgIDs, pageSize, 1)
	if err != nil {
		return nil, err
	}
	if uint64(len(data)) < totalCount {
		pageNum := int(math.Ceil(float64(totalCount) / float64(pageSize)))
		var w sync.WaitGroup
		runtime.GOMAXPROCS(runtime.NumCPU())
		for i := 2; i <= pageNum; i++ {
			w.Add(1)
			go func(i int) {
				d, _, _ := U._getAllUserByOrgIDs(orgIDs, pageSize, i)
				data = append(data, d[0:]...)
				w.Done()
			}(i)
		}
		w.Wait()
	}
	if uint64(len(data)) != totalCount {
		return nil, fmt.Errorf("GetAllUserByOrgIDs error")
	}
	return data, nil
}

//_getAllUserByOrgIDs 批量获取组织下所有用户
func (U *UMS) _getAllUserByOrgIDs(orgIDs []uint64, pageSize uint64, page int) ([]UMSUser, uint64, error) {
	url := fmt.Sprintf("%s/rs/organizations/query/orgs/users?pageNum=%d&pageSize=%d&productID=%d", UMSBusinessURL, page, pageSize, 20)
	body, _ := json.Marshal(orgIDs)
	statusCode, res, err := Request("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, 0, err
	}
	logs.Debug("_getAllUserByOrgIDs url:", url, "body:", string(body), "code:", statusCode)
	var data struct {
		RetCode uint64 `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		RetObj  struct {
			TotalCount uint64    `json:"totalCount"`
			UserList   []UMSUser `json:"dataSet"`
		} `json:"retObj"`
	}
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, 0, err
	}
	return data.RetObj.UserList, data.RetObj.TotalCount, nil
}
