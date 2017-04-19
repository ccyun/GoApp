package model

import "github.com/astaxie/beego/orm"

//BbsTaskUnreplyCount 任务未反馈数
type BbsTaskUnreplyCount struct {
	base
	ID           uint64 `orm:"column(id)"`
	SiteID       uint64 `orm:"column(site_id)"`
	BoardID      uint64 `orm:"column(board_id)"`
	UserID       uint64 `orm:"column(user_id)"`
	UnreplyCount uint64 `orm:"column(unreply_count)"`
}

//TableName 表名
func (B *BbsTaskUnreplyCount) TableName() string {
	return "bbs_task_unreply_count"
}

//GetUserIDs 查询未反馈用户列表
func (B *BbsTaskUnreplyCount) GetUserIDs(siteID, boardID uint64) ([]uint64, error) {
	var (
		data    []*BbsTaskUnreplyCount
		userIDs []uint64
	)
	if _, err := o.QueryTable(B).Filter("SiteID", siteID).Filter("BoardID", boardID).Limit(-1).All(&data, "UserID"); err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	for _, v := range data {
		userIDs = append(userIDs, v.UserID)
	}
	return userIDs, nil
}
