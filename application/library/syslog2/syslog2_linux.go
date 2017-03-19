package syslog2

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"log/syslog"

	"github.com/astaxie/beego/logs"
)

//SysLogWriter syslog结构体
type SysLogWriter struct {
	w     *syslog.Writer
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
	err := json.Unmarshal([]byte(config), s)
	if err != nil {
		return err
	}
	w, err := syslog.New(syslog.LOG_LOCAL0, s.Tag)
	if err != nil {
		return err
	}
	s.w = w
	return nil
}

//WriteMsg 写log
func (s *SysLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > s.Level {
		return nil
	}
	_, err := s.w.Write([]byte(fmt.Sprintf("%s %s", when.Format("2006-01-02 15:04:05"), msg)))
	log.Println(err)
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
