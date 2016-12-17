package application

import (
	"errors"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ccyun/GoApp/application/library/httpcurl"
	_ "github.com/ccyun/GoApp/application/library/redis"
	"github.com/ccyun/GoApp/application/model"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Conf 配置
var Conf config.Configer

func init() {
	func(f ...func() error) {
		for _, function := range f {
			if err := function(); err != nil {
				panic(err)
			}
		}
	}(InitConfig, InitLog, InitDB, InitHTTPCurl, InitCache)
}

//InitConfig 初始化配置
func InitConfig() error {
	conf, err := config.NewConfig("ini", "conf.ini")

	if err != nil {
		return err
	}
	Conf = conf
	return nil
}

//InitLog 初始化log
func InitLog() error {

	logs.SetLogger("file", `{"filename":"`+Conf.String("log_path")+`/`+time.Now().Format("2006-01-02")+`.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Async(1e3)
	return nil
}

//InitDB 初始化数据库
func InitDB() error {
	var err error
	debug, _ := Conf.Bool("debug")
	model.Debug = debug
	model.DBType = Conf.String("db_type")
	model.DBPrefix = Conf.String("db_prefix")
	dsn := Conf.String("db_dsn")
	pool, _ := Conf.Int("db_pool")
	if dsn == "" || pool <= 0 {
		return errors.New("InitDB error, Configuration error.[mysql_dsn,mysql_pool]")
	}
	switch model.DBType {
	case "mysql":
		err = orm.RegisterDriver(model.DBType, orm.DRMySQL)
	case "sqlite":
		err = orm.RegisterDriver(model.DBType, orm.DRSqlite)
	case "oracle":
		err = orm.RegisterDriver(model.DBType, orm.DROracle)
	case "pgsql":
		err = orm.RegisterDriver(model.DBType, orm.DRPostgres)
	case "TiDB":
		err = orm.RegisterDriver(model.DBType, orm.DRTiDB)
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
}

//InitHTTPCurl 初始化数据库
func InitHTTPCurl() error {
	httpcurl.UMSLoginURL = Conf.String("ums_login_url")
	httpcurl.UMSBusinessURL = Conf.String("ums_business_url")
	return nil
}

//InitCache 初始化缓存
func InitCache() error {
	ca, err := cache.NewCache("redis", Conf.String("cache"))
	if err != nil {
		return err
	}
	model.Cache = ca
	httpcurl.Cache = ca
	ca.Put("fdsfdsf", "fdsfdsf", 10*time.Second)
	return nil
}
