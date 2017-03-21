package queue

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/astaxie/beego/utils"
	"github.com/ccyun/GoApp/application/library/hook"
	"github.com/ccyun/GoApp/application/model"
)

//App 流程控制
type App struct {
	thread int
	done   chan bool
	close  bool
	DoFunc func(map[string]string)
}

//initRegister 初始化注册
func initRegister() {
	model.RegisterModels()
}

//Run 启动
func (app *App) Run() {
	hook.AppRunStart()
	initRegister()
	if len(os.Args) > 1 {
		app.thread, _ = strconv.Atoi(os.Args[1])
	}
	if app.thread < 1 { //使用CPU多核处理
		app.thread = runtime.NumCPU()
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
				log.Println(option)
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
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	app.close = true
}
