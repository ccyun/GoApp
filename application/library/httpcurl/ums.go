package httpcurl

import (
	"encoding/json"
	"fmt"
	"strings"

	"math"

	"bbs_server/application/library/redis"

	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils"
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
	OrganizationID   uint64 `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	Avatar           string `json:"iconUrl"`
	UserProductList  []struct {
		ProductID  int64 `json:"productId"`
		UserStatus int64 `json:"userStatus"`
	} `json:"userProductList"`
}

//UMSOrg 组织
type UMSOrg struct {
	OrgID        uint64 `json:"id"`
	OrgName      string `json:"name"`
	ParentID     uint64 `json:"parentId"`
	Type         uint   `json:"type"`
	NodeCode     string `json:"nodeCode"`
	NodeCodeArr  []uint64
	CustomerCode string `json:"customercode"`
}

//UMSTag tag结构
type UMSTag struct {
	TagID     uint64 `json:"tagId"`
	TagName   string `json:"tagName"`
	TagCode   string `json:"tagCode"`
	TagEnumID uint64 `json:"tagEnumId"`
	TagValue  string `json:"tagValue"`
}

//TagValueReq 查询tag用户参数
//{"conditionList":[{"tagId":187382,"tagValueList":[1588,1599]}]}
// type TagValueReq struct {
// 	TagID    uint64   `json:"tagId"`
// 	TagValue []uint64 `json:"tagValueList"`
// }

func (U *UMS) httpCurl(method string, url string, postData interface{}, resData interface{}) error {
	var (
		statusCode int
		res        []byte
		err        error
	)
	reqID := string(utils.RandomCreateBytes(8))
	body, _ := json.Marshal(postData)
	logs.Debug("%s->ums httpCurl url:%s body:%s", reqID, url, string(body))
	statusCode, res, err = Request(method, url, strings.NewReader(string(body)), "json")
	if statusCode != 200 {
		err = fmt.Errorf("%s->ums httpcurl status code: %d", reqID, statusCode)
	}
	logs.Debug("%s->ums httpcurl response:%s", reqID, string(res))
	if err = json.Unmarshal(res, resData); err != nil {
		return err
	}
	return err
}

//GetAllUserIDsByOrgIDs 批量获取组织下所有用户ID
func (U *UMS) GetAllUserIDsByOrgIDs(customerCode string, orgIDs []uint64) ([]uint64, error) {
	var data []uint64
	// cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetAllUserIDsByOrgIDs", orgIDs)
	// if cache.Get(&data) == true {
	// 	return data, nil
	// }
	UserList, err := U.GetAllUserByOrgIDs(customerCode, orgIDs)
	if err != nil {
		return nil, err
	}
	for _, userInfo := range UserList {
		data = append(data, userInfo.UserID)
	}
	//	cache.Set(data)
	return data, nil
}

//GetAllUserByOrgIDs 批量获取组织下所有用户
func (U *UMS) GetAllUserByOrgIDs(customerCode string, orgIDs []uint64) ([]UMSUser, error) {
	var (
		pageSize   uint64 = 500
		totalCount uint64
		err        error
		data       []UMSUser
	//	mutex      sync.Mutex
	)
	// cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetAllUserByOrgIDs", orgIDs)
	// if cache.Get(&data) == true {
	// 	return data, nil
	// }
	data, totalCount, err = U._getAllUserByOrgIDs(orgIDs, pageSize, 1)
	if err != nil {
		return nil, err
	}
	if uint64(len(data)) < totalCount {
		pageNum := int(math.Ceil(float64(totalCount) / float64(pageSize)))
		//	var w sync.WaitGroup
		//	runtime.GOMAXPROCS(runtime.NumCPU())
		for i := 2; i <= pageNum; i++ {
			//		w.Add(1)
			//		go func(i int) {
			if d, _, _ := U._getAllUserByOrgIDs(orgIDs, pageSize, i); len(d) > 0 {
				//			mutex.Lock()
				data = append(data, d...)
				//		mutex.Unlock()
			}
			//		w.Done()
			//	}(i)
		}
		//	w.Wait()
	}
	if uint64(len(data)) != totalCount {
		logs.Error("GetAllUserByOrgIDs error totalCount!=len(data)")
	}
	//cache.Set(data)
	return data, nil
}

//_getAllUserByOrgIDs 批量获取组织下所有用户
func (U *UMS) _getAllUserByOrgIDs(orgIDs []uint64, pageSize uint64, page int) ([]UMSUser, uint64, error) {
	url := fmt.Sprintf("%s/rs/organizations/query/orgs/users?pageNum=%d&pageSize=%d&productID=%d", UMSBusinessURL, page, pageSize, 20)
	var resData struct {
		RetCode uint64 `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		RetObj  struct {
			TotalCount uint64    `json:"totalCount"`
			UserList   []UMSUser `json:"dataSet"`
		} `json:"retObj"`
	}
	if err := U.httpCurl("POST", url, orgIDs, &resData); err != nil {
		return nil, 0, err
	}
	return resData.RetObj.UserList, resData.RetObj.TotalCount, nil
}

