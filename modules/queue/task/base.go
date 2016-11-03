package task

import "github.com/ccyun/GoApp/model"
import "github.com/astaxie/beego/orm"

//base 任务处理适配器（基类）
type base struct {
	//o orm对象
	o orm.Ormer
}

//Begin 开启事务
func (B *base) Begin() error {
	B.o = orm.NewOrm()
	return B.o.Begin()
}

//Commit 提交事务
func (B *base) Rollback() error {
	return B.o.Rollback()
}

//Commit 提交事务
func (B *base) Commit() error {
	return B.o.Commit()
}

//NewTask 新任务对象
func (B *base) NewTask(taskInfo model.Queue) error {

	return nil
}

//CreateFeed 创建Feed
func (B *base) CreateFeed() error {
	return nil
}

//CreateRelation 创建接收者关系
func (B *base) CreateRelation() error {
	return nil
}

//GetPublishScopeUsers 分析发布范围
func (B *base) GetPublishScopeUsers() error {
	return nil
}

//CreateUnread 创建未读计数
func (B *base) CreateUnread() error {
	return nil
}

//SendMsg 发送消息
func (B *base) SendMsg() error {
	return nil
}
