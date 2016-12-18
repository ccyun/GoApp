package main

import (
	_ "github.com/ccyun/GoApp/application"
	"github.com/ccyun/GoApp/application/module/quere"
	"github.com/ccyun/GoApp/application/module/quere/adapter"
)

func main() {
	app := new(queue.App)
	app.DoFunc = work
	app.Run()

}

func work(options map[string]string) {
	adapter.Run(options)
}
