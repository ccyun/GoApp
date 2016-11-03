package main

import (
	"github.com/ccyun/GoApp/modules/queue"
	"github.com/ccyun/GoApp/modules/queue/task"
)

func main() {

	app := new(queue.App)
	app.DoFunc = work
	app.Run()

}
func work() {
	task.Run()
}
