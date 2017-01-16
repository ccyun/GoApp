package model

import (
	"log"
	"testing"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"

	//redis 驱动
	_ "github.com/ccyun/GoApp/application/library/redis"
	//mysql driver

	"time"

	_ "github.com/go-sql-driver/mysql"
)

var isInit = false

//InitDB 初始化数据库
func InitDB() error {
	if isInit == true {
		return nil
	}
	var err error
	RequestID = "RequestID1234567890"
	Debug = true
	DBType = "mysql"
	DBPrefix = "bbs_"
	dsn := "root:root@tcp(127.0.0.1:3306)/bee_app?charset=utf8mb4"
	err = orm.RegisterDriver(DBType, orm.DRMySQL)
	if err != nil {
		return err
	}
	//最大数据库连接//最大空闲连接
	err = orm.RegisterDataBase("default", "mysql", dsn, 10, 10)
	if err != nil {
		return err
	}
	ca, err := cache.NewCache("redis", `{"nodes":["192.168.40.12:7000","192.168.40.12:8000","192.168.40.12:9000"],"prefix":"bee"}`)
	if err != nil {
		return err
	}
	Cache = ca
	RegisterModels()
	isInit = true
	return nil
}

///////////////////////////////////////////model case //////////////////////////////////////////////
func TestModelCache(t *testing.T) {
	InitDB()
	c := newCache("tableName", "getOne")
	a := map[string]string{"fdf": "fsdfds"}
	var aa map[string]string
	aa = make(map[string]string)
	c.setCache(a)
	if ok := c.getCache(&aa); ok != true {
		t.Error("model->Test_model_cache GetCache error", aa)
	}
	time.Sleep(1 * time.Second)
	if ok := c.clearCache("tableName"); ok != true {
		t.Error("model->Test_model_cache ClearCache error")
	}
}

//TestModelL 测试model语言
func TestModelL(t *testing.T) {
	InitDB()
	log := L("fdsfsd")
	if log != "RequestID1234567890  fdsfsd" {
		t.Error("model languge error")
	}

}

///////////////////////////////////////////////bbs case //////////////////////////////////////////////
func TestBbsPublishScopeHandle(t *testing.T) {
	InitDB()
	a := new(Bbs)
	s := `{"discuss_ids":["50032726"],"group_ids":["54299","54342"],"user_ids":["62073932"]}`
	v, err := a.publishScopeHandle(s)
	if err != nil {
		t.Error("model->bbs.publishScopeHandle err", err)
	}
	if v.DiscussIDs[0] != 50032726 || v.GroupIDs[0] != 54299 || v.GroupIDs[1] != 54342 || v.UserIDs[0] != 62073932 {
		t.Error("model->bbs.publishScopeHandle err", s, v)
	}
}

func TestAfterUpdate(t *testing.T) {
	InitDB()
	log.Println(AfterUpdate("bbs", 1, nil))
}

func TestBbsGetOne(t *testing.T) {
	InitDB()
	a := new(Bbs)
	var (
		err     error
		bbsinfo Bbs
		//bbsinfo2 Bbs
	)

	SiteID = 71058
	bbsinfo, err = a.GetOne(15)
	if err != nil {
		t.Error("model->bbs.Test_bbs_getOne err", err)
	}
	if bbsinfo.ID != 15 {
		t.Error("model->bbs.Test_bbs_getOne err bbsinfo = nil")
	}
	log.Println(bbsinfo)
	// c := NewCache("bbs", "GetOne")
	// c.GetCache(15, &bbsinfo2)
	// log.Println(bbsinfo2)
	// if bbsinfo2.ID != 15 {
	// 	t.Error("model->bbs.Test_bbs_getOne cache err bbsinfo = nil")
	// }

}

///////////////////////////////////////////////feed case //////////////////////////////////////////////
func TestSaveHbase(t *testing.T) {
	a := new(Feed)
	var userIDs []uint64
	for i := 100000; i < 101000; i++ {
		userIDs = append(userIDs, uint64(i))
	}
	feedData := HbaseFeed{
		BoardID:        111,
		BbsID:          123,
		FeedID:         113,
		FeedType:       "bbs",
		MsgID:          0,
		DiscussID:      0,
		Title:          "测试hbase广播标题",
		Description:    "测试hbase广播摘要",
		Thumb:          "http://testcloudb.quanshi.com:80/ucfserver/hddown?fid=NjE1MDY4MDQvOC90ZW1wX25vdGljZV5eXnRhbmdoZGZzXl5eYWQ2ZmQ5MTg2NzIxZjRlODE2MzliZmNiOGQ5ZWY4ZjleXl50YW5naGRmc15eXjYwMjA3OA$&u=61506804",
		Category:       "bbs",
		Type:           "default",
		CommentEnabled: 1,
		CreatedAt:      1481879624794,
	}
	a.SaveHbase(userIDs, feedData)
}

func TestMakeRowkey(t *testing.T) {
	log.Println(makeRowkey(45441266))
}
