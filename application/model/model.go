package model

import (

	//redis 驱动

	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/library/redis"
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
	siteID uint64 `orm:"-",json:"-"`
}

//RegisterModels 注册Model
func RegisterModels() {
	orm.Debug = Debug
	orm.RegisterModelWithPrefix(DBPrefix, new(Queue))
	orm.RegisterModelWithPrefix(DBPrefix, new(Board))
	orm.RegisterModelWithPrefix(DBPrefix, new(Bbs))
	orm.RegisterModelWithPrefix(DBPrefix, new(Editor))
	orm.RegisterModelWithPrefix(DBPrefix, new(PublishScope))
	orm.RegisterModelWithPrefix(DBPrefix, new(Feed))
	orm.RegisterModelWithPrefix(DBPrefix, new(BbsTask))

	o = orm.NewOrm()
	DB, _ = orm.NewQueryBuilder(DBType)
}

//AfterUpdate 错误处理,处理 增/删/改；以及更新缓存
func (b *base) AfterUpdate(tableName string, num int64, err error) bool {
	if num == 0 {
		logs.Notice("Model info:", orm.ErrNoRows)
		return false
	}
	if err != nil {
		logs.Error("Model error:", err)
		return false
	}
	//异步 clearCache
	go redis.NewCache(fmt.Sprintf("D%d%s", b.siteID, tableName)).Clear()
	return true
}
