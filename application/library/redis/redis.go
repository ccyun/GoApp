package redis

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/chasex/redis-go-cluster"

	"strings"

	"github.com/astaxie/beego/cache"
)

//DefaultKey 默认前缀
var DefaultKey = ""

//Cache is Redis cache adapter.
type Cache struct {
	startNodes []string
	prefix     string
	p          *redis.Cluster // redis connection pool
}

//NewRedisCache create new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &Cache{prefix: DefaultKey}
}

//actually do the redis cmds
func (rc *Cache) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return rc.p.Do(commandName, args...)
}

//realKey 处理key
func (rc *Cache) realKey(key string) string {
	return rc.prefix + ":" + key
}

//updateIndex 缓存key集合
func (rc *Cache) updateIndex(key string) error {
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
func (rc *Cache) clearIndex(key string) error {
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
func (rc *Cache) clearAll(key string) error {
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
	return err
}

//Get cache from redis.
func (rc *Cache) Get(key string) interface{} {
	key = rc.realKey(key)
	//todo 只是模糊查询
	// if strings.HasSuffix(key, "*") == true {
	// 	//key = strings.TrimRight(key, "*")
	// 	//return rc.clearAll(key)
	// }
	if v, err := rc.do("GET", key); err == nil {
		return v
	}
	return nil
}

//GetMulti get cache from redis.
func (rc *Cache) GetMulti(keys []string) []interface{} {
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
func (rc *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	key = rc.realKey(key)
	if err := rc.updateIndex(key); err != nil {
		return err
	}
	_, err := rc.do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

//Delete delete cache in redis.
func (rc *Cache) Delete(key string) error {
	key = rc.realKey(key)
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
func (rc *Cache) IsExist(key string) bool {
	key = rc.realKey(key)
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil || v == false {
		go rc.clearIndex(key)
		return false
	}
	return v
}

//Incr increase counter in redis.
func (rc *Cache) Incr(key string) error {
	key = rc.realKey(key)
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

//Decr decrease counter in redis.
func (rc *Cache) Decr(key string) error {
	key = rc.realKey(key)
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

//ClearAll clean all cache in redis. delete this redis collection.
func (rc *Cache) ClearAll() error {
	return rc.clearAll(rc.prefix)
}

//StartAndGC start redis cache adapter.
//config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
//the cache item in redis are stored forever,
//so no gc operation.
func (rc *Cache) StartAndGC(config string) error {
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
func (rc *Cache) connectInit() {
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
