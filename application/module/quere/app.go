package queue

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"bbs_server/application/library/conf"

	"github.com/astaxie/beego/utils"
)

//App 流程控制
type App struct {
	thread int
	done   chan bool
	close  bool
	DoFunc func(map[string]string)
}

//Run 启动
func (app *App) Run() {
	app.thread, _ = conf.Int("app_threads")
	if app.thread < 1 {
		app.thread = runtime.NumCPU() //使用CPU多核处理
	}
	runtime.GOMAXPROCS(app.thread)
	app.done = make(chan bool, app.thread)
	go app.listenSignal()
	app.work()
}

//work 执行程序
func (app *App) work() {
	for i := 0; i < app.thread; i++ {
		go func(i int) {
			for {
				if app.close == true {
					app.done <- true
					break
				}
				option := make(map[string]string)
				option["requestID"] = string(utils.RandomCreateBytes(32))
				app.DoFunc(option)
				time.Sleep(1 * time.Second)
			}
		}(i)
		time.Sleep(1 * time.Second)
	}
	for i := 0; i < app.thread; i++ {
		<-app.done
	}
}

//listenSignal 监听信号
func (app *App) listenSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	for {
		s := <-sig
		if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGQUIT {
			app.close = true
			log.Printf("exit signal:%v", s)
			break
		} else {
			log.Printf("ingore signal:%v", s)
		}
	}
}
