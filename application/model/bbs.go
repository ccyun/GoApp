package model

import (
	"encoding/json"
	"time"

	"sort"

	"github.com/astaxie/beego/orm"
)

//Bbs 任务表结构
type Bbs struct {
	base
	ID                 uint64              `orm:"column(id)"`
	SiteID             uint64              `orm:"column(site_id)"`
	BoardID            uint64              `orm:"column(board_id)"`
	DiscussID          uint64              `orm:"column(discuss_id)"`
	Title              string              `orm:"column(title)"`
	Description        string              `orm:"column(description)"`
	Content            string              `orm:"column(content)"`
	IsBrowser          uint8               `orm:"column(is_browser)"`
	IsAuth             uint8               `orm:"column(is_auth)"`
	Link               string              `orm:"column(link)"`
	PublishScopeString string              `orm:"column(publish_scope)"`
	PublishScope       PublishScoper       `orm:"-"`
	MsgCount           uint64              `orm:"column(msg_count)"`
	AttachmentsString  string              `orm:"column(attachments)"`
	Attachments        []map[string]string `orm:"-"`
	Thumb              string              `orm:"-"`
	UserID             uint64              `orm:"column(user_id)"`
	CreatedAt          uint64              `orm:"column(created_at)"`
	PublishAt          uint64              `orm:"column(publish_at)"`
	ModifiedAt         uint64              `orm:"column(modified_at)"`
	SetTimer           uint64              `orm:"column(set_timer)"`
	Category           string              `orm:"column(category)"`
	Type               string              `orm:"column(type)"`
	IsDeleted          string              `orm:"column(is_deleted)"`
	Status             uint8               `orm:"column(status)"`
	CommentEnabled     uint8               `orm:"column(comment_enabled)"`
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
	//c.Set(data[0])
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
		//item.PublishScope, err = B.publishScopeHandle(item.PublishScopeString)
		if err := json.Unmarshal([]byte(item.PublishScopeString), &item.PublishScope); err != nil {
			return nil, err
		}
		sort.Sort(Uint64Slice(item.PublishScope.GroupIDs))
		sort.Sort(Uint64Slice(item.PublishScope.UserIDs))
		if item.AttachmentsString != "" {
			err = json.Unmarshal([]byte(item.AttachmentsString), &item.Attachments)
		}
		if len(item.Attachments) > 0 {
			if thumb, ok := item.Attachments[0]["url"]; ok {
				item.Thumb = thumb
			}
		}
		if err == nil {
			data[key] = item
		}
	}
	return data, err
}
