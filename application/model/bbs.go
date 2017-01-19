package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//Bbs 任务表结构
type Bbs struct {
	ID                  uint64        `orm:"column(id)"`
	SiteID              uint64        `orm:"column(site_id)"`
	BoardID             uint64        `orm:"column(board_id)"`
	Title               string        `orm:"column(title)"`
	Description         string        `orm:"column(description)"`
	Content             string        `orm:"column(content)"`
	PublishScopeString  string        `orm:"column(publish_scope)"`
	PublishScope        PublishScoper `orm:"-"`
	PublishScopeUserIDs string        `orm:"column(publish_scope_user_ids)"`
	MsgCount            uint8         `orm:"column(msg_count)"`
	Attachments         string        `orm:"column(attachments)"`
	UsesID              uint64        `orm:"column(user_id)"`
	CreatedAt           uint64        `orm:"column(created_at)"`
	PublishAt           uint64        `orm:"column(publish_at)"`
	ModifiedAt          uint64        `orm:"column(modified_at)"`
	SetTimer            uint64        `orm:"column(set_timer)"`
	Category            string        `orm:"column(category)"`
	Type                string        `orm:"column(type)"`
	IsDeleted           string        `orm:"column(is_deleted)"`
	Status              uint8         `orm:"column(status)"`
	CommentEnabled      uint8         `orm:"column(comment_enabled)"`
}

//PublishScoper 广播发布范围
type PublishScoper struct {
	GroupIDs []uint64 `json:"group_ids"`
	UserIDs  []uint64 `json:"user_ids"`
}

//TableName 表名
func (B *Bbs) TableName() string {
	return "bbs"
}

//GetOne 读取单条数据
func (B *Bbs) GetOne(ID uint64) (Bbs, error) {
	bbsInfo := Bbs{ID: ID}
	c := newCache(B.TableName(), "GetOne", ID)
	if c.getCache(&bbsInfo) == true {
		return bbsInfo, nil
	}
	if err := o.Read(&bbsInfo); err != nil {
		return Bbs{}, err
	}
	data, err := B.afterSelectHandle([]Bbs{bbsInfo})
	if err != nil {
		return Bbs{}, err
	}
	c.setCache(data[0])
	return data[0], nil
}

//Update 修改数据
func (B *Bbs) Update(ID uint64) error {
	num, err := o.Update(&Queue{ID: ID, Status: 3, ModifiedAt: uint64(time.Now().UnixNano() / 1e6)}, "Status", "ModifiedAt")
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
		item.PublishScope, err = B.publishScopeHandle(item.PublishScopeString)
		if err == nil {
			data[key] = item
		}
	}
	return data, err
}

//publishScopeHandle 处理发布范围
func (B *Bbs) publishScopeHandle(publishScopeString string) (PublishScoper, error) {
	var data PublishScoper
	var publishScope map[string][]string
	var err error
	err = json.Unmarshal([]byte(publishScopeString), &publishScope)
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
	}
	return data, err
}
