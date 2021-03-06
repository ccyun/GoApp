package application

import (
	"errors"
	"runtime"

	"github.com/astaxie/beego/cache"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"bbs_server/application/library/conf"
	"bbs_server/application/library/httpcurl"
	"bbs_server/application/model"
	"bbs_server/application/module/feed"
	"bbs_server/application/module/pic"
	//syslog 驱动
	_ "bbs_server/application/library/syslog2"
	//redis 驱动
	"bbs_server/application/library/redis"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//io 重要参数
type io struct {
	requestID string
}

func init() {
	func(funcs ...func() error) {
		for _, f := range funcs {
			if err := f(); err != nil {
				panic(err)
			}
		}
	}(
		InitConfig,
		InitLog,
		InitDB,
		InitHTTPCurl,
		InitCache,
		//InitHbase,
		InitPackage,
		InitRouter,
	)
}

//InitConfig 初始化配置
func InitConfig() error {
	return conf.InitConfig("")
}

//InitLog 初始化log
func InitLog() error {
	if runtime.GOOS == "linux" && conf.String("log_type") == "syslog" {
		logs.SetLogger("syslog", `{"tag":"`+conf.String("log_tag")+`"}`)
	} else {
		logs.SetLogger(logs.AdapterFile, `{"filename":"`+conf.String("log_path")+`/`+conf.String("log_tag")+`.log"}`)
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
	defaultDSN := conf.String("db_dsn_default")
	msgDSN := conf.String("db_dsn_msg")
	pool, _ := conf.Int("db_pool")
	if defaultDSN == "" || pool <= 0 {
		return errors.New("InitDB error, Configuration error.[mysql_dsn,mysql_pool]")
	}
	//最大数据库连接//最大空闲连接
	if err = orm.RegisterDataBase("default", "mysql", defaultDSN, pool, pool); err != nil {
		return err
	}
	if err = orm.RegisterDataBase("msg", "mysql", msgDSN, pool, pool); err != nil {
		return err
	}
	o := new(io)
	orm.DebugLog = orm.NewLog(o)
	//注册model
	model.RegisterModels()
	return nil
}

//Write io.Writer 用于orm sql输出
func (o *io) Write(b []byte) (n int, err error) {
	logs.Info(string(b))
	return 0, nil
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

	//初始化solr配置
	httpcurl.SolrServerURL = conf.String("solr_url")

	//初始化solr配置
	httpcurl.BillServerURL = conf.String("bill_url")
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
	return pic.Init(config)
}
