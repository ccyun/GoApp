package main

import (
	_ "bbs_server/application"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"encoding/json"
	"log"
	"runtime"
	"sync"

	"time"

	"github.com/astaxie/beego/orm"
)

var (
	o      orm.Ormer
	ums    *httpcurl.UMS
	ucc    *httpcurl.UCC
	bbsIDs [][]uint64
)

type task struct {
	bbsInfo             model.Bbs
	taskInfo            model.BbsTask
	category            string
	publishScopeUserIDs map[uint64][]uint64
	taskStatus          uint8
}

type feedDataer struct {
	Status int8 `json:"status"`
}

// var mData map[uint64]msgData
func init() {
	o = orm.NewOrm()
	ums = new(httpcurl.UMS)
	ucc = new(httpcurl.UCC)
}

func main() {
	if err := getBbsIDs(); err != nil {
		log.Println(err)
		return
	}
	db := new(DB)
	log.Println(db.alterTable())
	if err := runTask(); err != nil {
		log.Println(err)
		return
	}

	//更新用户组织ID
	if err := updateMsgOrgID(); err != nil {
		log.Println(err)
	}
	if err := updateMsgTaskStatus(); err != nil {
		log.Println(err)
	}
	log.Println(db.clearTable())
}

func runTask() error {
	var w sync.WaitGroup
	runtime.GOMAXPROCS(8)
	// bbsIDs = [][]uint64{[]uint64{
	// 	50001145,
	// 	50001162,
	// 	50001599,
	// 	50001256,
	// 	50001260,
	// 	50001212,
	// 	50001381,
	// 	50001272,
	// 	50001389,
	// 	50001390,
	// 	50001352,
	// 	50001577,
	// }}
	w.Add(len(bbsIDs))
	for _, ids := range bbsIDs {
		go func(ids []uint64) {
			for _, id := range ids {
				handleBbs(id)
			}
			w.Done()
		}(ids)
	}
	w.Wait()
	return nil
}

func handleBbs(id uint64) error {
	log.Println(id)
	T := new(task)
	if err := T.getInfo(id); err != nil {
		log.Println(err)
		return err
	}
	if err := T.getPublishScope(); err != nil {
		log.Println(err)
		return err
	}
	if err := T.handleFeed(); err != nil {
		log.Println(err)
		return err
	}
	if err := T.handleTaskReply(); err != nil {
		log.Println(err)
		return err
	}
	if err := T.handleTaskReplyTags(); err != nil {
		log.Println(err)
		return err
	}
	if err := T.handleSubTaskText(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getBbsIDs() error {
	var bbsData []model.Bbs
	if _, err := o.Raw("select id from bbs_bbs where status=1 and is_deleted=0 and (discuss_id=0 or (discuss_id>0 and type='default')) order by id asc").QueryRows(&bbsData); err != nil {
		return err
	}
	pageSize := (len(bbsData) / 20) + 1
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
	if err = o.Raw("select id,site_id,board_id,title,description,discuss_id,category,comment_enabled,type,publish_at,attachments,user_id from bbs_bbs where id=? order by id asc limit 1", id).QueryRow(&T.bbsInfo); err != nil {
		return err
	}
	if T.bbsInfo.AttachmentsString != "" {
		if err = json.Unmarshal([]byte(T.bbsInfo.AttachmentsString), &T.bbsInfo.Attachments); err != nil {
			log.Println(T.bbsInfo)
		}
		if len(T.bbsInfo.Attachments) > 0 {
			if thumb, ok := T.bbsInfo.Attachments[0]["url"]; ok {
				T.bbsInfo.Thumb = thumb
			}
		}
	}
	T.category = T.bbsInfo.Category
	if T.category == "task" {
		if err = o.Raw("select end_time,allow_expired,is_close from bbs_bbs_task where bbs_id=?", T.bbsInfo.ID).QueryRow(&T.taskInfo); err != nil {
			return err
		}
		if T.taskInfo.IsClose == 1 || (T.taskInfo.AllowExpired == 0 && T.taskInfo.EndTime > uint64(time.Now().UnixNano()/1e6)) {
			T.taskStatus = 1
		}
	}
	return nil
}
