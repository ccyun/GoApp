package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//Bbs 任务表结构
type Bbs struct {
	base
	ID                     uint64              `orm:"column(id)"`
	SiteID                 uint64              `orm:"column(site_id)"`
	BoardID                uint64              `orm:"column(board_id)"`
	Title                  string              `orm:"column(title)"`
	Description            string              `orm:"column(description)"`
	Content                string              `orm:"column(content)"`
	PublishScopeString     string              `orm:"column(publish_scope)"`
	PublishScope           PublishScoper       `orm:"-"`
	PublishScopeUserIDs    string              `orm:"column(publish_scope_user_ids)"`
	PublishScopeUserIDsArr []uint64            `orm:"-"`
	MsgCount               uint64              `orm:"column(msg_count)"`
	AttachmentsString      string              `orm:"column(attachments)"`
	Attachments            []map[string]string `orm:"-"`
	UsesID                 uint64              `orm:"column(user_id)"`
	CreatedAt              uint64              `orm:"column(created_at)"`
	PublishAt              uint64              `orm:"column(publish_at)"`
	ModifiedAt             uint64              `orm:"column(modified_at)"`
	SetTimer               uint64              `orm:"column(set_timer)"`
	Category               string              `orm:"column(category)"`
	Type                   string              `orm:"column(type)"`
	IsDeleted              string              `orm:"column(is_deleted)"`
	Status                 uint8               `orm:"column(status)"`
	CommentEnabled         uint8               `orm:"column(comment_enabled)"`
}

//PublishScoper 广播发布范围
type PublishScoper struct {
	GroupIDs   []uint64 `json:"group_ids"`
	UserIDs    []uint64 `json:"user_ids"`
}

//TableName 表名
func (B *Bbs) TableName() string {
	return "bbs"
}

//GetOne 读取单条数据
func (B *Bbs) GetOne(ID uint64) (Bbs, error) {
	bbsInfo := Bbs{ID: ID}
	// c := redis.NewCache(fmt.Sprintf("D%d%s", B.siteID, B.TableName()), "GetOne", ID)
	// if c.Get(&bbsInfo) == true {
	// 	return bbsInfo, nil
	// }
	if err := o.Read(&bbsInfo); err != nil {
		return Bbs{}, err
	}

	data, err := B.afterSelectHandle([]Bbs{bbsInfo})
	if err != nil {
		return Bbs{}, err
	}
	// c.Set(data[0])
	return data[0], nil
}

//Update 更新数据
func (B *Bbs) Update(bbs Bbs, field ...string) error {
	bbs.ModifiedAt = uint64(time.Now().UnixNano() / 1e6)
	num, err := o.Update(&bbs, field...)
	if num == 0 {
		err = orm.ErrNoRows
	}
	return err
}

//afterSelectHandle 查询结果处理
func (B *Bbs) afterSelectHandle(data []Bbs) ([]Bbs, error) {
	var err error
	for key, item := range data {
		//处理发布范围
		item.PublishScope, item.PublishScopeUserIDsArr, err = B.publishScopeHandle(item.PublishScopeString, item.PublishScopeUserIDs)
		if item.AttachmentsString != "" {
			err = json.Unmarshal([]byte(item.AttachmentsString), &item.Attachments)
		}
		if err == nil {
			data[key] = item
		}
	}
	return data, err
}

//publishScopeHandle 处理发布范围
func (B *Bbs) publishScopeHandle(publishScopeString, publishScopeUserIDs string) (PublishScoper, []uint64, error) {
	var data PublishScoper
	var publishScope map[string][]string
	var err error
	if err = json.Unmarshal([]byte(publishScopeString), &publishScope); err != nil {
		return data, nil, err
	}
	for k, r := range publishScope {
		for _, v := range r {
			id, _ := strconv.Atoi(v)
			if id > 0 {
				switch k {
				case "group_ids":
					data.GroupIDs = append(data.GroupIDs, uint64(id))
				case "user_ids":
					data.UserIDs = append(data.UserIDs, uint64(id))
			}
		}
	}
	publishUser := []uint64{}
	if err = json.Unmarshal([]byte(publishScopeUserIDs), &publishUser); err != nil {
		return data, nil, err
	}
	return data, publishUser, err
}
