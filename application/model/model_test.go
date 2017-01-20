package model

import (
	"log"
	"testing"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"

	//redis 驱动
	"github.com/ccyun/GoApp/application/library/hbase"
	_ "github.com/ccyun/GoApp/application/library/redis"
	//mysql driver

	_ "github.com/go-sql-driver/mysql"
)

var isInit = false

//InitDB 初始化数据库
func InitDB() error {
	if isInit == true {
		return nil
	}
	var err error

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
	_, err = cache.NewCache("redis", `{"nodes":["192.168.40.12:7000","192.168.40.12:8000","192.168.40.12:9000"],"prefix":"bee"}`)
	if err != nil {
		return err
	}

	RegisterModels()
	isInit = true
	return nil
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
	log.Println(v)

}

func TestBbsGetOne(t *testing.T) {
	InitDB()
	a := new(Bbs)
	var (
		err     error
		bbsinfo Bbs
		//bbsinfo2 Bbs
	)

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
	hbase.InitHbase("192.168.197.128", "9090", 10)
	a := new(Feed)
	var userIDs []uint64
	for i := 100000; i < 100001; i++ {
		userIDs = append(userIDs, uint64(i))
	}
	for i := 10; i < 20; i++ {
		for ii := 100; ii < 200; ii++ {
			for iii := 10; iii < 20; iii++ {
				feedData := Feed{
					BoardID:   uint64(i),
					BbsID:     uint64(ii),
					ID:        uint64(iii),
					FeedType:  "task",
					MsgID:     0,
					CreatedAt: 1481879624794,
				}
				a.SaveHbase(userIDs, feedData)
			}
		}
	}
}
