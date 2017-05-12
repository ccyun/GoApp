package adapter

import (
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//base 任务处理适配器（基类）
type base struct {
	o                          orm.Ormer
	nowTime                    uint64
	taskID                     uint64
	siteID                     uint64
	customerCode               string
	action                     string
	bbsID                      uint64
	bbsInfo                    model.Bbs
	category                   string
	feedType                   string
	bbsTaskInfo                model.BbsTask
	boardID                    uint64
	boardInfo                  model.Board
	feedID                     uint64
	userList                   []httpcurl.UMSUser
	userIDs                    []uint64
	PublishScopeuserLoginNames []string
	attachmentsBase64          string
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
func (B *base) NewTask(task model.Queue) error {
	B.taskID = task.ID
	B.siteID = task.SiteID
	B.customerCode = task.CustomerCode
	B.action = task.Action
	B.nowTime = uint64(time.Now().UnixNano() / 1e6)
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *base) GetPublishScopeUsers() error {
	var err error
	if len(B.userIDs) > 0 {
		return fmt.Errorf("GetPublishScopeUsers error not found UNReply Users")
	}
	if B.userList, err = new(httpcurl.UMS).GetUsersDetail(B.customerCode, B.userIDs, true); err != nil {
		return err
	}
	for _, v := range B.userList {
		B.PublishScopeuserLoginNames = append(B.PublishScopeuserLoginNames, v.LoginName)
	}
	return nil
}

//CreateFeed 创建Feed
func (B *base) CreateFeed() error {
	return nil
}

//CreateRelation 创建接收者关系
func (B *base) CreateRelation() error {
	var (
		ackReadUserID     uint64
		defaultReadStatus uint8
	)
	if B.boardInfo.DiscussID > 0 {
		if B.bbsInfo.Type == "preview" {
			return nil
		}
		ackReadUserID = B.bbsInfo.UserID
	}
	msgData := model.Msg{
		SiteID:    B.siteID,
		BoardID:   B.boardID,
		DiscussID: B.boardInfo.DiscussID,
		BbsID:     B.bbsID,
		FeedType:  B.feedType,
		FeedID:    B.feedID,
		CreatedAt: B.nowTime,
	}
	return new(model.Msg).Create(msgData, B.userList, defaultReadStatus, ackReadUserID)
}

//CreateUnread 创建未处理数
func (B *base) CreateUnread() error {
	return nil
}

//UpdateStatus 更新状态
func (B *base) UpdateStatus() error {
	return nil
}

//SendMsg 发送消息
func (B *base) SendMsg() error {
	return nil
}

///////////////////////////////////////////////公共方法//////////////////////////////////////////////////////////////////////
//getBbsInfo 读取公告信息
func (B *base) getBbsInfo() error {
	var err error
	model := new(model.Bbs)
	if B.bbsInfo, err = model.GetOne(B.bbsID); err == nil {
		B.boardID = B.bbsInfo.BoardID
		B.category = B.bbsInfo.Category
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
