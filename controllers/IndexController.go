package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Title"] = "博客"
	c.Data["HomeIndex"] = 0
	c.TplName = "index.html"
}
