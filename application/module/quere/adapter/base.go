package adapter

import (
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/model"
)

//base 任务处理适配器（基类）
type base struct {
	o            orm.Ormer
	taskID       uint64
	siteID       uint64
	customerCode string
	action       string
	bbsID        uint64
	bbsInfo      model.Bbs
	PublishScope map[string][]uint64
	boardID      uint64
	boardInfo    model.Board
	feedID       uint64
}

//Begin 开启事务
func (B *base) Begin() error {
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
func (B *base) getBbsInfo() error {
	var err error
	model := new(model.Bbs)
	if B.bbsInfo, err = model.GetOne(B.bbsID); err == nil {
		B.boardID = B.bbsInfo.BoardID
	}
	return err
}

//getBoardInfo 读取公告信息
func (B *base) getBoardInfo() error {
	var err error
	model := new(model.Board)
	B.boardInfo, err = model.GetOne(B.boardID)
	return err
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
