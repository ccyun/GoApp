package application

import (
	"encoding/json"
	"errors"
	"runtime"
	"time"

	"github.com/astaxie/beego/cache"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/ccyun/GoApp/application/library/hbase"
	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/library/neo4j"
	"github.com/ccyun/GoApp/application/model"
	//syslog 驱动
	_ "github.com/ccyun/GoApp/application/library/log"
	//redis 驱动
	"github.com/ccyun/GoApp/application/library/redis"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Conf 配置
var Conf config.Configer

func init() {
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(InitConfig, InitLog, InitDB, InitHTTPCurl, InitCache, InitHbase)
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
	if runtime.GOOS == "linux" || Conf.String("log_type") == "syslog" {

	} else {
		logs.SetLogger("file", `{"filename":"`+Conf.String("log_path")+`/`+time.Now().Format("2006-01-02")+`.log"}`)
		logs.EnableFuncCallDepth(true)
		logs.SetLogFuncCallDepth(4)
		logs.Async(1e3)
	}
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

	//初始化ums配置
	httpcurl.UMSLoginURL = Conf.String("ums_login_url")
	httpcurl.UMSBusinessURL = Conf.String("ums_business_url")

	//初始化uc配置
	httpcurl.UcOpenAPIURL = Conf.String("uc_open_api_url")
	httpcurl.UcAPPID = Conf.String("uc_open_appid")
	httpcurl.UcPaddword = Conf.String("uc_open_password")

	return nil
}

//InitCache 初始化缓存
func InitCache() error {
	cache, err := cache.NewCache("redis", Conf.String("cache"))
	if err != nil {
		return err
	}
	redis.Cache = cache
	return nil
}

//InitHbase 初始化hbase
func InitHbase() error {
	var (
		err    error
		config struct {
			Host string `json:"host"`
			Port string `json:"port"`
			Pool int    `json:"pool"`
		}
	)
	if err = json.Unmarshal([]byte(Conf.String("hbase")), &config); err != nil {
		return err
	}
	return hbase.InitHbase(config.Host, config.Port, config.Pool)
}

//InitNeo4j 初始化Neo4j
func InitNeo4j() error {
	var (
		err    error
		config struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			UserName string `json:"username"`
			Password string `json:"password"`
			Pool     int    `json:"pool"`
		}
	)
	if err = json.Unmarshal([]byte(Conf.String("neo4j")), &config); err != nil {
		return err
	}
	return neo4j.Init(config.Host, config.Port, config.UserName, config.Password, config.Pool)
}
