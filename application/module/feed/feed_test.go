package feed

import (
	"log"
	"testing"

	"encoding/json"

	"github.com/astaxie/beego/config"
	"github.com/ccyun/GoApp/application/module/pic"
)

//Conf 配置
var Conf config.Configer
var isInit = false

//InitDB 初始化数据库
func InitModule() {
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
		conf, err := config.NewConfig("ini", "../../../cmd/TaskScript/conf.ini")
		if err != nil {
			return err
		}
		Conf = conf
		return nil
	}, func() error {
		config := map[string]string{
			"server_name": Conf.String("server_name"),
			"app_domain":  Conf.String("app_domain"),
			"app_path":    Conf.String("app_path"),
			"feed_icons":  Conf.String("feed_icons"),
		}
		if err := Init(config); err != nil {
			return err
		}
		if err := pic.Init(config); err != nil {
			return err
		}
		return nil
	})
	isInit = true
}

//TestNewBbs
func TestNewBbs(t *testing.T) {
	InitModule()
	feedData, err := NewBbs("bbs", Customizer{
		BoardID:        50000075,
		BoardName:      "王磊测试7",
		Avatar:         "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gN18yNDBweC5wbmdeXl50YW5naGRmc15eXjhmMmMyNTVmMmRkZWNhYWE0ODc1N2U4MjVmMThlMDdjXl5edGFuZ2hkZnNeXl4xMjE4Nw$&u=62051318",
		DiscussID:      0,
		BbsID:          50001140,
		FeedID:         1214,
		Title:          "111",
		Description:    "111",
		Thumb:          "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMi5wbmdeXl50YW5naGRmc15eXmIwNGQyMTIyYjVjOTFiMzkzZTQzNTcwZjUxNDU4ZjhlXl5edGFuZ2hkZnNeXl43ODY0MTA$&u=62051318",
		UserID:         63672505,
		Type:           "default",
		Category:       "bbs",
		CommentEnabled: 1,
		CreatedAt:      1481879395263,
	})
	if err != nil {
		t.Error("module->feed.NewBbs error:", err)
	}

	a, err := json.Marshal(feedData)
	if err != nil {
		t.Error("module->feed.NewBbs error:", err)
	}
	log.Println(string(a))
}

func TestNewTask(t *testing.T) {
	InitModule()
	feedData, err := NewTask("task", Customizer{
		BoardID:        50000075,
		BoardName:      "王磊测试7",
		Avatar:         "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gN18yNDBweC5wbmdeXl50YW5naGRmc15eXjhmMmMyNTVmMmRkZWNhYWE0ODc1N2U4MjVmMThlMDdjXl5edGFuZ2hkZnNeXl4xMjE4Nw$&u=62051318",
		DiscussID:      0,
		BbsID:          50001140,
		FeedID:         1214,
		Title:          "111",
		Description:    "111",
		Thumb:          "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMi5wbmdeXl50YW5naGRmc15eXmIwNGQyMTIyYjVjOTFiMzkzZTQzNTcwZjUxNDU4ZjhlXl5edGFuZ2hkZnNeXl43ODY0MTA$&u=62051318",
		UserID:         63672505,
		Type:           "default",
		Category:       "bbs",
		CommentEnabled: 1,
		CreatedAt:      1481879395263,
	}, CustomizeTasker{
		EndTime:      1481979595263,
		Status:       0,
		AllowExpired: 1,
	})
	if err != nil {
		t.Error("module->feed.NewBbs error:", err)
	}

	a, err := json.Marshal(feedData)
	if err != nil {
		t.Error("module->feed.NewBbs error:", err)
	}
	log.Println(string(a))
}

func TestNewTaskReply(t *testing.T) {
	InitModule()
	feedData, err := NewTask("taskReply", Customizer{
		BoardID:        50000075,
		BoardName:      "王磊测试7",
		Avatar:         "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_pu5jorqTlpLTlg48gN18yNDBweC5wbmdeXl50YW5naGRmc15eXjhmMmMyNTVmMmRkZWNhYWE0ODc1N2U4MjVmMThlMDdjXl5edGFuZ2hkZnNeXl4xMjE4Nw$&u=62051318",
		DiscussID:      0,
		BbsID:          50001140,
		FeedID:         1214,
		Title:          "111",
		Description:    "111",
		Thumb:          "http://testcloud.quanshi.com:80/ucfserver/hddown?fid=NjIwNTEzMTgvOC_mjqjojZAgMi5wbmdeXl50YW5naGRmc15eXmIwNGQyMTIyYjVjOTFiMzkzZTQzNTcwZjUxNDU4ZjhlXl5edGFuZ2hkZnNeXl43ODY0MTA$&u=62051318",
		UserID:         63672505,
		Type:           "default",
		Category:       "bbs",
		CommentEnabled: 1,
		CreatedAt:      1481879395263,
	}, CustomizeTasker{
		EndTime:      1481979595263,
		Status:       0,
		AllowExpired: 1,
	})
	if err != nil {
		t.Error("module->feed.NewBbs error:", err)
	}

	a, err := json.Marshal(feedData)
	if err != nil {
		t.Error("module->feed.NewBbs error:", err)
	}
	log.Println(string(a))
}
