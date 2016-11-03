package queue

import (
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/ccyun/GoApp/model"
)

//App 流程控制
type App struct {
	thread int
	done   chan bool
	close  bool
	DoFunc func()
}

//initRegister 初始化注册
func initRegister() {
	model.RegisterModels()
}

//Run 启动
func (app *App) Run() {
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
				app.DoFunc()
				time.Sleep(30 * time.Second)
			}
		}(i)
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
