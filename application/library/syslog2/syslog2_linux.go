package syslog2

import (
	"encoding/json"
	"fmt"
	"time"

	"log/syslog"

	"github.com/astaxie/beego/logs"
)

//SysLogWriter syslog结构体
type SysLogWriter struct {
	w     *syslog.Writer
	wc    *syslog.Writer
	Level int    `json:"level"`
	Tag   string `json:"tag"`
}

func newSysLogWriter() logs.Logger {
	s := &SysLogWriter{
		Level: logs.LevelTrace,
	}
	return s
}

//Init 初始化配置
func (s *SysLogWriter) Init(config string) error {
	var err error
	if err = json.Unmarshal([]byte(config), s); err != nil {
		return err
	}
	if s.w, err = syslog.New(syslog.LOG_LOCAL0, s.Tag); err != nil {
		return err
	}
	if s.wc, err = syslog.New(syslog.LOG_LOCAL5, s.Tag); err != nil {
		return err
	}
	return nil
}

//WriteMsg 写log
func (s *SysLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	var err error
	if level > s.Level {
		return nil
	}
	logMsg := fmt.Sprintf("%s %s", when.Format("2006-01-02 15:04:05"), msg)
	if level == logs.LevelCritical {
		err = s.wc.Crit(logMsg)
	}
	switch level {
	case logs.LevelEmergency:
		err = s.w.Emerg(logMsg)
	case logs.LevelAlert:
		err = s.w.Alert(logMsg)
	case logs.LevelCritical:
		err = s.w.Crit(logMsg)
	case logs.LevelError:
		err = s.w.Err(logMsg)
	case logs.LevelWarning:
		err = s.w.Warning(logMsg)
	case logs.LevelNotice:
		err = s.w.Notice(logMsg)
	case logs.LevelInfo:
		err = s.w.Info(logMsg)
	case logs.LevelDebug:
		err = s.w.Debug(logMsg)
	}
	return err
}

//Destroy 注销
func (s *SysLogWriter) Destroy() {
	s.w.Close()
}

//Flush 刷新
func (s *SysLogWriter) Flush() {}

func init() {
	logs.Register("syslog", newSysLogWriter)
}
