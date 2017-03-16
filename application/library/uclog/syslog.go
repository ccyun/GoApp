// +build !windows,!nacl,!plan9

package uclog

// import (
// 	"fmt"
// 	"log/syslog"
// 	"path"
// 	"runtime"

// 	"github.com/huandu/goroutine"
// )

// var (
// 	w           *syslog.Writer
// 	processName string
// 	isSyslog    bool
// )

// func Initialize(procName string, logpath string, loglevel string) (ok bool) {
// 	if logpath == "" || loglevel == "" {
// 		return false
// 	}
// 	if logpath == "syslog" {
// 		isSyslog = true
// 		t, err := syslog.New(syslog.LOG_LOCAL0, procName)
// 		if err != nil {
// 			fmt.Printf("open common syslog fail, %s", err.Error())
// 			return false
// 		}
// 		w = t
// 		processName = procName

// 		t, err = syslog.New(syslog.LOG_LOCAL6, procName)
// 		if err != nil {
// 			fmt.Printf("open monitor syslog fail, %s", err.Error())
// 			return false
// 		}
// 		mw = t

// 		switch loglevel {
// 		case "deubg":
// 			level = LOG_LEVEL_DEBUG
// 		case "info":
// 			level = LOG_LEVEL_INFO
// 		case "warn":
// 			level = LOG_LEVEL_WARN
// 		case "error":
// 			level = LOG_LEVEL_ERROR
// 		case "critical":
// 			level = LOG_LEVEL_CRITICAL
// 		default:
// 			level = LOG_LEVEL_DEBUG
// 		}
// 		ok = true
// 	} else {
// 		ok = initFileLog(logpath, loglevel)
// 	}
// 	return
// }

// func GetProcName() string {
// 	return processName
// }

// func getCaller() string {
// 	pc, file, line, ok := runtime.Caller(2)
// 	if ok {
// 		_, filename := path.Split(file)
// 		fnname := runtime.FuncForPC(pc).Name()
// 		return fmt.Sprintf("[%s:%d:%s] ", filename, line, fnname)
// 	}

// 	return ""
// }
// func getGid() string {
// 	var r string
// 	defer func() {
// 		if err := recover(); err != nil {
// 			r = ""
// 		}
// 	}()
// 	gid := goroutine.GoroutineId()
// 	r = fmt.Sprintf("<gid:%d>", gid)
// 	return r
// }

// func (this *UcLog) genLogPrefix() string {

// 	this.logIndex++
// 	headerPrefix := fmt.Sprintf("<requestid:%s>%s<lognum:%d> ", this.requestId, getGid(), this.logIndex)
// 	for _, v := range this.header {
// 		headerPrefix = headerPrefix + "<" + v + "> "
// 	}

// 	callerPrefix := ""
// 	pc, file, line, ok := runtime.Caller(2)
// 	if ok {
// 		_, filename := path.Split(file)
// 		fnname := runtime.FuncForPC(pc).Name()
// 		callerPrefix = fmt.Sprintf("<%s:%d:%s> ", filename, line, fnname)
// 	}

// 	return headerPrefix + callerPrefix
// }

// func (this *UcLog) Log_Debug(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_DEBUG {
// 			w.Debug(this.genLogPrefix() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Debug(this.genLogPrefix()+format, v...)
// 	}
// }

// func (this *UcLog) Log_Info(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_INFO {
// 			w.Info(this.genLogPrefix() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Informational(this.genLogPrefix()+format, v...)
// 	}
// }

// func (this *UcLog) Log_Warn(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_WARN {
// 			w.Warning(this.genLogPrefix() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Warning(this.genLogPrefix()+format, v...)
// 	}
// }

// func (this *UcLog) Log_Error(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_ERROR {
// 			w.Err(this.genLogPrefix() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Error(this.genLogPrefix()+format, v...)
// 	}
// }

// func (this *UcLog) Log_Critical(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_CRITICAL {
// 			w.Crit(this.genLogPrefix() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Critical(this.genLogPrefix()+format, v...)
// 	}
// }

// func Debug(format string, v ...interface{}) {

// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_DEBUG {
// 			w.Debug(getGid() + getCaller() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Debug(getGid()+getCaller()+format, v...)
// 	}
// }

// func Info(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_INFO {
// 			w.Info(getGid() + getCaller() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Informational(getGid()+getCaller()+format, v...)
// 	}
// }

// func Warn(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_WARN {
// 			w.Warning(getGid() + getCaller() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Warning(getGid()+getCaller()+format, v...)
// 	}
// }

// func Error(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_ERROR {
// 			w.Err(getGid() + getCaller() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Error(getGid()+getCaller()+format, v...)
// 	}
// }

// func Critical(format string, v ...interface{}) {
// 	if isSyslog {
// 		if w != nil && level >= LOG_LEVEL_CRITICAL {
// 			w.Crit(getGid() + getCaller() + fmt.Sprintf(format, v...))
// 		}
// 	} else {
// 		log.Critical(getGid()+getCaller()+format, v...)
// 	}
// }
func init() {

}
