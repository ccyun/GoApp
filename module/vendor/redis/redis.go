package redis

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"

	"github.com/astaxie/beego/cache"
)

var (
	// DefaultKey the collection name of redis for cache adapter.
	DefaultKey = "beecacheRedis"
)

func init() {
	cache.Register("redis", NewRedisCache)
}

// Cache is Redis cache adapter.
type Cache struct {
	p        []*redis.Pool // redis connection pool
	conninfo []string
	dbNum    int
	key      string
	password string
}

// NewRedisCache create new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &Cache{key: DefaultKey}
}

// actually do the redis cmds
func (rc *Cache) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// Get cache from redis.
func (rc *Cache) Get(key string) interface{} {
	if v, err := rc.do("GET", key); err == nil {
		return v
	}
	return nil
}

// GetMulti get cache from redis.
func (rc *Cache) GetMulti(keys []string) []interface{} {
	size := len(keys)
	var rv []interface{}
	c := rc.p.Get()
	defer c.Close()
	var err error
	for _, key := range keys {
		err = c.Send("GET", key)
		if err != nil {
			goto ERROR
		}
	}
	if err = c.Flush(); err != nil {
		goto ERROR
	}
	for i := 0; i < size; i++ {
		if v, err := c.Receive(); err == nil {
			rv = append(rv, v.([]byte))
		} else {
			rv = append(rv, err)
		}
	}
	return rv
ERROR:
	rv = rv[0:0]
	for i := 0; i < size; i++ {
		rv = append(rv, nil)
	}

	return rv
}

// Put put cache to redis.
func (rc *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	var err error
	if _, err = rc.do("SETEX", key, int64(timeout/time.Second), val); err != nil {
		return err
	}

	if _, err = rc.do("HSET", rc.key, key, true); err != nil {
		return err
	}
	return err
}

// Delete delete cache in redis.
func (rc *Cache) Delete(key string) error {
	var err error
	if _, err = rc.do("DEL", key); err != nil {
		return err
	}
	_, err = rc.do("HDEL", rc.key, key)
	return err
}

// IsExist check cache's existence in redis.
func (rc *Cache) IsExist(key string) bool {
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil {
		return false
	}
	if v == false {
		if _, err = rc.do("HDEL", rc.key, key); err != nil {
			return false
		}
	}
	return v
}

// Incr increase counter in redis.
func (rc *Cache) Incr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

// Decr decrease counter in redis.
func (rc *Cache) Decr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

// ClearAll clean all cache in redis. delete this redis collection.
func (rc *Cache) ClearAll() error {
	cachedKeys, err := redis.Strings(rc.do("HKEYS", rc.key))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if _, err = rc.do("DEL", str); err != nil {
			return err
		}
	}
	_, err = rc.do("DEL", rc.key)
	return err
}

// StartAndGC start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *Cache) StartAndGC(config string) error {
	var cf struct {
		Key  string   `json:"Key"`
		Conn []string `json:"Conn"`
	}
	cf.Key = DefaultKey
	if err := json.Unmarshal([]byte(config), &cf); err != nil {
		return err
	}
	if cf.Key != "" {
		rc.key = cf.Key
	}
	if len(cf.Conn) == 0 {
		return errors.New("config has no conn key")
	}
	rc.conninfo = cf.Conn
	rc.connectInit()
	c := rc.p.Get()
	defer c.Close()
	return c.Err()
}

//getNodeInfo 读取节点信息
func (rc *Cache) getNodeInfo() error {
	for _, coon := range rc.conninfo {
		if c, err := redis.Dial("tcp", coon); err == nil {
			//读取节点信息
			node, err := c.Do("CLUSTER", "NODES")
			if err != nil {
				return err
			}

			break
		}
	}
	return nil
}

// connect to redis.
func (rc *Cache) connectInit(nodes []string) {
	for _, node := range nodes {
		rc.p = append(rc.p, &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 180 * time.Second,
			Dial: func() (c redis.Conn, err error) {
				c, err = redis.Dial("tcp", node)
				if err != nil {
					return nil, err
				}
				return
			},
		})
	}
}
