package httpcurl

import (
	"fmt"
	"log"
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
func (U *UMS) GetAllUserByOrgIDs(orgIDs []uint64) error {

	//$res = $this->ums_api('/rs/organizations/query/orgs/users?pageNum=%s&pageSize=%s&productID=%s', 'post', array($page, $page_size, 20), $org_ids)

	return nil
}
func (U *UMS) _getAllUserByOrgIDs(orgIDs []uint64, page uint64) ([]UMSUser, uint64, error) {
	var totalCount uint64
	url := fmt.Sprintf("%s/rs/organizations/query/orgs/users?pageNum=%d&pageSize=%d&productID=%d", UMSBusinessURL, page, 500, 20)
	log.Println(url)
	return nil, totalCount, nil
}
