package model

//BbsTaskAudit 任务审核表
type BbsTaskAudit struct {
	base
	ID           uint64 `orm:"column(id)"`
	SiteID       uint64 `orm:"column(site_id)"`
	BoardID      uint64 `orm:"column(board_id)"`
	BbsID        uint64 `orm:"column(bbs_id)"`
	ReplyID      uint64 `orm:"column(reply_id)"`
	SubTaskID    uint64 `orm:"column(sub_task_id)"`
	SubReplyID   uint64 `orm:"column(sub_reply_id)"`
	AuditUserID  uint64 `orm:"column(audit_user_id)"`
	AuditAt      uint64 `orm:"column(audit_at)"`
	AuditScore   int    `orm:"column(audit_score)"`
	AuditOpinion string `orm:"column(audit_opinion)"`
	Data         string `orm:"column(data)"`
	Status       int8   `orm:"column(status)"`
	CreatedAt    uint64 `orm:"column(created_at)"`
}

//TableName 表名
func (B *BbsTaskAudit) TableName() string {
	return "bbs_task_audit"
}

//GetUNAuditUserIDs 读取已反馈用户列表
func (B *BbsTaskAudit) GetUNAuditUserIDs(BbsID uint64) ([]uint64, error) {
	var (
		auditList []BbsTaskAudit
		data      []uint64
	)
	if _, err := o.QueryTable(B).Filter("BbsID", BbsID).Filter("Status", 0).Limit(-1).All(&auditList, "UserID"); err != nil {
		return nil, err
	}
	for _, v := range auditList {
		data = append(data, v.AuditUserID)
	}
	return data, nil
}
