package routers

import (
	"github.com/shelmesky/rtmonitor_dashboard/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
