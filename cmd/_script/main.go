package main

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/ccyun/GoApp/application"
	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/model"
)

//FeedMsg msg结构
type FeedMsg struct {
	BoardID  uint64 `orm:"column(board_id)"`
	BbsID    uint64 `orm:"column(bbs_id)"`
	FeedID   uint64 `orm:"column(feed_id)"`
	FeedType string `orm:"column(feed_type)"`
	SendType string `orm:"column(send_type)"`
	SendTo   uint64 `orm:"column(send_to)"`
}
type msgData struct {
	BoardID  uint64
	BbsID    uint64
	FeedID   uint64
	UserIDs  []uint64
	OrgIDs   []uint64
	FeedType string
}

var mData map[uint64]msgData

func main() {
	o := orm.NewOrm()
	i := 0
	for true {
		var data []FeedMsg
		o.Raw("select board_id,bbs_id,feed_id,feed_type,send_type,send_to from bbs_msg where feed_id=(select feed_id from bbs_msg where send_type='user' or send_type='org' group by feed_id order by feed_id asc limit ?,1)", i).QueryRows(&data)

		if len(data) == 0 {
			break
		} else {
			var d msgData
			for _, v := range data {
				d.BoardID = v.BoardID
				d.BbsID = v.BbsID
				d.FeedID = v.FeedID
				d.FeedType = v.FeedType
				if v.SendType == "user" {
					d.UserIDs = append(d.UserIDs, v.SendTo)
				} else if v.SendType == "org" {
					d.OrgIDs = append(d.OrgIDs, v.SendTo)
				}
			}

			if len(d.OrgIDs) > 0 {
				a := new(httpcurl.UMS)
				userIDs, err := a.GetAllUserIDsByOrgIDs("00000", d.OrgIDs)
				if err != nil {
					log.Println(err)
				}
				if len(userIDs) > 0 {
					d.UserIDs = append(d.UserIDs, userIDs[0:]...)
				}
			}
			d.UserIDs = userIDsUnique(d.UserIDs)
			userIDstr, _ := json.Marshal(d.UserIDs)
			o.Raw("UPDATE bbs_bbs SET send_user_ids = ?,msg_count=? where id=?", string(userIDstr), len(d.UserIDs), d.BbsID).Exec()
			saveHbase(d)
		}
		i++
	}
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

//保存到hbase
func saveHbase(data msgData) {
	feed := new(model.Feed)
	d := model.Feed{
		ID:       data.FeedID,
		BoardID:  data.BoardID,
		BbsID:    data.BbsID,
		FeedType: data.FeedType,
	}
	if err := feed.SaveHbase(data.UserIDs, d); err != nil {
		log.Println(err)
	}

}
