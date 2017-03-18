package syslog2

import (
	"testing"

	"github.com/astaxie/beego/logs"
)

func TestSyslog(t *testing.T) {
	log := logs.NewLogger(10000)
	log.SetLogger("syslog", `{"tag":"bbsapp","path":"D:/Go/WorkSpace/src/github.com/ccyun/GoApp/cmd/_script/logs"}`)
	log.Info("sendmail critical")
}
