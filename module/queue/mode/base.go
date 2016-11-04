package mode

import (
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/model"
)

//base 任务处理适配器（基类）
type base struct {
	requestID    string
	o            orm.Ormer
	taskID       uint64
	siteID       uint64
	customerCode string
	action       string
	bbsID        uint64
	bbsInfo      model.BbsInfo
	PublishScope map[string][]uint64
	boardID      uint64
	feedID       uint64
}

//Begin 开启事务
func (B *base) Begin(requestID string) error {
	B.requestID = requestID
	B.o = orm.NewOrm()
	return B.o.Begin()
}

//Commit 提交事务
func (B *base) Rollback() error {
	return B.o.Rollback()
}

//Commit 提交事务
func (B *base) Commit() error {

	return B.o.Commit()
}

//NewTask 新任务对象
func (B *base) NewTask(taskInfo model.Queue) error {
	B.taskID = taskInfo.ID
	B.siteID = taskInfo.SiteID
	B.customerCode = taskInfo.CustomerCode
	B.action = taskInfo.Action
	return nil
}

//getBbsInfo 读取公告信息
func (B *base) getBbsInfo() {
	model := new(model.Bbs)
	bbsInfo, _ := model.GetOne(B.bbsID)
	B.bbsInfo = bbsInfo
}

//CreateFeed 创建Feed
func (B *base) CreateFeed() error {
	return nil
}

//CreateRelation 创建接收者关系
func (B *base) CreateRelation() error {
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *base) GetPublishScopeUsers() error {
	return nil
}

//CreateUnread 创建未读计数
func (B *base) CreateUnread() error {
	return nil
}

//SendMsg 发送消息
func (B *base) SendMsg() error {
	return nil
}
