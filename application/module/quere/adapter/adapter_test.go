package adapter

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type form interface{}

var Data form

func TestFormData(t *testing.T) {
	s := "bbs_id=2888&sub_reply_id=68&status=2&audit_opinion=11111111&data=eKKQopgUkGFaVqHxwzwWVmAlSPPIqRZt&data%5B0%5D%5Bstatus%5D=2&data%5B0%5D%5Bcomments%5D%5B0%5D%5Btop%5D=12&data%5B0%5D%5Bcomments%5D%5B0%5D%5Bleft%5D=45&data%5B0%5D%5Bcomments%5D%5B0%5D%5Btext%5D=fdfdsfsdfdfdsfsdsdds&data%5B1%5D%5Bid%5D=GcMIKLXinNHhbQqBepHVtuBVLjEWJknz&data%5B1%5D%5Bstatus%5D=2&data%5B2%5D%5Bid%5D=ijHupuvShCfxtlrNWFLgoKSxDXhYmoYk&data%5B2%5D%5Bstatus%5D=2"
	req := &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": {`application/x-www-form-urlencoded`}},
		Body:   ioutil.NopCloser(strings.NewReader(s)),
	}
	req.ParseForm()

	reflect
	// for k, _ := range req.Form {
	// 	log.Println(k)
	// }

}
