package controllers

import (
	"LearnBeeGo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	beego.Controller
}

func (c *Comment) Get() {
	reslt:=orm.NewOrm()
	commet:=models.Comment{
		Name:"test",
		Message:"testMessage",
	}
	_,err:=reslt.Insert(&commet)
	fmt.Print(err)
}