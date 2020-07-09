package controllers

import "github.com/astaxie/beego"

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["Title"] = "文章"
	c.Data["HomeIndex"] = 2
	c.TplName = "topic.html"
}
