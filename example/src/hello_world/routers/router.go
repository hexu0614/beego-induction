package routers

import (
	"hello_world/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello_world", &controllers.HWController{})
}
