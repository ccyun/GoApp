package model

import "fmt"

//Unread 任务表结构
type Unread struct {
	base
	ID          uint64 `orm:"column(id)"`
	SiteID      uint64 `orm:"column(site_id)"`
	BoardID     uint64 `orm:"column(board_id)"`
	UserID      uint64 `orm:"column(user_id)"`
	UnreadCount uint64 `orm:"column(unread_count)"`
}

//TableName 表名
func (U *Unread) TableName() string {
	return "unread2"
}

//IncrCount 未读计数+1
func (U *Unread) IncrCount(siteID, boardID uint64, userIDs []uint64) error {
	if len(userIDs) == 0 {
		return nil
	}
	sql := "INSERT INTO " + DBPrefix + U.TableName() + " (site_id,board_id,user_id,unread_count) "
	for k, userID := range userIDs {
		if k == 0 {
			sql += "VALUES"
		} else {
			sql += ","
		}
		sql += fmt.Sprintf("(%d,%d,%d,%d)", siteID, boardID, userID, 1)
	}
	sql += " ON DUPLICATE KEY update unread_count=unread_count+1"
	_, err := o.Raw(sql).Exec()
	return err

}
