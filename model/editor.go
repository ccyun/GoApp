package model

//Editor 任务表结构
type Editor struct {
	ID       uint64 `orm:"column(id)"`
	SiteID   uint64 `orm:"column(site_id)"`
	BoardID  uint64 `orm:"column(board_id)"`
	EditorID uint64 `orm:"column(editor_id)"`
}

//TableName 表名
func (E *Editor) TableName() string {
	return "editor"
}

//GetEditor 读取单条数据
func (E *Editor) GetEditor(boardIDs []uint64) (map[uint64][]uint64, error) {
	var editorList []*Editor
	data := make(map[uint64][]uint64)
	if _, err := o.QueryTable(E).Filter("BoardID__in", boardIDs).Limit(-1).All(&editorList, "BoardID", "EditorID"); err != nil {
		return map[uint64][]uint64{}, err
	}
	for _, v := range editorList {
		data[v.BoardID] = append(data[v.BoardID], v.EditorID)
	}
	return data, nil
}
