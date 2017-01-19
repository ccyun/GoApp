package model

//BbsTask 任务表结构
type BbsTask struct {
	ID           uint64 `orm:"column(id)"`
	SiteID       uint64 `orm:"column(site_id)"`
	BoardID      uint64 `orm:"column(board_id)"`
	BbsID        uint64 `orm:"column(bbs_id)"`
	Restriction  string `orm:"column(restriction)"`
	EndTime      uint64 `json:"end_time"`
	AllowExpired uint8  `json:"allow_expired"`
	IsCycle      uint8  `json:"is_cycle"`
	CycleRule    string `json:"cycle_rule"`
	IsClose      uint8  `json:"is_close"`
	CloseAt      uint64 `json:"close_at"`
}

//TableName 表名
func (B *BbsTask) TableName() string {
	return "bbs_task"
}

//GetOne 读取单条数据
func (B *BbsTask) GetOne(ID uint64) (BbsTask, error) {
	bbsTaskInfo := BbsTask{BbsID: ID}
	c := newCache(B.TableName(), "GetOne", ID)
	if c.getCache(&bbsTaskInfo) == true {
		return bbsTaskInfo, nil
	}
	if err := o.Read(&bbsTaskInfo); err != nil {
		return BbsTask{}, err
	}
	c.setCache(bbsTaskInfo)
	return bbsTaskInfo, nil
}
