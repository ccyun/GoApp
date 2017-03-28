package model

import (

	//redis 驱动

	"fmt"

	"github.com/astaxie/beego/orm"
	"bbs_server/application/library/redis"
)

var (
	//Debug 调试模式
	Debug bool
	//DBType 数据库类型
	DBType string
	//DBPrefix 表前缀
	DBPrefix string
	//DB 高级查询
	DB orm.QueryBuilder
	//o orm对象
	o orm.Ormer
)

type base struct {
}

//RegisterModels 注册Model
func RegisterModels() {
	orm.Debug = Debug
	orm.RegisterModelWithPrefix(
		DBPrefix,
		new(Queue),
		new(Board),
		new(Bbs),
		new(Editor),
		new(PublishScope),
		new(Feed),
		new(BbsTask),
		new(Todo),
		new(BbsTaskReply),
	)
	o = orm.NewOrm()
	DB, _ = orm.NewQueryBuilder(DBType)
}

//AfterUpdate 错误处理,处理 增/删/改；以及更新缓存
func AfterUpdate(tableName string, siteID uint64) bool {
	//异步 clearCache
	go redis.NewCache(fmt.Sprintf("D%d%s", siteID, tableName)).Clear()
	return true
}

//Begin 开启事务
func Begin() error {
	return o.Begin()
}

//Rollback 事务回滚
func Rollback() error {
	return o.Rollback()
}

//Commit 提交事务
func Commit() error {
	return o.Commit()
}
