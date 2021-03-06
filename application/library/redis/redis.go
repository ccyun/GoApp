package redis

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/chasex/redis-go-cluster"
)

//DefaultKey 默认前缀
var DefaultKey = ""

//Redis is Redis cache adapter.
type Redis struct {
	startNodes []string
	prefix     string
	p          *redis.Cluster // redis connection pool
}

//NewRedisCache create new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &Redis{prefix: DefaultKey}
}

//actually do the redis cmds
func (rc *Redis) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return rc.p.Do(commandName, args...)
}

//realKey 处理key
func (rc *Redis) realKey(key string) string {
	return rc.prefix + ":" + key
}

//lock 加锁
func (rc *Redis) lock(key string) bool {
	key += "_lock"
	value := time.Now().UnixNano()
	n, err := rc.do("SETNX", key, value)
	if res, ok := n.(int64); ok && res == 1 {
		if _, err := rc.do("EXPIRE", key, 300); err == nil {
			return true
		}
	} else {
		logs.Error("redis setnx error:", n, key, err)
	}
	return false
}

//unlock 解锁
func (rc *Redis) unlock(key string) bool {
	key += "_lock"
	n, err := rc.do("DEL", key)
	if res, ok := n.(int64); ok && res == 1 {
		return true
	}
	logs.Error("redis del error:", n, key, err)
	return false
}

//isLock 检查是否是锁定状态
func (rc *Redis) isLock(key string) bool {
	var (
		val    string
		keyArr []string
	)
	keyArr = strings.Split(key, ":")
	for _, k := range keyArr {
		if val == "" {
			val = k
		} else {
			val += ":" + k
		}
		if val != key {
			n, err := rc.do("EXISTS", val+"_lock")
			if v, ok := n.(int64); (ok && v == 1) || err != nil {
				return true
			}
		}
	}
	return false
}

//updateIndex 缓存key集合
func (rc *Redis) updateIndex(key string) error {
	var (
		val      string
		indexVal string
		keyArr   []string
		keys     []string
	)
	keyArr = strings.Split(key, ":")
	lev := len(keyArr) - 1
	for i := lev; i > 0; i-- {
		val = strings.Join(keyArr[:i], ":")
		if _, err := rc.do("HSET", val, key, "cache"); err != nil {
			return err
		}
		keys = append(keys, val)
		if i > 1 {
			indexVal = strings.Join(keyArr[:(i-1)], ":")
			for _, v := range keys {
				if _, err := rc.do("HSET", indexVal, v, "index"); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

//clearIndex 清理索引
func (rc *Redis) clearIndex(key string) error {
	var (
		val    string
		keyArr []string
	)
	keyArr = strings.Split(key, ":")
	for _, k := range keyArr {
		if val == "" {
			val = k
		} else {
			val += ":" + k
		}
		if val != key {
			if _, err := rc.do("HDEL", val, key); err != nil {
				return err
			}
		}
	}
	return nil
}

//clearAll 删除数据
func (rc *Redis) clearAll(key string) error {
	//加锁
	if rc.lock(key) == false {
		return errors.New("clearAll lock error")
	}
	cachedKeys, err := redis.Strings(rc.do("HKEYS", key))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if err := rc.clearIndex(str); err != nil {
			return err
		}
		if _, err = rc.do("DEL", str); err != nil {
			return err
		}
	}
	_, err = rc.do("DEL", key)
	//解锁
	if rc.unlock(key) == false {
		return errors.New("clearAll unlock error")
	}
	return err
}

//Get cache from redis.
func (rc *Redis) Get(key string) interface{} {
	key = rc.realKey(key)
	if rc.isLock(key) == true {
		return nil
	}
	if v, err := rc.do("GET", key); err == nil {
		return v
	}
	return nil
}

//GetMulti get cache from redis.
func (rc *Redis) GetMulti(keys []string) []interface{} {
	var rv []interface{}
	for _, key := range keys {
		key = rc.realKey(key)
		if v, err := rc.do("GET", key); err == nil {
			rv = append(rv, v.([]byte))
		} else {
			rv = append(rv, err)
		}
	}
	return rv
}

//Put put cache to redis.
func (rc *Redis) Put(key string, val interface{}, timeout time.Duration) error {
	key = rc.realKey(key)
	if rc.isLock(key) == true {
		return nil
	}
	if err := rc.updateIndex(key); err != nil {
		return err
	}
	_, err := rc.do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

//Delete delete cache in redis.
func (rc *Redis) Delete(key string) error {
	key = rc.realKey(key)
	logs.Info("Redis delete key:", key)
	if strings.HasSuffix(key, "*") == true {
		key = strings.TrimRight(key, "*")
		return rc.clearAll(key)
	}
	if _, err := rc.do("DEL", key); err != nil {
		return err
	}
	return rc.clearIndex(key)
}

//IsExist check cache's existence in redis.
func (rc *Redis) IsExist(key string) bool {
	key = rc.realKey(key)
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil || v == false {
		go rc.clearIndex(key)
		return false
	}
	return v
}

//Incr increase counter in redis.
func (rc *Redis) Incr(key string) error {
	key = rc.realKey(key)
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

//Decr decrease counter in redis.
func (rc *Redis) Decr(key string) error {
	key = rc.realKey(key)
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

//ClearAll clean all cache in redis. delete this redis collection.
func (rc *Redis) ClearAll() error {
	return rc.clearAll(rc.prefix)
}

//StartAndGC start redis cache adapter.
//config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
//the cache item in redis are stored forever,
//so no gc operation.
func (rc *Redis) StartAndGC(config string) error {
	var conf struct {
		Nodes  []string `json:"nodes"`
		Prefix string   `json:"prefix"`
	}
	json.Unmarshal([]byte(config), &conf)
	if len(conf.Nodes) == 0 {
		return errors.New("config has no conn key")
	}
	rc.startNodes = conf.Nodes
	rc.prefix = conf.Prefix
	rc.connectInit()
	return nil
}

//connectInit connect to redis.
func (rc *Redis) connectInit() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   rc.startNodes,
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    20,
			AliveTime:    50 * time.Second,
		})
	if err != nil {
		log.Fatalf("redis.New error: %s", err.Error())
	}
	rc.p = cluster
}

func init() {
	cache.Register("redis", NewRedisCache)
}
