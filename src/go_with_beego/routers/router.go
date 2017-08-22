package routers

import (
	"go_with_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{})
}