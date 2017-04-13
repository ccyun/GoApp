package adapter

import (
	"bbs_server/application/function"
	"bbs_server/application/model"

	"github.com/astaxie/beego/orm"
)

//base 任务处理适配器（基类）
type base struct {
	o                          orm.Ormer
	taskID                     uint64
	siteID                     uint64
	customerCode               string
	action                     string
	bbsID                      uint64
	bbsInfo                    model.Bbs
	category                   string
	feedType                   string
	bbsTaskInfo                model.BbsTask
	PublishScope               map[string][]uint64
	PublishScopeuserLoginNames []string
	boardID                    uint64
	boardInfo                  model.Board
	feedID                     uint64
	userIDs                    []uint64
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
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *base) GetPublishScopeUsers() error {
	return nil
}

//CreateFeed 创建Feed
func (B *base) CreateFeed() error {
	return nil
}

//CreateRelation 创建接收者关系
func (B *base) CreateRelation() error {
	return nil
}

//CreateUnread 创建未处理数
func (B *base) CreateUnread() error {
	var (
		err        error
		InsertData []model.Unread
		userIDs    []uint64
	)
	db := new(model.Unread)
	if userIDs, err = db.GetUserIDs(B.siteID, B.boardID, B.feedType); err != nil {
		return err
	}
	if len(userIDs) > 0 {
		if _, err = B.o.QueryTable(db).Filter("SiteID", B.siteID).Filter("BoardID", B.boardID).Filter("FeedType", B.feedType).Filter("UserID__in", function.SliceIntersect(B.userIDs, userIDs).Uint64()).Update(orm.Params{
			"UnreadCount": orm.ColValue(orm.ColAdd, 1),
		}); err != nil {
			return err
		}
	}
	for _, userID := range function.SliceDiff(B.userIDs, userIDs).Uint64() {
		InsertData = append(InsertData, model.Unread{
			SiteID:      B.siteID,
			BoardID:     B.boardID,
			FeedType:    B.feedType,
			UnreadCount: 1,
			UserID:      userID,
		})
	}
	if len(InsertData) > 0 {
		_, err = B.o.InsertMulti(100000, InsertData)
	}
	return err

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
