package routers

import (
	"LearnBeeGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/comment", &controllers.Comment{})
}