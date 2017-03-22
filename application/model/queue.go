package model

import (
	"log"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
)

//Queue 任务表结构
type Queue struct {
	base
	ID           uint64 `orm:"column(id)"`
	SiteID       uint64 `orm:"column(site_id)"`
	CustomerCode string `orm:"column(customer_code)"`
	TaskType     string `orm:"column(task_type)"`
	Action       string `orm:"column(action)"`
	Status       uint8  `orm:"column(status)"`
	TryCount     uint8  `orm:"column(try_count)"`
	SetTimer     uint64 `orm:"column(set_timer)"`
	ModifiedAt   uint64 `orm:"column(modified_at)"`
}

//TableName 表名
func (Q *Queue) TableName() string {
	return "task"
}

//Pull 读取单条数据
func (Q *Queue) Pull() (Queue, error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	taskInfo := Queue{}
	cond := orm.NewCondition()
	condition := cond.And("SetTimer__lt", uint64(time.Now().UnixNano()/1e6)).AndCond(cond.And("Status", 0).OrCond(cond.And("Status", 3).And("TryCount__lt", 3)))
	if err := o.QueryTable(Q).SetCond(condition).OrderBy("Status", "ID").One(&taskInfo); err != nil {
		return Queue{}, err
	}
	if err := Q.lockTask(taskInfo); err != nil {
		return Queue{}, err
	}
	log.Println(taskInfo.ID)
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

//TimeOut 处理超时任务
func (Q *Queue) TimeOut() error {
	nowTime := uint64(time.Now().UnixNano() / 1e6)
	data := orm.Params{
		"Status":     3,
		"ModifiedAt": nowTime,
	}
	if _, err := o.QueryTable(Q).Filter("Status", 1).Filter("ModifiedAt__lt", (nowTime - 7200000)).Update(data); err != nil {
		return err
	}
	return nil
}

//Fail 修改数据
func (Q *Queue) Fail(ID uint64) bool {
	if _, err := o.Update(&Queue{ID: ID, Status: 3, ModifiedAt: uint64(time.Now().UnixNano() / 1e6)}, "Status", "ModifiedAt"); err != nil {
		return false
	}
	return true

}

//Delete 删除数据
func (Q *Queue) Delete(ID uint64) bool {
	if _, err := o.Delete(&Queue{ID: ID}); err != nil {
		return false
	}
	return true
}
