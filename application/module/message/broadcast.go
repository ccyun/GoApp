package message

import (
	"bbs_server/application/function"
	"bbs_server/application/library/conf"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/library/thrift/uc"
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"git.apache.org/thrift.git/lib/go/thrift"
)

//BROADCAST 广播
type BROADCAST struct {
	SiteID  uint64
	MsgType int8
	Head    *uc.UcMessageHead
	Body    *uc.UcMessageBody
	Data    []byte
}

//OASendTitleElementser 标题元素
type OASendTitleElementser struct {
	Status uint8  `json:"status"`
	Title  string `json:"title"`
	Color  string `json:"Color"`
}

//OASendElementser 内容元素
type OASendElementser struct {
	Type      string `json:"type"`
	Status    uint8  `json:"status"`
	ImageType string `json:"ImageType,omitempty"`
	ImageID   string `json:"imageId,omitempty"`
	Content   string `json:"content,omitempty"`
}

//CustomizedSender 定制消息
type CustomizedSender struct {
	SiteID      string   `json:"siteId"`
	AppID       string   `json:"appId"`
	History     int      `json:"history"`
	WebPushData string   `json:"webPushData,omitempty"`
	ToUsers     []string `json:"toUsers,omitempty"`
	ToPartyIds  []uint64 `json:"toPartyIds,omitempty"`
	Data1       string   `json:"data1"`
	Data2       string   `json:"data2,omitempty"`
	Data3       string   `json:"data3,omitempty"`
	Data4       string   `json:"data4,omitempty"`
}

//SignalMsg 信令消息
type SignalMsg struct {
	Action  string `json:"action"`
	BoardID uint64 `json:"board_id"`
}

//消息类型
const (
	FEED_MSG   int8 = 1
	SIGNAL_MSG int8 = 2
)

//NewBroadcastMsg 广播消息
func NewBroadcastMsg(siteID uint64, msgType int8) *BROADCAST {
	b := new(BROADCAST)
	b.SiteID = siteID
	b.MsgType = msgType

	b.Head = uc.NewUcMessageHead()
	//b.Head = new(uc.UcMessageHead)
	b.Body = uc.NewUcMessageBody()
	//b.Body = new(uc.UcMessageBody)
	return b
}

//PackHead 组装消息头
func (B *BROADCAST) PackHead() {
	B.Head.Appid = int16(uc.AppId_AppAPI)
	B.Head.Version = 256
	//信令消息 仅通知不保存消息
	if B.MsgType == SIGNAL_MSG {
		B.Head.ControlPri = int8(uc.ControlPriType_webpush_count_type)
	}
	ChannelPri := int16(0)
	B.Head.ChannelPri = &ChannelPri
	B.Head.Conversation = 0
	from := uc.NewJID()
	userID, _ := conf.Int("uc_open_appid")
	from.UserID = int32(userID)
	from.SiteID = int32(B.SiteID)
	B.Head.From = from
	to := uc.NewJID()
	B.Head.To = to
	B.Head.Id = -1768291471
	B.Head.Pri = int8(uc.PriType_thrift_type)
	B.Head.Protocolid = 4
	B.Head.Protocoltype = 1
}

//OAMsg OA消息
func (B *BROADCAST) OAMsg(titleElements []OASendTitleElementser, customizedData string, detailURL string, elements []OASendElementser) {
	B.Head.Protocolid = int16(uc.APIMessageId_OA)
	elementbyte, _ := json.Marshal(elements)
	color := "yellow"
	customizedType := "application/json"
	status := int16(1)
	titleStyle := "simple"
	detailAuth := int8(1)
	titleData, _ := json.Marshal(titleElements)
	title := string(titleData)
	B.Body.ApiOA = uc.NewAPIOAContent()
	B.Body.ApiOA.Title = titleElements[0].Title
	B.Body.ApiOA.TitleElements = &title
	B.Body.ApiOA.Color = &color
	B.Body.ApiOA.Status = &status
	B.Body.ApiOA.DetailURL = &detailURL
	B.Body.ApiOA.CustomizedType = &customizedType
	B.Body.ApiOA.TitleStyle = &titleStyle
	B.Body.ApiOA.DetailAuth = &detailAuth
	B.Body.ApiOA.CustomizedData = &customizedData
	B.Body.ApiOA.Elements = string(elementbyte)
}

//CustomizedMsg 定制消息
func (B *BROADCAST) CustomizedMsg(data1, data2, data3, data4 string) {
	B.Head.Protocolid = int16(uc.APIMessageId_Customized)
	B.Body.ApiCustomized = uc.NewAPICustomizedContent()
	if B.MsgType == FEED_MSG {
		webPushData := "您有一个“i 广播”消息"
		B.Body.ApiCustomized.WebPushData = &webPushData
	}
	if data1 != "" {
		B.Body.ApiCustomized.Data1 = data1
	}
	if data2 != "" {
		B.Body.ApiCustomized.Data2 = &data2
	}
	if data3 != "" {
		B.Body.ApiCustomized.Data3 = &data3
	}
	if data4 != "" {
		B.Body.ApiCustomized.Data4 = &data4
	}
}

//SerilizeMsg 序列化消息
func (B *BROADCAST) SerilizeMsg() {
	bodyBuf := thrift.NewTMemoryBuffer()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	B.Body.Write(protocolFactory.GetProtocol(bodyBuf))
	bodyBytes := bodyBuf.Bytes()

	B.Head.Length = int32(len(bodyBytes))

	headBuf := thrift.NewTMemoryBuffer()
	messageProtocol := thrift.NewTBinaryProtocolFactoryDefault()
	B.Head.Write(messageProtocol.GetProtocol(headBuf))
	headBuf.Write(bodyBytes)
	B.Data = headBuf.Bytes()
}

//FormatSingleMsg 按照msgserver接口要求将消息格式化
func (B *BROADCAST) FormatSingleMsg() {
	length := len(B.Data)
	rData := make([]byte, 4+length)
	copy(rData[:4], function.Int32ToBytes(int32(length)))
	copy(rData[4:], B.Data)
	B.Data = function.CompressData(rData)
}

//Send 发送消息
func (B *BROADCAST) Send(userIDs []uint64) error {
	B.SerilizeMsg()
	B.FormatSingleMsg()
	head, _ := json.Marshal(B.Head)
	logs.Debug("broadcast msg head:%s", string(head))
	body, _ := json.Marshal(B.Body)
	logs.Debug("broadcast msg body:%s", string(body))
	return new(httpcurl.UCC).SendBroadcastMsg(B.Data, B.SiteID, userIDs)
}
