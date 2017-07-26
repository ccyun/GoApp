package httpcurl

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"bbs_server/application/library/conf"
	"bbs_server/application/library/redis"

	"github.com/astaxie/beego/cache"
)

var isInit = false

//initHTTPCurl 初始化httpcurl
func initHTTPCurl() {
	if isInit == false {
		conf.InitConfig("../../../cmd/_script/_script.ini")
		cache, err := cache.NewCache("redis", conf.String("cache"))
		if err != nil {
			log.Println(err)
			return
		}
		redis.Cache = cache
		UMSLoginURL = conf.String("ums_login_url")
		UMSBusinessURL = conf.String("ums_business_url")
		//初始化uc配置
		UcOpenAPIURL = conf.String("uc_open_api_url")
		UcAPPID = conf.String("uc_open_appid")
		UcPaddword = conf.String("uc_open_password")
		//初始化ucc配置
		UccServerURL = conf.String("uccserver_url")
		isInit = true
	}
}

func TestRequest(t *testing.T) {
	initHTTPCurl()
	url := fmt.Sprintf("%s/rs/organizations/query/orgs/users?pageNum=%d&pageSize=%d&productID=%d", UMSBusinessURL, 1, 500, 20)
	b, _ := json.Marshal([]uint{2752})
	statusCode, _, err := Request("POST", url, strings.NewReader(string(b)), "json")
	if err != nil {
		t.Error(err)
	}
	if statusCode != 200 {
		t.Error(statusCode)
	}
}

func TestGetAllUserByOrgIDs(t *testing.T) {
	initHTTPCurl()
	ums := new(UMS)
	data, err := ums.GetAllUserIDsByOrgIDs("0000445", []uint64{2752})
	if err != nil {
		t.Error(err)
	}
	log.Println(len(data))
	t.Log(len(data))
}

func TestGetToken(t *testing.T) {
	initHTTPCurl()
	a := new(UC)
	log.Println(a.GetToken())
}

func TestGetUsersDetail(t *testing.T) {
	initHTTPCurl()
	ums := new(UMS)
	data, err := ums.GetUsersDetail("000092", []uint64{63706854}, true)
	if err != nil {
		t.Error(err)
	}
	log.Println(data)
	t.Log(len(data))
}

func TestGetPublishScope(t *testing.T) {
	log.Println(new(UC).getPublishScope([]string{"1", "2", "3", "4", "5", "6", "7", "8"}, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0))
}

func TestCheckSession(t *testing.T) {
	initHTTPCurl()
	data, _ := new(UCC).CheckSession(63706854, "b4857f2dbfeeaf36f339510655e577e2e439e8c2")
	if data.UserID == 0 {
		t.Error(errors.New("check session error"))
	}
	log.Println(data)
}

func TestBill(t *testing.T) {
	userIDs := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	new(BILL).Accepter(111, userIDs)
}

func TestGetDiscussInfo(t *testing.T) {
	initHTTPCurl()
	//50033583
	info, err := new(UCC).GetDiscussInfo(63672505, 50033583)
	if err != nil {
		t.Error(err)
	}
	if info.DiscussID != 50033583 {
		t.Error("err")
	}
	log.Println(info)
}

func TestGetTags(t *testing.T) {
	initHTTPCurl()
	ums := new(UMS)
	data, err := ums.GetUserTags("0000000", 72112, []uint64{63706854, 63524288, 63661770})
	log.Println(err)
	log.Println(data)
	a, errr := json.Marshal(data)
	log.Println(errr)
	log.Println(string(a))
}
