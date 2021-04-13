package routers

import (
	"mbook/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	//首页&分类&详情
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/2", &controllers.HomeController{}, "get:Index2")

}
