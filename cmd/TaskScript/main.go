package main

import (
	"github.com/ccyun/GoApp/module/queue"
	"github.com/ccyun/GoApp/module/queue/mode"
)

func main() {

	app := new(queue.App)
	app.DoFunc = work
	app.Run()

}
func work() {
	mode.Run()
}
