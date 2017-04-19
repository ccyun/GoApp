package controller

import (
	"bbs_server/application/library/httpcurl"
	"encoding/json"
)

//Org 组织通讯录
type Org struct {
	Base
}
type orgInfoer struct {
	OrgID        uint64       `json:"org_id"`
	OrgName      string       `json:"org_name"`
	ParentID     uint64       `json:"parent_id"`
	Type         uint         `json:"type"`
	NodeCode     string       `json:"node_code"`
	NodeCodeArr  []uint64     `json:"node_code_arr"`
	CustomerCode string       `json:"customer_code"`
	DeptList     []orgInfoer  `json:"dept_list"`
	MemberList   []userInfoer `json:"member_list"`
}
type userInfoer struct {
	UserID           uint64 `json:"user_id"`
	LoginName        string `json:"user_account"`
	UserStatus       uint   `json:"user_status"`
	ProductStatus    uint   `json:"product_status"`
	DisplayName      string `json:"display_name"`
	Avatar           string `json:"avatar"`
	OrganizationID   uint64 `json:"org_id"`
	OrganizationName string `json:"org_name"`
}

func (O *Org) umsOrgHandle(data []httpcurl.UMSOrg) (data2 []orgInfoer) {
	for _, orgInfo := range data {
		data2 = append(data2, orgInfoer{
			OrgID:        orgInfo.OrgID,
			OrgName:      orgInfo.OrgName,
			ParentID:     orgInfo.ParentID,
			Type:         orgInfo.Type,
			NodeCode:     orgInfo.NodeCode,
			NodeCodeArr:  orgInfo.NodeCodeArr,
			CustomerCode: orgInfo.CustomerCode,
		})
	}
	return data2
}

func (O *Org) umsUserHandle(data []httpcurl.UMSUser) (data2 []userInfoer) {
	for _, userInfo := range data {
		data2 = append(data2, userInfoer{
			UserID:           userInfo.UserID,
			LoginName:        userInfo.LoginName,
			UserStatus:       userInfo.UserStatus,
			ProductStatus:    userInfo.ProductStatus,
			DisplayName:      userInfo.DisplayName,
			Avatar:           userInfo.Avatar,
			OrganizationID:   userInfo.OrganizationID,
			OrganizationName: userInfo.OrganizationName,
		})
	}
	return data2
}

//GetOrgUser 查询子组织及成员
func (O *Org) GetOrgUser() {
	ums := new(httpcurl.UMS)
	orgID, err := O.GetUint64("org_id")
	if err != nil {
		O.Error(ErrorCodeParamsValidationError, "org_id is not valid", err)
		return
	}
	orgInfo, _ := ums.BatchQueryOrg(O.CustomerCode, []uint64{orgID})
	if len(orgInfo) == 0 {
		O.Success([]interface{}{})
	}

	var data orgInfoer
	data = O.umsOrgHandle(orgInfo)[0]
	if orgList, _ := ums.GetOrgChilds(O.CustomerCode, orgID); len(orgList) > 0 {
		data.DeptList = O.umsOrgHandle(orgList)
	}
	if userList, _ := ums.GetOrgMembers(O.CustomerCode, orgID); len(userList) > 0 {
		data.MemberList = O.umsUserHandle(userList)
	}
	O.Success(data)
}

//UserList 批量查询用户信息
func (O *Org) UserList() {
	var inPutData struct {
		UserIDs []uint64 `json:"user_ids"`
	}
	if err := json.Unmarshal([]byte(O.PostData), &inPutData); err != nil {
		O.Error(ErrorCodeParamsValidationError, "userIDs is not valid", err)
	}
	if userList, _ := new(httpcurl.UMS).GetUsersDetail(O.CustomerCode, inPutData.UserIDs, false); len(userList) > 0 {
		O.Success(O.umsUserHandle(userList))
	}
	O.Success([]interface{}{})
}

//OrgByCode 根据customer_code查询组织详情
func (O *Org) OrgByCode() {
	customerCode := O.GetString("customer_code")
	if customerCode == "" {
		O.Error(ErrorCodeParamsValidationError, "customer_code is not valid", nil)
	}
	data, err := new(httpcurl.UMS).GetOrgByCustomerCode(customerCode)
	if err != nil {
		O.Success([]interface{}{})
	}
	O.Success(O.umsOrgHandle([]httpcurl.UMSOrg{data})[0])
}

//SearchOrg 搜索组织
func (O *Org) SearchOrg() {

}

//SearchUser 搜索用户
func (O *Org) SearchUser() {

}
