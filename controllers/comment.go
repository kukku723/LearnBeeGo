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
	c.TplName="main.html"
}
func (c *Comment) Post() {
	reslt:=orm.NewOrm()
	commet:=models.Comment{
		Name:c.GetString("title"),
		Message:"body",
	}
	_,err:=reslt.Insert(&commet)
	fmt.Print(err)
	c.TplName="main.html"
}