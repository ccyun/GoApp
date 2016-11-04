package model

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

//Bbs 任务表结构
type Bbs struct {
	ID                 uint64 `orm:"column(id)"`
	SiteID             uint64 `orm:"column(site_id)"`
	BoardID            uint64 `orm:"column(board_id)"`
	Title              string `orm:"column(title)"`
	Description        string `orm:"column(description)"`
	Content            string `orm:"column(content)"`
	PublishScopeString string `orm:"column(publish_scope)"`
	MsgCount           uint8  `orm:"column(msg_count)"`
	Attachments        string `orm:"column(attachments)"`
	UsesID             uint64 `orm:"column(user_id)"`
	CreatedAt          uint64 `orm:"column(created_at)"`
	PublishAt          uint64 `orm:"column(publish_at)"`
	ModifiedAt         uint64 `orm:"column(modified_at)"`
	SetTimer           uint64 `orm:"column(set_timer)"`
	Category           string `orm:"column(category)"`
	Type               string `orm:"column(type)"`
	IsDeleted          string `orm:"column(is_deleted)"`
	Status             uint8  `orm:"column(status)"`
	CommentEnabled     uint8  `orm:"column(comment_enabled)"`
}

//BbsInfo 公告
type BbsInfo struct {
	Bbs
	PublishScope map[string][]uint64
}

//TableName 表名
func (B *Bbs) TableName() string {
	return "bbs"
}

//GetOne 读取单条数据
func (B *Bbs) GetOne(ID uint64) (BbsInfo, error) {
	bbsInfo := Bbs{}
	if err := o.QueryTable(B).Filter("ID", ID).One(&bbsInfo); err != nil {
		return BbsInfo{}, err
	}
	return B.afterSelectHandle([]Bbs{bbsInfo})
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
func (B *Bbs) afterSelectHandle(data []Bbs) (BbsInfo, error) {
	log.Println(data)

	for key, item := range data {
		log.Println(key)
		log.Println(item)
	}
	return BbsInfo{}, nil
}