//GetUsersDetail 批量查询用户详情
func (U *UMS) GetUsersDetail(customerCode string, userIDs []uint64, isValid bool) ([]UMSUser, error) {
	var data []UMSUser
	// cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetUsersDetail", userIDs, isValid)
	// if cache.Get(&data) == true {
	// 	return data, nil
	// }
	url := fmt.Sprintf("%s/rs/users/id/in?requestType=0", UMSBusinessURL)
	userCount := len(userIDs)
	start := 0
	for {
		end := start + 500
		if start >= userCount {
			break
		}
		if end > userCount {
			end = userCount
		}
		resData := []UMSUser{}
		if err := U.httpCurl("POST", url, userIDs[start:end], &resData); err != nil {
			return nil, err
		}
		if isValid == false {
			data = append(data, resData...)
		} else {
			for _, user := range resData {
				for _, v := range user.UserProductList {
					if v.ProductID == 20 && (v.UserStatus == 82 || v.UserStatus == 9) {
						data = append(data, user)
					}
				}
			}
		}
		start += 500
	}
	//cache.Set(data)
	return data, nil
}

//BatchQueryOrg 批量查询组织信息
func (U *UMS) BatchQueryOrg(customerCode string, orgIDs []uint64) ([]UMSOrg, error) {
	var data []UMSOrg
	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "BatchQueryOrg", orgIDs)
	if cache.Get(&data) == true {
		return data, nil
	}
	url := fmt.Sprintf("%s/rs/organizations/batchquery?productID=20&child=0", UMSBusinessURL)
	if err := U.httpCurl("POST", url, orgIDs, &data); err != nil {
		return nil, err
	}
	cache.Set(data)
	return data, nil
}

//GetOrgChilds 查询子组织
func (U *UMS) GetOrgChilds(customerCode string, orgID uint64) ([]UMSOrg, error) {
	var data []UMSOrg
	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetOrgChilds", orgID)
	if cache.Get(&data) == true {
		return data, nil
	}
	url := fmt.Sprintf("%s/rs/organizations/%d?scope=nextlevel&types=1,2,3,4,5", UMSBusinessURL, orgID)
	if err := U.httpCurl("GET", url, "", &data); err != nil {
		return nil, err
	}
	cache.Set(data)
	return data, nil
}

//GetOrgMembers 查询组织成员
func (U *UMS) GetOrgMembers(customerCode string, orgID uint64) ([]UMSUser, error) {
	var data []UMSUser
	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetOrgMembers", orgID)
	if cache.Get(&data) == true {
		return data, nil
	}
	url := fmt.Sprintf("%s/rs/organizations/%d/users?productID=20", UMSBusinessURL, orgID)
	if err := U.httpCurl("GET", url, "", &data); err != nil {
		return nil, err
	}
	cache.Set(data)
	return data, nil
}

//GetOrgByCustomerCode 根据customerCode查询组织信息
func (U *UMS) GetOrgByCustomerCode(customerCode string) (data UMSOrg, err error) {
	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetOrgByCustomerCode")
	if cache.Get(&data) == true {
		return data, nil
	}
	url := fmt.Sprintf("%s/rs/organizations?customer_code=%s", UMSBusinessURL, customerCode)
	if err = U.httpCurl("GET", url, "", &data); err == nil {
		cache.Set(data)
	}
	return data, err
}

