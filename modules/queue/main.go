package main

import "github.com/ccyun/daemon/example/app/common"

func main() {

	app := new(common.App)
	app.DoFunc = work
	app.Run()

}
func work() {
	common.Run()
}
