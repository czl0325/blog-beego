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
	topics, err :=  models.GetAllTopic(true)
	if err != nil {
		println(err.Error())
	}
	c.Data["Topics"] = topics
	c.TplName = "index.html"
}
