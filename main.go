package main

import (
	"github.com/shelmesky/rtmonitor_dashboard/Godeps/_workspace/src/github.com/astaxie/beego"
	_ "github.com/shelmesky/rtmonitor_dashboard/routers"
)

func main() {
	beego.SetStaticPath("/static/", "static")
	beego.Run()
}
