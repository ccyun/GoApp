package httpcurl

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/config"
	"github.com/ccyun/GoApp/application/library/redis"
)

//Conf 配置
var Conf config.Configer

func initHTTPCurl() {
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(func() error {
		conf, err := config.NewConfig("ini", "../../../cmd/TaskScript/conf.ini")
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
		UMSLoginURL = Conf.String("ums_login_url")
		UMSBusinessURL = Conf.String("ums_business_url")
		return nil

	})
}

func TestRequest(t *testing.T) {
	initHTTPCurl()
	url := fmt.Sprintf("%s/rs/organizations/query/orgs/users?pageNum=%d&pageSize=%d&productID=%d", UMSBusinessURL, 1, 500, 20)
	b, _ := json.Marshal([]uint{2752})
	statusCode, _, err := Request("POST", url, strings.NewReader(string(b)))
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
