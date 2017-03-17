package log2

import (
	"testing"

	"github.com/astaxie/beego/logs"
)

func TestSyslog(t *testing.T) {
	log := logs.NewLogger(10000)
	log.SetLogger("syslog", `{"tag":"bbsapp"}`)
	log.Info("sendmail critical")

}
