package model

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
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
	// Cache cache对象
	Cache cache.Cache
	//RequestID 请求ID
	RequestID string
)

//RegisterModels 注册Model
func RegisterModels() {
	orm.Debug = Debug
	orm.RegisterModelWithPrefix(DBPrefix, new(Queue))
	orm.RegisterModelWithPrefix(DBPrefix, new(Board))
	orm.RegisterModelWithPrefix(DBPrefix, new(Bbs))
	orm.RegisterModelWithPrefix(DBPrefix, new(Editor))
	orm.RegisterModelWithPrefix(DBPrefix, new(PublishScope))
	o = orm.NewOrm()
	DB, _ = orm.NewQueryBuilder(DBType)
}

//Init 初始化
func Init(option map[string]string) {
	if option["requestID"] != "" {
		RequestID = option["requestID"]
	}
}

//L 语言log
func L(log string) string {
	return RequestID + "  " + log
}
