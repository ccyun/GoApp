package main

import (
	"bbs_server/application/function"
	"encoding/json"
	"fmt"
	"strings"
)

type replyer struct {
	ID           uint64               `orm:"column(id)"`
	SiteID       uint64               `orm:"column(site_id)"`
	BoardID      uint64               `orm:"column(board_id)"`
	BbsID        uint64               `orm:"column(bbs_id)"`
	SubTaskID    uint64               `orm:"column(sub_task_id)"`
	ReplyID      uint64               `orm:"column(reply_id)"`
	UserID       uint64               `orm:"column(user_id)"`
	DataStr      string               `orm:"column(data)"`
	Data         map[string]oldDataer `orm:"-"`
	CreatedAt    uint64               `orm:"column(created_at)"`
	Status       int8                 `orm:"column(status)"`
	AuditOpinion string               `orm:"column(audit_opinion)"`
	AuditUserID  uint64               `orm:"column(audit_user_id)"`
}
type tager struct {
	ID        uint64 `orm:"column(id)"`
	SiteID    uint64 `orm:"column(site_id)"`
	BbsID     uint64 `orm:"column(bbs_id)"`
	ReplyID   uint64 `orm:"column(reply_id)"`
	TagID     uint64 `orm:"column(tag_id)"`
	TagName   string `orm:"column(tag_name)"`
	TagCode   string `orm:"column(tag_code)"`
	TagEnumID uint64 `orm:"column(tag_enum_id)"`
	TagValue  string `orm:"column(tag_value)"`
}
type commenter struct {
	Top  float64 `json:"top"`
	Left float64 `json:"left"`
	Text string  `json:"text"`
}

type oldDataer struct {
	newDataer
	Status       uint8       `json:"status"`
	AuditOpinion string      `json:"audit_opinion"`
	Comments     []commenter `json:"comments"`
}

type newDataer struct {
	Src         string `json:"src"`
	Description string `json:"description"`
	ShootTime   uint64 `json:"shoot_time"`
	Address     string `json:"address"`
	Location    geo    `json:"location"`
}

type geo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (T *task) handleTaskReply() error {
	if T.bbsInfo.Category != "task" {
		return nil
	}
	var replyList []replyer
	if _, err := o.Raw("select reply.*,task.type from bbs_bbs_task_reply_sub reply INNER JOIN bbs_bbs_task_sub task on task.id=reply.sub_task_id where task.type='image' and reply.bbs_id=?", T.bbsInfo.ID).QueryRows(&replyList); err != nil {
		return err
	}
	if len(replyList) > 0 {
		values := []string{}
		sql := "insert into  bbs_bbs_task_reply_sub_ext (`site_id`,`board_id`,`bbs_id`,`sub_task_id`,`reply_id`,`sub_task_reply_id`,`value`,`audit_opinion`,`comments`,`status`) values"
		for _, item := range replyList {
			if err := json.Unmarshal([]byte(item.DataStr), &item.Data); err == nil {
				for _, v := range item.Data {
					newData := newDataer{
						Src:         v.Src,
						Description: v.Description,
						ShootTime:   v.ShootTime,
						Address:     v.Address,
						Location: geo{
							Latitude:  v.Location.Latitude,
							Longitude: v.Location.Longitude,
						},
					}
					newData2, _ := json.Marshal(newData)
					if len(v.Comments) == 0 {
						v.Comments = []commenter{}
					}
					comments, _ := json.Marshal(v.Comments)
					values = append(values, fmt.Sprintf("(%d,%d,%d,%d,%d,%d,\"%s\",\"%s\",\"%s\",%d)", item.SiteID, item.BoardID, item.BbsID, item.SubTaskID, item.ReplyID, item.ID, function.MysqlEscapeString(string(newData2)), v.AuditOpinion, function.MysqlEscapeString(string(comments)), v.Status))
				}
			}
		}
		if len(values) > 0 {
			if _, err := o.Raw(sql + strings.Join(values, ",")).Exec(); err != nil {
				return err
			}
		}

	}
	return nil
}

func (T *task) handleTaskReplyTags() error {
	var replyList []struct {
		ID     uint64 `orm:"column(id)"`
		UserID uint64 `orm:"column(user_id)"`
	}
	if _, err := o.Raw("select id,user_id from bbs_bbs_task_reply where bbs_id=?", T.bbsInfo.ID).QueryRows(&replyList); err != nil {
		return err
	}
	if len(replyList) == 0 {
		return nil
	}
	userIDs := []uint64{}
	for _, item := range replyList {
		userIDs = append(userIDs, item.UserID)
	}
	tagList, err := ums.GetUserTags("000000", T.bbsInfo.SiteID, userIDs)
	if err != nil {
		return err
	}
	values := []string{}
	sql := "insert into  bbs_bbs_task_reply_tags (`site_id`,`bbs_id`,`reply_id`,`tag_id`,`tag_name`,`tag_code`,`tag_enum_id`,`tag_value`) values"
	for _, item := range replyList {
		if tags, ok := tagList[item.UserID]; ok {
			for _, tag := range tags {
				values = append(values, fmt.Sprintf("(%d,%d,%d,%d,\"%s\",\"%s\",%d,\"%s\")", T.bbsInfo.SiteID, T.bbsInfo.ID, item.ID, tag.TagID, function.MysqlEscapeString(tag.TagName), function.MysqlEscapeString(tag.TagCode), tag.TagEnumID, function.MysqlEscapeString(tag.TagValue)))
			}
		}
	}
	if len(values) > 0 {
		if _, err := o.Raw(sql + strings.Join(values, ",")).Exec(); err != nil {
			return err
		}
	}
	return nil
}
