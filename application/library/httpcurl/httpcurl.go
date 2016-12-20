package httpcurl //Request UMSRequest
import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/cache"
)

var (
	// Cache cache对象
	Cache cache.Cache
)

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