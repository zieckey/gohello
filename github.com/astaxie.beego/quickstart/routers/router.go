package routers

import (
	"github.com/zieckey/gohello/beego/quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
