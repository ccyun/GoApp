package application

import (
	"errors"
	"runtime"
	"time"

	"github.com/astaxie/beego/cache"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/ccyun/GoApp/application/library/conf"
	"github.com/ccyun/GoApp/application/library/hbase"
	"github.com/ccyun/GoApp/application/library/httpcurl"
	"github.com/ccyun/GoApp/application/model"
	"github.com/ccyun/GoApp/application/module/feed"
	"github.com/ccyun/GoApp/application/module/pic"
	//syslog 驱动
	_ "github.com/ccyun/GoApp/application/library/syslog2"
	//redis 驱动
	"github.com/ccyun/GoApp/application/library/redis"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(InitConfig, InitLog, InitDB, InitHTTPCurl, InitCache, InitHbase, InitPackage)
}

//InitConfig 初始化配置
func InitConfig() error {
	return conf.InitConfig()
}

//InitLog 初始化log
func InitLog() error {
	if runtime.GOOS == "linux" && conf.String("log_type") == "syslog" {
		logs.SetLogger("syslog", `{"tag":"`+conf.String("log_tag")+`"}`)
	} else {
		logs.SetLogger("file", `{"filename":"`+conf.String("log_path")+`/`+time.Now().Format("2006-01-02")+`.log"}`)
	}
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(4)
	logs.Async(1e3)
	return nil
}

//InitDB 初始化数据库
func InitDB() error {
	var err error
	debug, _ := conf.Bool("debug")
	model.Debug = debug
	model.DBPrefix = conf.String("db_prefix")
	dsn := conf.String("db_dsn_default")
	pool, _ := conf.Int("db_pool")
	if dsn == "" || pool <= 0 {
		return errors.New("InitDB error, Configuration error.[mysql_dsn,mysql_pool]")
	}
	err = orm.RegisterDriver(model.DBType, orm.DRMySQL)
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
	httpcurl.UMSLoginURL = conf.String("ums_login_url")
	httpcurl.UMSBusinessURL = conf.String("ums_business_url")

	//初始化uc配置
	httpcurl.UcOpenAPIURL = conf.String("uc_open_api_url")
	httpcurl.UcAPPID = conf.String("uc_open_appid")
	httpcurl.UcPaddword = conf.String("uc_open_password")

	//初始化ucc配置
	httpcurl.UccServerURL = conf.String("uccserver_url")

	return nil
}

//InitCache 初始化缓存
func InitCache() error {
	cache, err := cache.NewCache("redis", conf.String("cache"))
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
	if err = conf.JSON("hbase", &config); err != nil {
		return err
	}
	return hbase.Init(config.Host, config.Port, config.Pool)
}

//InitPackage 初始化其他包
func InitPackage() error {
	config := map[string]string{
		"server_name": conf.String("server_name"),
		"app_domain":  conf.String("app_domain"),
		"app_path":    conf.String("app_path"),
		"feed_icons":  conf.String("feed_icons"),
	}
	if err := feed.Init(config); err != nil {
		return err
	}
	if err := pic.Init(config); err != nil {
		return err
	}
	return nil
}
