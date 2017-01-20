package model

//PublishScope 任务表结构
type PublishScope struct {
	base
	ID               uint64 `orm:"column(id)"`
	SiteID           uint64 `orm:"column(site_id)"`
	BoardID          uint64 `orm:"column(board_id)"`
	PublishScopeType string `orm:"column(publish_scope_type)"`
	PublishScope     uint64 `orm:"column(publish_scope)"`
}

//TableName 表名
func (P *PublishScope) TableName() string {
	return "publishscope"
}

//GetPublishScope 读取单条数据
func (P *PublishScope) GetPublishScope(boardIDs []uint64) (map[uint64][]uint64, error) {
	var PublishScopeList []*PublishScope
	data := make(map[uint64][]uint64)
	if _, err := o.QueryTable(P).Filter("BoardID__in", boardIDs).Limit(-1).All(&PublishScopeList, "BoardID", "PublishScope"); err != nil {
		return map[uint64][]uint64{}, err
	}
	for _, v := range PublishScopeList {
		data[v.BoardID] = append(data[v.BoardID], v.PublishScope)
	}
	return data, nil
}
