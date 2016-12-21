package redis

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils"
	"github.com/chasex/redis-go-cluster"
)

func initRedis(t *testing.T) cache.Cache {
	bm, err := cache.NewCache("redis", `{"nodes":["192.168.40.12:7000","192.168.40.12:8000","192.168.40.12:9000"],"prefix":"bee"}`)
	if err != nil {
		t.Error("init err")
	}
	return bm
}
func TestBatch1(t *testing.T) {
	bm := initRedis(t)
	var w sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	for index := 0; index < 500; index++ {
		w.Add(1)
		go func() {
			key := "db:name:"
			key1 := string(utils.RandomCreateBytes(5))
			key2 := string(utils.RandomCreateBytes(5))
			key3 := string(utils.RandomCreateBytes(5))
			key4 := string(utils.RandomCreateBytes(5))
			key += key1 + ":" + key2 + ":" + key3 + ":" + key4

			if err := bm.Put(key, "true", 86400*time.Second); err != nil {
				t.Error("cache put err", err)
			}
			if ok := bm.IsExist(key); ok == false {
				t.Error("cache put err")
			}
			w.Done()
		}()
	}
	w.Wait()
	//time.Sleep(10 * time.Second)
	if err := bm.ClearAll(); err != nil {
		t.Error("ClearAll Error", err)
	}
}

func TestBatch2(t *testing.T) {
	bm := initRedis(t)
	var w sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	for index := 0; index < 500; index++ {
		w.Add(1)
		go func() {
			key := "db:name:"
			key1 := string(utils.RandomCreateBytes(5))
			key2 := string(utils.RandomCreateBytes(5))
			key3 := string(utils.RandomCreateBytes(5))
			key4 := string(utils.RandomCreateBytes(5))
			key += key1 + ":" + key2 + ":" + key3 + ":" + key4

			if err := bm.Put(key, "true", 86400*time.Second); err != nil {
				t.Error("cache put err", err)
			}
			if ok := bm.IsExist(key); ok == false {
				t.Error("cache put err")
			}
			if err := bm.Delete(key); err != nil {
				t.Error("cache Delete err", err)
			}
			w.Done()
		}()
	}
	w.Wait()
	if err := bm.ClearAll(); err != nil {
		t.Error("ClearAll Error", err)
	}
}

func TestIndex(t *testing.T) {
	//updateIndex
	bm := initRedis(t)
	if err := bm.Put("db:name:1:2:3:4:5:6", 1, 2*time.Second); err != nil {
		t.Error("set Error", err)
	}
	if ok := bm.IsExist("db:name:1:2:3:4:5"); ok != true {
		t.Error("index Error")
	}

	if ok := bm.IsExist("db"); ok != true {
		t.Error("index Error")
	}
	bm.ClearAll()
	if ok := bm.IsExist("db:name:1:2:3:4"); ok == true {
		t.Error("index Error")
	}
	if ok := bm.IsExist("db:name"); ok == true {
		t.Error("index Error")
	}
}

func TestRedisCache(t *testing.T) {
	var err error
	bm := initRedis(t)

	timeoutDuration := 10 * time.Second
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}
	time.Sleep(11 * time.Second)
	if bm.IsExist("astaxie") {
		t.Error("check err")
	}
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}
	if err = bm.Incr("astaxie"); err != nil {
		t.Error("Incr Error", err)
	}
	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 2 {
		t.Error("get err")
	}
	if err = bm.Decr("astaxie"); err != nil {
		t.Error("Decr Error", err)
	}
	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}
	bm.Delete("astaxie")
	if bm.IsExist("astaxie") {
		t.Error("delete err")
	}
	//test string
	if err = bm.Put("astaxie", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}
	if v, _ := redis.String(bm.Get("astaxie"), err); v != "author" {
		t.Error("get err")
	}
	//test GetMulti
	if err = bm.Put("astaxie1", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie1") {
		t.Error("check err")
	}
	vv := bm.GetMulti([]string{"astaxie", "astaxie1"})
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
	if err = bm.ClearAll(); err != nil {
		t.Error("clear all err")
	}

}
