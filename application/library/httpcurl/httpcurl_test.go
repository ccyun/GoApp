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
	"bbs_server/application/library/thrift/uc"

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
	log.Println(new(UC).getPublishScope([]string{"1", "2", "3", "4", "5", "6", "7", "8", "1", "2", "3", "4", "5", "6", "7", "8", "1", "2", "3", "4", "5", "6", "7", "8", "1", "2", "3", "4", "5", "6", "7", "8", "1", "2", "3", "4", "5", "6", "7", "8", "1", "2", "3", "4", "5", "6", "7", "8", "1", "2", "3", "4", "5", "6", "7", "8"}, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 9}, 1))
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

func TestSliceSort(t *testing.T) {
	a := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	userCount := len(a)
	start := 0
	for {
		end := start + 5
		if start >= userCount {
			break
		}
		if end > userCount {
			end = userCount
		}
		bb := a[start:end]
		start += 5
		fmt.Println(bb)
	}
}
func TestMsgSend(t *testing.T) {

	head := uc.NewUcMessageHead()
	head.Appid = 256
	head.Version = 1
	head.ControlPri = 15
	head.Conversation = 0
	from := uc.NewJID()
	from.UserID = 10488557
	from.SiteID = 72112
	head.From = from
	to := uc.NewJID()
	head.To = to
	head.Id = -1768291471
	head.Pri = 1
	head.Protocolid = 4
	head.Protocoltype = 1

	body := uc.NewUcMessageBody()

	body.ApiOA = uc.NewAPIOAContent()

	body.ApiOA.Title = "111"
	Color := "yellow"
	body.ApiOA.Color = &Color
	titleElements := `[{\"status\":1,\"title\":\"121212121212121212\",\"Color\":\"white\"}]`
	body.ApiOA.TitleElements = &titleElements
	customizedData := `{\"board_id\":50000201,\"board_name\":\"new公告\",\"avatar\":\"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gMTJfMjQwcHgucG5nXl5edGFuZ2hkZnNeXl42NzM5NTExNWFjMmFmM2NmMzk2MjYwNDNlMGI1Y2ZlOF5eXnRhbmdoZGZzXl5eMTAwMDE$\\u0026u=62051318\",\"discuss_id\":0,\"bbs_id\":10931,\"feed_id\":11401,\"title\":\"111\",\"description\":\"\",\"thumb\":\"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgOC5wbmdeXl50YW5naGRmc15eXjY4YjkyOTJkZmE3ZDFkYWQ0ODhlMzJmNTkyNjMxMzg1Xl5edGFuZ2hkZnNeXl41OTM4OTk$\\u0026u=62051318\",\"user_id\":64055154,\"type\":\"default\",\"category\":\"bbs\",\"link\":\"\",\"is_browser\":0,\"is_auth\":1,\"comment_enabled\":1,\"created_at\":1501470261095}`
	body.ApiOA.CustomizedData = &customizedData
	customizedType := "application/json"
	body.ApiOA.CustomizedType = &customizedType
	detailAuth := int8(1)
	body.ApiOA.DetailAuth = &detailAuth
	detailURL := "https://testcloudb.quanshi.com/bbsapp/bbs/show/bbs.html?id=10931&category=bbs&v=2"
	body.ApiOA.DetailURL = &detailURL
	body.ApiOA.Elements = `[{\"imageId\":\"http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgOC5wbmdeXl50YW5naGRmc15eXjY4YjkyOTJkZmE3ZDFkYWQ0ODhlMzJmNTkyNjMxMzg1Xl5edGFuZ2hkZnNeXl41OTM4OTk$&u=62051318\",\"status\":1,\"type\":\"image\"},{\"content\":\"新公告\",\"status\":1,\"type\":\"text\"}]`
	status := int16(1)
	body.ApiOA.Status = &status
	titleStyle := "simple"
	body.ApiOA.TitleStyle = &titleStyle

	//ucc := new(UCC)
	//	ucc.SendBroadcastMsg(head, body, 72112, []uint64{63706854, 63706854, 63706854, 63706854})

	// body_buf := thrift.NewTMemoryBuffer()
	// protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	// body.Write(protocolFactory.GetProtocol(body_buf))
	// body_bytes := body_buf.Bytes()
	// head.Length = int32(len(body_bytes))

	// head_buf := thrift.NewTMemoryBuffer()
	// message_protocol := thrift.NewTBinaryProtocolFactoryDefault()
	// head.Write(message_protocol.GetProtocol(head_buf))

	// head_buf.Write(body_bytes)
	// length := int32(len(head_buf.Bytes()))

	// rData := make([]byte, 4+length)
	// copy(rData[:4], function.Int32ToBytes(int32(length)))
	// copy(rData[4:], head_buf.Bytes())
	//fmt.Println(rData)
	// bodyBytes := bodyBuf.Bytes()
	// head.Length = int32(len(bodyBytes))
	//fmt.Println(string(function.CompressData(rData)))
	//	data := bytes.NewReader(function.CompressData(rData))
	//	rData = function.CompressData(rData)
	// initHTTPCurl()
	// v := url.Values{
	// 	//"user_id": []string{"64055154"},
	// 	//"data":  []string{string(function.CompressData(rData))},
	// 	"range": []string{`{"siteid":72112,"userid":[63706854]}`},
	// 	"type":  []string{"openapi"},
	// }
	//	url := "http://testcloud3.quanshi.com/uccserver/uccapi/message/broadcast?" + v.Encode()
	// fmt.Println(url)
	// var msgList *bytes.Buffer = new(bytes.Buffer)
	// lengte2 := uint32(len(rData))
	// err := binary.Write(msgList, binary.LittleEndian, lengte2)
	// err = binary.Write(msgList, binary.LittleEndian, rData)
	// fmt.Println(err)
	//fmt.Println(msgList)

	// _, o, _ := Request("POST", url, bytes.NewReader(function.CompressData(rData)), "octet")
	// fmt.Println(string(o))
	// fmt.Println(0x04)
}
