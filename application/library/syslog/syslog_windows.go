package syslog

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

//SysLogWriter syslog结构体
type SysLogWriter struct {
}

func newSysLogWriter() logs.Logger {
	return &SysLogWriter{}
}

//Init 初始化配置
func (s *SysLogWriter) Init(config string) error {
	return nil
}

//WriteMsg 写log
func (s *SysLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	fmt.Println(fmt.Sprintf("%s %s", when.Format("2006-01-02 15:04:05"), msg))
	return nil
}

//Destroy
func (s *SysLogWriter) Destroy() {}

//Flush d
func (s *SysLogWriter) Flush() {}

func init() {
	logs.Register("syslog", newSysLogWriter)
}
