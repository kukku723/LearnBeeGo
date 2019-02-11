package main

import (
	_ "LearnBeeGo/routers"
	_ "LearnBeeGo/models"
	"github.com/astaxie/beego"


)

func main() {
	beego.Run()
}
