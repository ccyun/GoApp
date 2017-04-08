package httpcurl

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"

	"bbs_server/application/library/conf"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/library/redis"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/config"
)

//Conf 配置
var Conf config.Configer

var isInit = false

//initHTTPCurl 初始化httpcurl
func initHTTPCurl() {
	if isInit == false {
		conf.InitConfig("../../../cmd/TaskScript/conf.ini")

		cache, err := cache.NewCache("redis", conf.String("cache"))
		if err != nil {
			log.Println(err)
			return
		}
		redis.Cache = cache

		httpcurl.UMSLoginURL = conf.String("ums_login_url")
		httpcurl.UMSBusinessURL = conf.String("ums_business_url")
		//初始化uc配置
		httpcurl.UcOpenAPIURL = conf.String("uc_open_api_url")
		httpcurl.UcAPPID = conf.String("uc_open_appid")
		httpcurl.UcPaddword = conf.String("uc_open_password")
		//初始化ucc配置
		httpcurl.UccServerURL = conf.String("uccserver_url")
		isInit = true
	}
}

func initHTTPCurl() {
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(func() error {
		conf, err := config.NewConfig("ini", "../../../cmd/WebService/conf.ini")
		if err != nil {
			return err
		}
		Conf = conf
		return nil
	}, func() error {
		cache, err := cache.NewCache("redis", Conf.String("cache"))
		if err != nil {
			return err
		}
		redis.Cache = cache
		return nil
	}, func() error {
		//初始化ums配置
		UMSLoginURL = Conf.String("ums_login_url")
		UMSBusinessURL = Conf.String("ums_business_url")
		//初始化uc配置
		UcOpenAPIURL = Conf.String("uc_open_api_url")
		UcAPPID = Conf.String("uc_open_appid")
		UcPaddword = Conf.String("uc_open_password")
		return nil

	})
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
	data, err := ums.GetAllUserIDsByOrgIDs("0000445", []uint64{54169})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(data))
}

func TestGetToken(t *testing.T) {
	initHTTPCurl()
	a := new(UC)
	a.GetToken()
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

func TestGetUsersLoginName(t *testing.T) {
	initHTTPCurl()
	ums := new(UMS)
	data, err := ums.GetUsersLoginName("000092", []uint64{63706854}, true)
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

	new(UCC).CheckSession(1111, "fdsfsdf")
}
