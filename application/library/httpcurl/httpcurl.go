package httpcurl //Request UMSRequest
import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/ccyun/GoApp/application/function"
	"github.com/garyburd/redigo/redis"
)

var (
	// Cache cache对象
	Cache cache.Cache
	//RequestID 请求ID
	RequestID string
)

//C 缓存结构
type C struct {
	customerCode string
	key          string
}

//Request curl请求
func Request(method string, url string, body io.Reader) (int, []byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return -1, nil, fmt.Errorf("construct http request failed, requrl = %s, err:%s", url, err.Error())
	}
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
		var respBody []byte
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			defer reader.Close()
			respBody, _ = ioutil.ReadAll(reader)
		default:
			respBody, _ = ioutil.ReadAll(response.Body)
		}
		if response.StatusCode > 300 {
			return response.StatusCode, respBody, fmt.Errorf("http request fail, url: %s", url)
		}
		return response.StatusCode, respBody, nil
	}
	if err != nil {
		return -1, nil, fmt.Errorf("http request fail, url: %s, error:%s", url, err.Error())
	}
	return -1, nil, fmt.Errorf("http request fail, url: %s, error:%s", url, err.Error())
}

//L 语言log
func L(log string) string {
	return RequestID + "  " + log
}

///////////////////////////////Cache//////////////////////////////////////////////////////////////////////////////////////////////////////

//newCache 初始化缓存对象
func newCache(customerCode string, args ...interface{}) *C {
	c := new(C)
	c.customerCode = customerCode
	c.key = c.makeKey(args)
	return c
}

//makeKey 参数产生Key
func (c *C) makeKey(args ...interface{}) string {
	k, err := json.Marshal(args)
	if err != nil {
		logs.Error(L("GetCache make key error"), err)
		return ""
	}
	return fmt.Sprintf("U%s:%s", c.customerCode, function.Md5(string(k), 32))
}

//setCache 设置缓存
func (c *C) setCache(data interface{}) bool {
	var (
		val []byte
		err error
	)
	if val, err = json.Marshal(data); err != nil {
		logs.Error(L("SetCache data Marshal error"), err)
		return false
	}
	if err := Cache.Put(c.key, val, 48*time.Hour); err != nil {
		logs.Error(L("SetCache Put error"), err)
		return false
	}
	return true
}

//getCache 读取缓存
func (c *C) getCache(data interface{}) bool {
	var (
		err error
		val string
	)
	if val, err = redis.String(Cache.Get(c.key), nil); err != nil {
		logs.Info(L("GetCache value Assertion error"), err)
		return false
	}
	if err = json.Unmarshal([]byte(val), data); err != nil {
		logs.Error(L("GetCache data Unmarshal error"), err)
		return false
	}
	return true
}
