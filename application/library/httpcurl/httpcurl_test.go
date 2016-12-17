package httpcurl

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func Test_Request(t *testing.T) {
	UMSBusinessURL = "http://192.168.28.173:8081/umsapi"
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
func Test__getAllUserByOrgIDs(t *testing.T) {
	UMSBusinessURL = "http://192.168.28.173:8081/umsapi"
	a := new(UMS)
	_, totalCount, err := a._getAllUserByOrgIDs([]uint64{2752}, 100, 1)
	if err != nil {
		t.Error(err)
	}
	if totalCount <= 0 {
		t.Error(totalCount)
	}
}

func Test_GetAllUserByOrgIDs(t *testing.T) {
	UMSBusinessURL = "http://192.168.28.173:8081/umsapi"
	a := new(UMS)
	data, err := a.GetAllUserByOrgIDs([]uint64{2752})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(data))
}
