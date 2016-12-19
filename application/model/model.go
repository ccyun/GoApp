package model

import (
	"encoding/json"
	"strconv"

	//redis 驱动

	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/function"
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
	funcName  string
}

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
	if option["userID"] != "" {
		userID, _ := strconv.Atoi(option["userID"])
		UserID = uint64(userID)
	}
	if option["siteID"] != "" {
		siteID, _ := strconv.Atoi(option["siteID"])
		UserID = uint64(siteID)
	}
}

//L 语言log
func L(log string) string {
	return RequestID + "  " + log
}

//NewCache 初始化缓存对象
func NewCache(tableName string, funcName string) *C {
	c := new(C)
	c.tableName = tableName
	c.funcName = funcName
	return c
}

//makeKey 参数产生Key
func (c *C) makeKey(arg interface{}) (string, error) {
	k, err := json.Marshal(arg)
	if err != nil {
		logs.Error(L("GetCache make key error"), err)
		return "", err
	}
	key := "db:" + strconv.FormatUint(SiteID, 10) + ":"
	if c.tableName != "" {
		key += c.tableName + ":"
	}
	if c.funcName != "" {
		key += c.funcName + ":"
	}
	key += function.Md5(string(k))
	return key, nil
}

//SetCache 设置缓存
func (c *C) SetCache(arg interface{}, data interface{}) bool {
	key, err := c.makeKey(arg)
	if err != nil {
		return false
	}
	value, err := json.Marshal(data)
	if err != nil {
		logs.Error(L("SetCache data Marshal error"), err)
		return false
	}
	err2 := Cache.Put(key, value, 240*time.Hour)
	if err2 != nil {
		logs.Error(L("SetCache Put error"), err2)
		return false
	}
	return true
}

//GetCache 读取缓存
func (c *C) GetCache(arg interface{}, data interface{}) bool {
	key, err := c.makeKey(arg)
	if err != nil {
		return false
	}

	val := Cache.Get(key)

	switch val.(type) {
	case []byte:
		value := string(val.([]byte))

		err := json.Unmarshal([]byte(value), data)
		if err != nil {
			logs.Error(L("GetCache data Unmarshal error"), err)
			return false
		}
		return true
	}
	logs.Error(L("GetCache value Assertion error"))
	return false
}

//ClearCache 清除缓存
func (c *C) ClearCache(keys ...string) bool {
	key := "db:" + strconv.FormatUint(SiteID, 10) + ":"
	if len(keys) == 0 {
		keys[0] = c.tableName
	}
	for _, v := range keys {
		if err := Cache.Delete(key + v + "*"); err != nil {
			logs.Error(L("ClearCache error"), err)
			return false
		}
	}
	return true
}