//GetUserTags 批量查询用户标签
func (U *UMS) GetUserTags(customerCode string, siteID uint64, userIDs []uint64) (data map[uint64][]UMSTag, err error) {
	data = make(map[uint64][]UMSTag)
	data2 := make(map[string][]UMSTag)
	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetUserTags", siteID, userIDs)
	if cache.Get(&data2) == false {
		postData := struct {
			SiteID    uint64   `json:"siteId"`
			ProductID uint64   `json:"productId"`
			UserIDs   []uint64 `json:"userIds"`
		}{
			SiteID:    siteID,
			ProductID: 20,
			UserIDs:   userIDs,
		}
		var resData struct {
			RetCode uint64 `json:"retCode"`
			RetMsg  string `json:"retMsg"`
			RetObj  []struct {
				UserID   uint64   `json:"userId"`
				TagGroup []UMSTag `json:"tagGroup"`
			} `json:"retObj"`
		}
		url := fmt.Sprintf("%s/rs/users/tagGroup", UMSBusinessURL)
		if err = U.httpCurl("POST", url, postData, &resData); err == nil {
			for _, item := range resData.RetObj {
				k := strconv.FormatUint(item.UserID, 10)
				data2[k] = item.TagGroup
			}
			cache.Set(data2)
		}
	}
	for k, item := range data2 {
		key, _ := strconv.ParseUint(k, 10, 0)
		data[key] = item
	}
	return data, err
}

// //GetTagUserIDs 查询标签组的用户IDs
// func (U *UMS) GetTagUserIDs(customerCode string, siteID uint64, tagValue []TagValueReq) ([]uint64, error) {
// 	var data []uint64
// 	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetTagUserIDs", siteID, tagValue)
// 	if cache.Get(&data) == true {
// 		return data, nil
// 	}
// 	UserList, err := U.GetTagUserList(customerCode, siteID, tagValue)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, userInfo := range UserList {
// 		data = append(data, userInfo.UserID)
// 	}
// 	cache.Set(data)
// 	return data, nil
// }

// //GetTagUserList 批量获取组织下所有用户列表
// func (U *UMS) GetTagUserList(customerCode string, siteID uint64, tagValue []TagValueReq) ([]UMSUser, error) {
// 	var (
// 		pageSize   uint64 = 50
// 		totalCount uint64
// 		err        error
// 		tempData   []UMSUser
// 		data       []UMSUser
// 		postData   struct {
// 			ConditionList []TagValueReq `json:"conditionList"`
// 		}
// 	)

// 	cache := redis.NewCache(fmt.Sprintf("U%s", customerCode), "GetTagUserList", siteID, tagValue)
// 	if cache.Get(&data) == true {
// 		return data, nil
// 	}
// 	postData.ConditionList = tagValue
// 	tempData, totalCount, err = U._getTagUserList(postData, pageSize, 1, siteID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if uint64(len(tempData)) < totalCount {
// 		pageNum := int(math.Ceil(float64(totalCount) / float64(pageSize)))
// 		var w sync.WaitGroup
// 		runtime.GOMAXPROCS(runtime.NumCPU())
// 		for i := 2; i <= pageNum; i++ {
// 			w.Add(1)
// 			go func(i int) {
// 				d, _, _ := U._getTagUserList(postData, pageSize, i, siteID)
// 				tempData = append(tempData, d[0:]...)
// 				w.Done()
// 			}(i)
// 		}
// 		w.Wait()
// 	}
// 	if uint64(len(tempData)) != totalCount {
// 		return nil, fmt.Errorf("GetAllUserByOrgIDs error")
// 	}
// 	for _, user := range tempData {
// 		for _, v := range user.UserProductList {
// 			if v.ProductID == 20 && (v.UserStatus == 82 || v.UserStatus == 9) {
// 				data = append(data, user)
// 			}
// 		}
// 	}
// 	cache.Set(data)
// 	return data, nil
// }

// func (U *UMS) _getTagUserList(postData interface{}, pageSize uint64, page int, siteID uint64) ([]UMSUser, uint64, error) {
// 	url := fmt.Sprintf("%s/rs/tag/search?pageNum=%d&pageSize=%d&productID=%d&siteId=%d", UMSBusinessURL, page, pageSize, 20, siteID)
// 	var (
// 		resData struct {
// 			RetCode uint64 `json:"retCode"`
// 			RetMsg  string `json:"retMsg"`
// 			RetObj  struct {
// 				TotalCount uint64 `json:"totalCount"`
// 				UserList   []struct {
// 					Info UMSUser `json:"user"`
// 				} `json:"dataSet"`
// 			} `json:"retObj"`
// 		}
// 		data []UMSUser
// 	)
// 	if err := U.httpCurl("POST", url, postData, &resData); err != nil {
// 		return nil, 0, err
// 	}
// 	for _, UserList := range resData.RetObj.UserList {
// 		data = append(data, UserList.Info)
// 	}
// 	return data, resData.RetObj.TotalCount, nil
// }
