package main

import (
	_ "bbs_server/application"
	"bbs_server/application/module/quere"
	"bbs_server/application/module/quere/adapter"
)

func main() {
	app := new(queue.App)
	app.DoFunc = work
	app.Run()
}

func work(options map[string]string) {
	adapter.Run(options)
}
