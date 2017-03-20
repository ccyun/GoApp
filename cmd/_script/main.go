package main

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/astaxie/beego/orm"
	_ "github.com/ccyun/GoApp/application"
	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/model"
	"github.com/ccyun/GoApp/application/module/feed"
)

var (
	o      orm.Ormer
	ums    *httpcurl.UMS
	bbsIDs [][]uint64
)

type task struct {
	bbsInfo   model.Bbs
	boardInfo model.Board
	taskInfo  model.BbsTask
	msgInfo   msger
	category  string
}
type msger struct {
	FeedID   uint64 `orm:"column(feed_id)"`
	FeedType string `orm:"column(feed_type)"`
	SendType string `orm:"column(send_type)"`
	SendTo   uint64 `orm:"column(send_to)"`
}

type feedDataer struct {
	Status int8 `json:"status"`
}

type msgData struct {
	BoardID   uint64
	DiscussID uint64
	BbsID     uint64
	FeedID    uint64
	UserIDs   []uint64
	FeedType  string
}

// var mData map[uint64]msgData
func init() {
	o = orm.NewOrm()
	ums = new(httpcurl.UMS)
}

func main() {
	if err := getBbsIDs(); err != nil {
		log.Println(err)
		return
	}
	runTask()
}

func runTask() {
	var w sync.WaitGroup
	w.Add(len(bbsIDs))
	for _, ids := range bbsIDs {
		go func(ids []uint64) {
			for _, id := range ids {
				T := new(task)
				if err := T.getInfo(id); err != nil {
					log.Println(err)
					return
				}
				if err := T.handleBbs(); err != nil {
					log.Println(err)
					return
				}
				if err := T.handleFeed(); err != nil {
					log.Println(err)
					return
				}
				if err := T.handleRelation(); err != nil {
					log.Println(err)
					return
				}
			}
			w.Done()
		}(ids)

	}
	w.Wait()

}

func getBbsIDs() error {
	var bbsData []model.Bbs
	if _, err := o.Raw("select id from bbs_bbs order by id asc").QueryRows(&bbsData); err != nil {
		return err
	}
	pageSize := (len(bbsData) / 5) + 1
	var (
		tempIDs []uint64
	)
	for j, v := range bbsData {
		tempIDs = append(tempIDs, v.ID)
		if (j+1)%pageSize == 0 {
			bbsIDs = append(bbsIDs, tempIDs)
			tempIDs = []uint64{}
		}
	}
	bbsIDs = append(bbsIDs, tempIDs)
	return nil
}
func (T *task) getInfo(id uint64) error {
	var err error
	if err = o.Raw("select id,board_id,title,description,discuss_id,category,comment_enabled,type,publish_at,publish_scope,attachments,user_id from bbs_bbs where id=? order by id asc limit 1", id).QueryRow(&T.bbsInfo); err != nil {
		return err
	}
	if T.bbsInfo.AttachmentsString != "" {
		if err = json.Unmarshal([]byte(T.bbsInfo.AttachmentsString), &T.bbsInfo.Attachments); err != nil {
			return err
		}
	}

	var publishScope map[string][]string
	if err = json.Unmarshal([]byte(T.bbsInfo.PublishScopeString), &publishScope); err != nil {
		return err
	}

	if _, ok := publishScope["user_ids"]; ok {
		for _, u := range publishScope["user_ids"] {
			uu, _ := strconv.Atoi(u)
			T.bbsInfo.PublishScopeUserIDsArr = append(T.bbsInfo.PublishScopeUserIDsArr, uint64(uu))
		}
	}

	log.Println(T.bbsInfo.ID)
	T.category = T.bbsInfo.Category
	msgData := []msger{}
	if _, err = o.Raw("select send_type,send_to from bbs_msg where feed_type=? and bbs_id=?", T.category, T.bbsInfo.ID).QueryRows(&msgData); err != nil {
		return err
	}
	T.getPublishScopeUserIDsArr(msgData)
	if T.category == "task" {
		if err = o.Raw("select end_time,allow_expired from bbs_bbs_task where bbs_id=?", T.bbsInfo.ID).QueryRow(&T.taskInfo); err != nil {
			return err
		}
	}
	return nil
}

func (T *task) getPublishScopeUserIDsArr(data []msger) error {
	orgIDs := []uint64{}
	for _, v := range data {
		if v.SendType == "user" {
			T.bbsInfo.PublishScopeUserIDsArr = append(T.bbsInfo.PublishScopeUserIDsArr, v.SendTo)
		} else if v.SendType == "org" {
			orgIDs = append(orgIDs, v.SendTo)
		}
	}
	if len(orgIDs) > 0 {
		userIDs, err := ums.GetAllUserIDsByOrgIDs("00000", orgIDs)
		if err != nil {
			return err
		}
		if len(userIDs) > 0 {
			T.bbsInfo.PublishScopeUserIDsArr = append(T.bbsInfo.PublishScopeUserIDsArr, userIDs[0:]...)
		}
	}
	T.bbsInfo.PublishScopeUserIDsArr = T.userIDsUnique(T.bbsInfo.PublishScopeUserIDsArr)
	return nil
}

