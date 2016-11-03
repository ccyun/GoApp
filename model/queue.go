package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Queue 任务表结构
type Queue struct {
	ID           uint64 `orm:"column(id)"`
	SiteID       uint64 `orm:"column(site_id)"`
	CustomerCode string `orm:"column(customer_code)"`
	TaskType     string `orm:"column(task_type)"`
	Action       string `orm:"column(action)"`
	Status       uint8  `orm:"column(status)"`
	TryCount     uint8  `orm:"column(try_count)"`
	SetTimer     uint64 `orm:"column(set_timer)"`
	ModifiedAt   uint64 `orm:"column(modified_at),index"`
}

//TableName 表名
func (Q *Queue) TableName() string {
	return "task"
}

//GetOneTask 读取单条数据
func (Q *Queue) GetOneTask() (Queue, error) {
	taskInfo := Queue{}
	cond := orm.NewCondition()
	condition := cond.And("SetTimer__lt", uint64(time.Now().UnixNano()/1e6)).AndCond(cond.And("Status", 0).OrCond(cond.And("Status", 3).And("TryCount__lt", 3)))
	if err := o.QueryTable(Q).SetCond(condition).OrderBy("Status", "ID").One(&taskInfo); err != nil {
		return Queue{}, err
	}
	if err := Q.lockTask(taskInfo); err != nil {
		return Queue{}, err
	}
	return taskInfo, nil
}

//lockTask 上锁，处理中
func (Q *Queue) lockTask(taskInfo Queue) error {
	data := orm.Params{
		"Status":     1,
		"ModifiedAt": uint64(time.Now().UnixNano() / 1e6),
		"TryCount":   orm.ColValue(orm.ColAdd, 1),
	}
	num, err := o.QueryTable(Q).Filter("ID", taskInfo.ID).Filter("Status", taskInfo.Status).Filter("ModifiedAt", taskInfo.ModifiedAt).Update(data)
	if num == 0 {
		err = orm.ErrNoRows
	}
	return err
}

//Update 修改数据
func (Q *Queue) Update(ID uint64) error {
	num, err := o.Update(&Queue{ID: ID, Status: 3, ModifiedAt: uint64(time.Now().UnixNano() / 1e6)}, "Status", "ModifiedAt")
	if num == 0 {
		err = orm.ErrNoRows
	}
	return err
}

//Delete 删除数据
func (Q *Queue) Delete(ID uint64) error {
	num, err := o.Delete(&Queue{ID: ID})
	if num == 0 {
		err = orm.ErrNoRows
	}
	return err
}
