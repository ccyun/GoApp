package model

import (
	"fmt"

	"github.com/ccyun/GoApp/application/library/redis"
)

//Board 任务表结构
type Board struct {
	base
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
	boardInfo := Board{ID: ID}
	//c := newCache(B.TableName(), "GetOne", ID)
	c := redis.NewCache(fmt.Sprintf("D%d%s", B.siteID, B.TableName()), "GetOne", ID)
	if c.Get(&boardInfo) == true {
		return boardInfo, nil
	}
	if err := o.Read(&boardInfo); err != nil {
		return Board{}, err
	}
	data, err := B.afterSelectHandle([]Board{boardInfo})
	if err != nil {
		return Board{}, err
	}
	c.Set(data[0])
	return data[0], nil
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
