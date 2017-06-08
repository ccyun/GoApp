package model

//BbsTask 任务表结构
type BbsTask struct {
	base
	ID                 uint64 `orm:"column(id)"`
	SiteID             uint64 `orm:"column(site_id)"`
	BoardID            uint64 `orm:"column(board_id)"`
	BbsID              uint64 `orm:"column(bbs_id)"`
	ReplyType          uint8  `orm:"column(reply_type)"`
	IsResubmit         uint8  `orm:"column(is_resubmit)"`
	AuditType          uint8  `orm:"column(audit_type)"`
	IsFlow             uint8  `orm:"column(is_flow)"`
	AuditUserIDsString string `orm:"column(audit_user_ids)"`
	EndTime            uint64 `orm:"column(end_time)"`
	AllowExpired       uint8  `orm:"column(allow_expired)"`
	AuditEndTime       uint64 `orm:"column(audit_end_time)"`
	ReplyRemindAt      uint64 `orm:"column(reply_remind_at)"`
	AuditRemindAt      uint64 `orm:"column(audit_remind_at)"`
	IsRedo             uint8  `orm:"column(is_redo)"`
	IsStopReply        uint8  `orm:"column(is_stop_reply)"`
	IsDone             uint8  `orm:"column(is_done)"`
	IsClose            uint8  `orm:"column(is_close)"`
	CloseAt            uint64 `orm:"column(close_at)"`
}

//TableName 表名
func (B *BbsTask) TableName() string {
	return "bbs_task"
}

//GetOne 读取单条数据
func (B *BbsTask) GetOne(BbsID uint64) (BbsTask, error) {
	var bbsTaskInfo BbsTask
	// c := redis.NewCache(fmt.Sprintf("D%d%s", B.siteID, B.TableName()), "GetOne", ID)
	// if c.Get(&bbsTaskInfo) == true {
	// 	return bbsTaskInfo, nil
	// }
	if err := o.QueryTable(B).Filter("BbsID", BbsID).One(&bbsTaskInfo); err != nil {
		return bbsTaskInfo, err
	}
	//c.Set(bbsTaskInfo)
	return bbsTaskInfo, nil
}