//handleFeed 处理关系
func (T *task) handleBbs() error {
	if T.bbsInfo.DiscussID != 0 && T.bbsInfo.Type == "default" {
		return nil
	}
	userIDstr, _ := json.Marshal(T.bbsInfo.PublishScopeUserIDsArr)
	T.bbsInfo.PublishScopeUserIDs = string(userIDstr)
	userCount := len(T.bbsInfo.PublishScopeUserIDsArr)
	if userCount > 0 {
		_, err := o.Raw("UPDATE bbs_bbs SET publish_scope_user_ids = ?,msg_count=? where id=?", T.bbsInfo.PublishScopeUserIDs, userCount, T.bbsInfo.ID).Exec()
		return err
	}
	return nil
}

//handleFeed 处理关系
func (T *task) handleFeed() error {
	feedData := model.FeedData{
		Title:          T.bbsInfo.Title,
		Description:    T.bbsInfo.Description,
		CreatedAt:      T.bbsInfo.PublishAt,
		UserID:         T.bbsInfo.UsesID,
		Type:           T.bbsInfo.Type,
		Category:       T.category,
		CommentEnabled: T.bbsInfo.CommentEnabled,
	}
	if T.category == "bbs" {
		if thumb, ok := T.bbsInfo.Attachments[0]["url"]; ok {
			feedData.Thumb = thumb
		}
	} else if T.category == "task" {
		feedData.EndTime = T.taskInfo.EndTime
		feedData.AllowExpired = T.taskInfo.AllowExpired
	}
	feedList := []model.Feed{}
	if _, err := o.Raw("select id,bbs_id,feed_type,data,created_at from bbs_feed where bbs_id=?", T.bbsInfo.ID).QueryRows(&feedList); err != nil {
		return err
	}

	for _, v := range feedList {
		switch v.FeedType {
		case "task":
			feedData.Status = feed.BbsTaskStatus
		case "taskReply":
			feedData.Status = feed.BbsTaskReplyStatus
			feedData.CreatedAt = v.CreatedAt
		case "taskClose":
			feedData.Status = feed.BbsTaskCloseStatus
			feedData.CreatedAt = v.CreatedAt
		case "taskAudit":
			fD := feedDataer{}
			if err := json.Unmarshal([]byte(v.Data), &fD); err != nil {
				return err
			}
			feedData.Status = fD.Status
			feedData.CreatedAt = v.CreatedAt
		}
		s, _ := json.Marshal(feedData)
		if _, err := o.Raw("UPDATE bbs_feed SET data = ? where id=?", string(s), v.ID).Exec(); err != nil {
			return err
		}

	}
	return nil
}

//handleRelation 处理关系

func (T *task) handleRelation() error {
	if T.bbsInfo.DiscussID > 0 && T.bbsInfo.Type == "preview" {
		return nil
	}
	msgList := []msger{}
	data := map[uint64](map[string](map[string][]uint64)){}
	if _, err := o.Raw("select feed_id,feed_type,send_type,send_to from bbs_msg where bbs_id=?", T.bbsInfo.ID).QueryRows(&msgList); err != nil {
		return err
	}
	for _, v := range msgList {
		if _, ok := data[v.FeedID]; !ok {
			data[v.FeedID] = make(map[string](map[string][]uint64))
		}
		if _, ok := data[v.FeedID][v.FeedType]; !ok {
			data[v.FeedID][v.FeedType] = make(map[string][]uint64)
		}
		data[v.FeedID][v.FeedType][v.SendType] = append(data[v.FeedID][v.FeedType][v.SendType], v.SendTo)
	}

	for key, value := range data {
		for k, item := range value {
			userIDs := []uint64{}
			for kk, v := range item {
				if kk == "user" {
					userIDs = append(userIDs, v...)
				} else if kk == "org" {
					users, err := ums.GetAllUserIDsByOrgIDs("00000", v)
					if err != nil {
						return err
					}
					userIDs = append(userIDs, users...)
				}
			}
			if err := T.saveHbase(key, k, userIDs); err != nil {
				return err
			}
		}
	}
	return nil
}

//userIDsUnique 去重复
func (T *task) userIDsUnique(data []uint64) []uint64 {
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

//保存到hbase
func (T *task) saveHbase(feedID uint64, feedType string, UserIDs []uint64) error {
	feed := new(model.Feed)
	d := model.Feed{
		ID:       feedID,
		BoardID:  T.bbsInfo.BoardID,
		BbsID:    T.bbsInfo.ID,
		FeedType: feedType,
	}
	return feed.SaveHbase(UserIDs, d, T.bbsInfo.DiscussID)
}
