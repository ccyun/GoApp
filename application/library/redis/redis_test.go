package redis

import (
	"bbs_server/application/library/conf"
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/chasex/redis-go-cluster"
)

var isInit = false

//initCache 初始化缓存
func initCache() {
	if isInit == false {
		conf.InitConfig("../../../cmd/_script/_script.ini")
		cache, err := cache.NewCache("redis", conf.String("cache"))
		if err != nil {
			log.Println(err)
			return
		}
		Cache = cache
		isInit = true
	}
}

//TestLock 测试锁机制
func TestLock(t *testing.T) {
	initCache()
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		n := 0
		for i := 0; i < 10000; i++ {
			key := "db:name:" + strconv.Itoa(i)
			if err := Cache.Put(key, "1", 86400*time.Second); err != nil {
				t.Error("set Error", err)
			}
			if Cache.IsExist(key) == false {
				n++
			}
		}
		t.Log("success :", (10000 - n))
		w.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		Cache.Delete("db:name*")
		time.Sleep(3 * time.Second)
		Cache.Delete("db*")
		//	time.Sleep(20 * time.Second)
		w.Done()
	}()
	w.Wait()
	Cache.ClearAll()
}

func TestIndex(t *testing.T) {
	initCache()
	if err := Cache.Put("db:name:1:2:3:4:5:6", 1, 2*time.Second); err != nil {
		t.Error("set Error", err)
	}
	if ok := Cache.IsExist("db:name:1:2:3:4:5"); ok != true {
		t.Error("index Error")
	}
	if ok := Cache.IsExist("db"); ok != true {
		t.Error("index Error")
	}
	Cache.ClearAll()
	if ok := Cache.IsExist("db:name:1:2:3:4"); ok == true {
		t.Error("index Error")
	}
	if ok := Cache.IsExist("db:name"); ok == true {
		t.Error("index Error")
	}
}

func TestRedisCache(t *testing.T) {
	initCache()
	var err error
	timeoutDuration := 10 * time.Second
	if err = Cache.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !Cache.IsExist("astaxie") {
		t.Error("check err")
	}
	time.Sleep(11 * time.Second)
	if Cache.IsExist("astaxie") {
		t.Error("check err")
	}
	if err = Cache.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if v, _ := redis.Int(Cache.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}
	if err = Cache.Incr("astaxie"); err != nil {
		t.Error("Incr Error", err)
	}
	if v, _ := redis.Int(Cache.Get("astaxie"), err); v != 2 {
		t.Error("get err")
	}
	if err = Cache.Decr("astaxie"); err != nil {
		t.Error("Decr Error", err)
	}
	if v, _ := redis.Int(Cache.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}
	Cache.Delete("astaxie")
	if Cache.IsExist("astaxie") {
		t.Error("delete err")
	}
	//test string
	if err = Cache.Put("astaxie", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !Cache.IsExist("astaxie") {
		t.Error("check err")
	}
	if v, _ := redis.String(Cache.Get("astaxie"), err); v != "author" {
		t.Error("get err")
	}
	//test GetMulti
	if err = Cache.Put("astaxie1", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !Cache.IsExist("astaxie1") {
		t.Error("check err")
	}
	vv := Cache.GetMulti([]string{"astaxie", "astaxie1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[0], nil); v != "author" {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[1], nil); v != "author1" {
		t.Error("GetMulti ERROR")
	}
	// test clear all
	if err = Cache.ClearAll(); err != nil {
		t.Error("clear all err")
	}

}

func TestCacheGet(t *testing.T) {
	initCache()
	var data []uint64
	orgids := func() []uint64 {
		return []uint64{55183, 55184, 55185, 55182, 55186, 55187, 55188, 55189, 55190, 55191, 55192, 55193, 55194, 55195, 55196, 55197, 55198, 55199, 55200, 55201, 55202, 55203, 55204, 55205, 55206, 55208, 55207, 55209, 55210, 55211, 55212, 55213, 55214, 55215, 55216, 55217, 55218, 55219, 55220, 55221, 55222, 55225, 55223, 55224, 55226, 55227, 55228, 63686573}
	}()

	cache := NewCache(fmt.Sprintf("U%s", "000000"), "GetAllUserIDsByOrgIDs", orgids)
	log.Println(cache.key)
	log.Println(cache.Get(&data))
	log.Println(data)
}
