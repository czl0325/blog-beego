package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Title"] = "博客"
	c.Data["HomeIndex"] = 0
	topics, err :=  models.GetAllTopic(true, c.Input().Get("c"), c.Input().Get("l"))
	if err != nil {
		println(err.Error())
	}
	categories, err := models.GetAllCategory()
	if err != nil {
		println(err.Error())
	}
	c.Data["Categories"] = categories
	c.Data["Topics"] = topics
	c.TplName = "index.html"
}
