package model

import (
	"errors"
	"log"
	"testing"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/config"

	//redis 驱动

	"bbs_server/application/library/redis"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Conf 配置
var Conf config.Configer
var isInit = false

//InitDB 初始化数据库
func InitDB() {
	if isInit == true {
		return
	}
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(func() error {
		conf, err := config.NewConfig("ini", "../../cmd/TaskScript/conf.ini")
		if err != nil {
			return err
		}
		Conf = conf
		return nil
	}, func() error {
		cache, err := cache.NewCache("redis", Conf.String("cache"))
		if err != nil {
			return err
		}
		redis.Cache = cache
		return nil
	}, func() error {
		var err error
		debug, _ := Conf.Bool("debug")
		Debug = debug
		DBType = Conf.String("db_type")
		DBPrefix = Conf.String("db_prefix")
		dsn := Conf.String("db_dsn")
		pool, _ := Conf.Int("db_pool")
		if dsn == "" || pool <= 0 {
			return errors.New("InitDB error, Configuration error.[mysql_dsn,mysql_pool]")
		}
		switch DBType {
		case "mysql":
			err = orm.RegisterDriver(DBType, orm.DRMySQL)
		case "sqlite":
			err = orm.RegisterDriver(DBType, orm.DRSqlite)
		case "oracle":
			err = orm.RegisterDriver(DBType, orm.DROracle)
		case "pgsql":
			err = orm.RegisterDriver(DBType, orm.DRPostgres)
		case "TiDB":
			err = orm.RegisterDriver(DBType, orm.DRTiDB)
		}
		if err != nil {
			return err
		}
		//最大数据库连接//最大空闲连接
		err = orm.RegisterDataBase("default", "mysql", dsn, pool, pool)
		if err != nil {
			return err
		}
		return nil
	})
	RegisterModels()
	isInit = true
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
	a := new(Bbs)
	s := `{"discuss_ids":["50032726"],"group_ids":["54299","54342"],"user_ids":["62073932"]}`
	ss := `[11,22,33,44,55]`
	v, vv, err := a.publishScopeHandle(s, ss)
	if err != nil {
		t.Error("model->bbs.publishScopeHandle error1", err)
	}
	if v.GroupIDs[1] != 54342 {
		t.Error("model->bbs.publishScopeHandle error2", err)
	}
	if v.UserIDs[0] != 62073932 {
		t.Error("model->bbs.publishScopeHandle error3", err)
	}
	if vv[3] != 44 {
		t.Error("model->bbs.publishScopeHandle error4", err)
	}
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

///////////////////////////////////////////////BbsTaskReply case //////////////////////////////////////////////
func TestGetReplyUserIDs(t *testing.T) {
	InitDB()
	a := new(BbsTaskReply)
	userids, err := a.GetReplyUserIDs(50001129)
	if err != nil {
		t.Error("model->GetReplyUserIDs error:", err)
	}
	if userids[0] != 62051317 {
		t.Error("model->GetReplyUserIDs error:!=62051317")
	}
}
