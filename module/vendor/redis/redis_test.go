// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"testing"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/garyburd/redigo/redis"
)

func initCache(t *testing.T) cache.Cache {
	bm, err := cache.NewCache("redis", `{"nodes":["192.168.40.12:7000","192.168.40.12:8000","192.168.40.12:9000"],"prefix":"bee"}`)
	if err != nil {
		t.Error("init err")
	}
	return bm
}

func TestIndex(t *testing.T) {
	//updateIndex
	bm := initCache(t)

	if err := bm.Put("db:name:1:2:3:4:5:6", 1, 2*time.Second); err != nil {
		t.Error("set Error", err)
	}
	time.Sleep(3 * time.Second)
	//clearIndex
	bm.Get("db:name:1:2:3:4:5:6")
	time.Sleep(3 * time.Second)
	bm.IsExist("db:name:1:2:3:4")
	time.Sleep(3 * time.Second)
	bm.IsExist("db:name:1:2:3")
	time.Sleep(3 * time.Second)
	bm.IsExist("db:name:1:2")
	time.Sleep(3 * time.Second)
	bm.IsExist("db:name:1")
	time.Sleep(3 * time.Second)
	bm.IsExist("db:name")
	time.Sleep(3 * time.Second)
	bm.IsExist("db")
	time.Sleep(3 * time.Second)

}

func TestRedisCache(t *testing.T) {
	var err error
	bm := initCache(t)

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
