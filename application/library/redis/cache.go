package redis

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"bbs_server/application/function"
	"github.com/chasex/redis-go-cluster"
)

//C 缓存结构
type C struct {
	prefix string
	key    string
}

// Cache cache对象
var Cache cache.Cache

//NewCache 初始化缓存对象
func NewCache(prefix string, args ...interface{}) *C {
	c := new(C)
	c.prefix = prefix
	c.key = c.makeKey(args)
	return c
}

//makeKey 参数产生Key
func (c *C) makeKey(args ...interface{}) string {
	k, err := json.Marshal(args)
	if err != nil {
		logs.Error("GetCache make key error", err, args)
		return ""
	}
	return c.prefix + ":" + function.Md5(string(k), 32)
}

//Set 设置缓存
func (c *C) Set(data interface{}) bool {
	var (
		val []byte
		err error
	)
	if val, err = json.Marshal(data); err != nil {
		logs.Error("SetCache data Marshal error", err, data)
		return false
	}
	if err := Cache.Put(c.key, val, 48*time.Hour); err != nil {
		logs.Error("SetCache Put error", err, c.key, val)
		return false
	}
	return true
}

//Get 读取缓存
func (c *C) Get(data interface{}) bool {
	var (
		err error
		val string
	)

	if val, err = redis.String(Cache.Get(c.key), nil); err != nil {
		logs.Error("GetCache value Assertion error", err, c.key, val)
		return false
	}
	if err = json.Unmarshal([]byte(val), data); err != nil {
		logs.Error("GetCache data Unmarshal error", err, val, data)
		return false
	}
	return true
}

//Clear 清除缓存
func (c *C) Clear() {
	if err := Cache.Delete(c.prefix + "*"); err != nil {
		logs.Error("Clear cahce:", c.prefix, "*")
	}
}
