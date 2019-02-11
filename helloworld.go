package main

import (
	_ "LearnBeeGo/routers"
	_ "LearnBeeGo/models"
	"github.com/astaxie/beego"


)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
