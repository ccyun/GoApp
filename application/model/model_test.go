package model

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"

	//redis 驱动

	"bbs_server/application/library/conf"
	"bbs_server/application/library/redis"

	//mysql driver

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

//Conf 配置

var isInit = false

//InitDB 初始化数据库
func InitDB() {
	if isInit == false {
		conf.InitConfig("../../cmd/base.ini")
		cache, err := cache.NewCache("redis", conf.String("cache"))
		if err != nil {
			log.Println(err)
			return
		}
		redis.Cache = cache
		Debug, _ = conf.Bool("debug")
		DBPrefix = conf.String("db_prefix")
		dsn := conf.String("db_dsn_default")
		pool, _ := conf.Int("db_pool")
		if dsn == "" || pool <= 0 {
			log.Println(errors.New("InitDB error, Configuration error.[mysql_dsn,mysql_pool]"))
			return
		}
		//最大数据库连接//最大空闲连接
		if err := orm.RegisterDataBase("default", "mysql", dsn, pool, pool); err != nil {
			log.Println(err)
			return
		}
		//注册model
		RegisterModels()
		isInit = true
	}
}

///////////////////////////////////////////////board case //////////////////////////////////////////////
func TestBoardGetOne(t *testing.T) {
	InitDB()
	a := new(Board)
	var (
		err       error
		boardInfo Board
	)
	boardInfo, err = a.GetOne(50000018)
	if err != nil {
		t.Error("model->board.GetOne err", err)
	}
	if boardInfo.ID != 50000018 {
		t.Error("model->board.GetOne err bbsinfo = nil")
	}
	log.Println(boardInfo)
}

///////////////////////////////////////////////bbs case //////////////////////////////////////////////
func TestBbsPublishScopeHandle(t *testing.T) {

	s := `{
	"discuss_ids": [50032726],
	"group_ids": [54299, 54342],
	"user_ids": [62073932],
	"tag_ids": [{
			"tag_id": 111,
			"tag_value": [111, 222, 333]
		}, {
			"tag_id": 222,
			"tag_value": [111, 222, 333]
		}
	]
}`

	var v PublishScoper

	json.Unmarshal([]byte(s), &v)

	log.Println(v)
}

func TestBbsGetOne(t *testing.T) {
	InitDB()
	a := new(Bbs)
	var (
		err     error
		bbsinfo Bbs
	)
	bbsinfo, err = a.GetOne(50001141)
	if err != nil {
		t.Error("model->bbs.GetOne err", err)
	}
	if bbsinfo.ID != 50001141 {
		t.Error("model->bbs.GetOne err bbsinfo = nil")
	}
	log.Println(bbsinfo)
}

func TestBbsUpdate(t *testing.T) {
	InitDB()
	a := new(Bbs)
	var (
		err     error
		bbsinfo Bbs
	)
	bbsinfo.ID = 50001141
	bbsinfo.Status = 1
	bbsinfo.MsgCount = 100
	if err = a.Update(bbsinfo, "Status", "MsgCount"); err != nil {
		t.Error("model->bbs.Update err", err)
	}

}

///////////////////////////////////////////////feed case //////////////////////////////////////////////

///////////////////////////////////////////////BbsTask case //////////////////////////////////////////////
func TestBbsTaskGetOne(t *testing.T) {
	InitDB()
	a := new(BbsTask)
	var (
		err         error
		bbsTaskinfo BbsTask
	)
	bbsTaskinfo, err = a.GetOne(50001597)
	if err != nil {
		t.Error("model->bbsTask.GetOne err", err)
	}
	if bbsTaskinfo.ID != 50001597 {
		t.Error("model->bbsTask.GetOne err bbsinfo = nil")
	}
	log.Println(bbsTaskinfo)
}

func TestKingShard(t *testing.T) {
	orm.Debug = true
	if err := orm.RegisterDataBase("default", "mysql", "kingshard:kingshard@tcp(192.168.197.131:9696)/kingshard?charset=utf8mb4", 10, 10); err != nil {
		t.Error("orm error", err)
	}
	o = orm.NewOrm()
	sql := ""
	values := []string{}
	for i := 13220001; i <= 1000000000; i++ {
		values = append(values, fmt.Sprintf("('%d','%d','%d','%d','%s')", (i%100), i, (i%1000), i, "bbs"))
		if i%10000 == 0 {
			sql = "insert into `bbs_unread`(`board_id`,`bbs_id`,`user_id`,`feed_id`,`feed_type`) values" + strings.Join(values, ",")
			o.Raw(sql).Exec()
			sql = ""
			values = []string{}
		}
	}
}

func TestString(t *testing.T) {
	data := []uint64{3, 2, 1, 6, 7, 5}
	sort.Sort(Uint64Slice(data))

	log.Println(data)
}
