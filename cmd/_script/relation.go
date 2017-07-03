package main

import (
	"bbs_server/application/model"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type msger struct {
	FeedID   uint64 `orm:"column(feed_id)"`
	FeedType string `orm:"column(feed_type)"`
	SendType string `orm:"column(send_type)"`
	SendTo   uint64 `orm:"column(send_to)"`
}

type msgData struct {
	BoardID   uint64
	DiscussID uint64
	BbsID     uint64
	FeedID    uint64
	UserIDs   []uint64
	FeedType  string
}

func (T *task) getPublishScope() error {
	msgList := []msger{}
	if _, err := o.Raw("select feed_id,feed_type,send_type,send_to from bbs_msg2 where bbs_id=?", T.bbsInfo.ID).QueryRows(&msgList); err != nil {
		return err
	}
	T.publishScopeUserIDs = make(map[uint64][]uint64)
	orgIDs := make(map[uint64][]uint64)
	for _, v := range msgList {
		if v.SendType == "user" {
			T.publishScopeUserIDs[v.FeedID] = append(T.publishScopeUserIDs[v.FeedID], v.SendTo)
		} else if v.SendType == "org" {
			orgIDs[v.FeedID] = append(orgIDs[v.FeedID], v.SendTo)
		} else if v.SendType == "discuss" {
			discussInfo, err := ucc.GetDiscussInfo(T.bbsInfo.UserID, T.bbsInfo.DiscussID)
			if err != nil {
				log.Println(T.bbsInfo.DiscussID)
				log.Printf("bbs_id:%d,GetDiscussInfo error:%s", T.bbsInfo.ID, err.Error())
				return err
			}
			T.publishScopeUserIDs[v.FeedID] = discussInfo.ValidMemberIDs
		}
		T.publishScopeUserIDs[v.FeedID] = userIDsUnique(T.publishScopeUserIDs[v.FeedID])
	}
	for feedID, orgIDs := range orgIDs {
		users, err := ums.GetAllUserIDsByOrgIDs("00000", orgIDs)
		if err != nil {
			log.Println(orgIDs)
			log.Printf("bbs_id:%d,GetAllUserIDsByOrgIDs error:%s", T.bbsInfo.ID, err.Error())
		} else {
			T.publishScopeUserIDs[feedID] = append(T.publishScopeUserIDs[feedID], users...)
		}
	}
	return nil
}

//Create 创建消息
func createRelation(msgData model.Msg, userIDs []uint64, taskStatus uint8) error {
	db := orm.NewOrm()
	db.Using("msg")
	sql := "insert into  bbs_msg (`site_id`,`board_id`,`discuss_id`,`bbs_id`,`feed_type`,`feed_id`,`user_id`,`user_org_id`,`task_status`,`is_read`,`created_at`) values"
	values := []string{}
	startIndex := 0
	userCount := len(userIDs)
	for true {
		if startIndex >= userCount {
			break
		}
		endIndex := startIndex + 10000
		if endIndex > userCount {
			endIndex = userCount
		}
		for _, u := range userIDs[startIndex:endIndex] {
			values = append(values, fmt.Sprintf("(%d,%d,%d,%d,'%s',%d,%d,%d,%d,%d,%d)", msgData.SiteID, msgData.BoardID, msgData.DiscussID, msgData.BbsID, msgData.FeedType, msgData.FeedID, u, 0, taskStatus, 1, msgData.CreatedAt))
		}
		if _, err := db.Raw(sql + strings.Join(values, ",")).Exec(); err != nil {
			return err
		}
		values = []string{}
		startIndex += 10000
	}
	return nil
}

//userIDsUnique 去重复
func userIDsUnique(data []uint64) []uint64 {
	_data := make(map[uint64]bool)
	for _, v := range data {
		_data[v] = true
	}
	data = []uint64{}
	for v := range _data {
		data = append(data, v)
	}
	return data
}

func updateMsgOrgID() error {
	db := orm.NewOrm()
	db.Using("msg")
	for true {
		userIDs := []uint64{}
		msgList := []model.Msg{}
		if _, err := db.Raw("select user_id from bbs_msg where user_org_id=0 group by user_id limit 100").QueryRows(&msgList); err != nil {
			return err
		}
		if len(msgList) == 0 {
			break
		}
		uu := []string{}
		for _, u := range msgList {
			uu = append(uu, strconv.FormatUint(u.UserID, 10))
			userIDs = append(userIDs, u.UserID)
		}
		userList, err := ums.GetUsersDetail("00000", userIDs, true)
		if err != nil {
			log.Printf("GetUsersDetail error:%s", err.Error())
			return err
		}
		if len(userList) > 0 {
			for _, u := range userList {
				if _, err := db.Raw(fmt.Sprintf("UPDATE bbs_msg SET user_org_id=%d where user_id=%d", u.OrganizationID, u.UserID)).Exec(); err != nil {
					return err
				}
			}
		}
		if _, err := db.Raw("UPDATE bbs_msg SET user_org_id = 18446744073709551615 where user_id in(" + strings.Join(uu, ",") + ") and user_org_id=0").Exec(); err != nil {
			return err
		}
	}

	if _, err := db.Raw("UPDATE bbs_msg SET user_org_id = 0 where  user_org_id=18446744073709551615").Exec(); err != nil {
		return err
	}
	return nil
}

func updateMsgTaskStatus() error {
	db := orm.NewOrm()
	db.Using("msg")

	maxID := uint64(0)
	for true {
		replyList := []model.Msg{}
		if _, err := o.Raw(fmt.Sprintf("select user_id,bbs_id,status from bbs_bbs_task_reply where status=1 and user_id>%d order by user_id asc limit 1000", maxID)).QueryRows(&replyList); err != nil {
			return err
		}
		if len(replyList) == 0 {
			break
		}
		for _, reply := range replyList {
			if _, err := db.Raw(fmt.Sprintf("UPDATE bbs_msg SET task_status=1 where feed_type='task' and bbs_id=%d and user_id=%d", reply.BbsID, reply.UserID)).Exec(); err != nil {
				return err
			}
			if reply.UserID > maxID {
				maxID = reply.UserID
			}
		}
	}
	return nil
}
