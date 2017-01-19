package model

import (
	"encoding/json"
	"strconv"

	//redis 驱动

	"time"

	"fmt"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/function"
	"github.com/chasex/redis-go-cluster"
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
	//UserID 用户ID
	UserID uint64
	//SiteID 站点ID
	SiteID uint64
)

//C model cache
type C struct {
	tableName string
	key       string
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

//Init 初始化
func Init(option map[string]string) {
	if option["requestID"] != "" {
		RequestID = option["requestID"]
	}
	if option["userID"] != "" {
		userID, _ := strconv.Atoi(option["userID"])
		UserID = uint64(userID)
	}
	if option["siteID"] != "" {
		siteID, _ := strconv.Atoi(option["siteID"])
		UserID = uint64(siteID)
	}
}

//AfterUpdate 错误处理,处理 增/删/改；以及更新缓存
func AfterUpdate(tableName string, num int64, err error) bool {
	if num == 0 {
		logs.Notice(L("Model info:"), orm.ErrNoRows)
		return false
	}
	if err != nil {
		logs.Error(L("Model error:"), err)
		return false
	}
	//异步 clearCache
	go newCache(tableName).clearCache(tableName)
	return true
}

//L 语言log
func L(log string) string {
	return RequestID + "  " + log
}

///////////////////////////////Cache//////////////////////////////////////////////////////////////////////////////////////////////////////

//newCache 初始化缓存对象
func newCache(tableName string, args ...interface{}) *C {
	c := new(C)
	c.tableName = tableName
	c.key = c.makeKey(args)
	return c
}

//makeKey 参数产生Key
func (c *C) makeKey(args ...interface{}) string {
	k, err := json.Marshal(args)
	if err != nil {
		logs.Error(L("GetCache make key error"), err)
		return ""
	}
	return fmt.Sprintf("D%d%s:%s", SiteID, c.tableName, function.Md5(string(k), 32))
}

//setCache 设置缓存
func (c *C) setCache(data interface{}) bool {
	var (
		val []byte
		err error
	)
	if val, err = json.Marshal(data); err != nil {
		logs.Error(L("SetCache data Marshal error"), err)
		return false
	}
	if err := Cache.Put(c.key, val, 48*time.Hour); err != nil {
		logs.Error(L("SetCache Put error"), err)
		return false
	}
	return true
}

//getCache 读取缓存
func (c *C) getCache(data interface{}) bool {
	var (
		err error
		val string
	)
	if val, err = redis.String(Cache.Get(c.key), nil); err != nil {
		logs.Info(L("GetCache value Assertion error"), err)
		return false
	}
	if err = json.Unmarshal([]byte(val), data); err != nil {
		logs.Error(L("GetCache data Unmarshal error"), err)
		return false
	}
	return true
}

//clearCache 清除缓存
func (c *C) clearCache(keys ...string) bool {
	key := "D" + strconv.FormatUint(SiteID, 10)
	for _, v := range keys {
		if err := Cache.Delete(key + v + "*"); err != nil {
			logs.Error(L("ClearCache error"), err)
			return false
		}
	}
	return true
}
