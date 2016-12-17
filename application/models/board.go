package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Board 任务表结构
type Board struct {
	ID           uint64   `orm:"column(id)"`
	SiteID       uint64   `orm:"column(site_id)"`
	BoardName    string   `orm:"column(board_name)"`
	DiscussID    uint64   `orm:"column(discuss_id)"`
	Avatar       string   `orm:"column(avatar)"`
	CreatedAt    uint64   `orm:"column(created_at)"`
	ModifiedAt   uint64   `orm:"column(modified_at)"`
	UsesID       uint64   `orm:"column(user_id)"`
	Status       uint8    `orm:"column(status)"`
	IsDeleted    uint8    `orm:"column(is_deleted)"`
	EditorIDs    []uint64 `orm:"-"`
	PublishScope []uint64 `orm:"-"`
}

//TableName 表名
func (B *Board) TableName() string {
	return "board"
}

//GetOne 读取单条数据
func (B *Board) GetOne(ID uint64) (Board, error) {
	boardInfo := Board{}
	if err := o.QueryTable(B).Filter("ID", ID).One(&boardInfo); err != nil {
		return Board{}, err
	}
	data, err := B.afterSelectHandle([]Board{boardInfo})
	if err != nil {
		return Board{}, err
	}
	return data[0], nil
}

//Update 修改数据
func (B *Board) Update(ID uint64) error {
	num, err := o.Update(&Queue{ID: ID, Status: 3, ModifiedAt: uint64(time.Now().UnixNano() / 1e6)}, "Status", "ModifiedAt")
	if num == 0 {
		err = orm.ErrNoRows
	}
	return err
}

//afterSelectHandle 查询结果处理
func (B *Board) afterSelectHandle(data []Board) ([]Board, error) {
	var err error
	boardIDs := []uint64{}
	for _, item := range data {
		boardIDs = append(boardIDs, item.ID)
	}
	editor := new(Editor)
	editorIDs, _ := editor.GetEditor(boardIDs)
	publishscope := new(PublishScope)
	publishScopeIDs, _ := publishscope.GetPublishScope(boardIDs)
	for k, item := range data {
		if editorIDs[item.ID] != nil {
			item.EditorIDs = editorIDs[item.ID]
		}
		if publishScopeIDs[item.ID] != nil {
			item.PublishScope = publishScopeIDs[item.ID]
			data[k] = item
		}
	}
	return data, err
}
