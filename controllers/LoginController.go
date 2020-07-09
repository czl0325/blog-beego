package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	exit := c.Input().Get("exit") == "true"
	if exit {
		c.Ctx.SetCookie("name", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	account := c.Input().Get("account")
	password := c.Input().Get("password")
	remember := c.Input().Get("remember") == "on"
	if account == beego.AppConfig.String("account") && password == beego.AppConfig.String("password") {
		maxAge := 0
		if remember {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("name", account, maxAge, "/")
		c.Ctx.SetCookie("pwd", password, maxAge, "/")
		c.Redirect("/", 301)
	} else {
		c.Data["Error"] = "用户名密码错误"
	}
	return
}

func checkAccount(c *beego.Controller) bool {
	ck, err := c.Ctx.Request.Cookie("name")
	if err != nil {
		return false
	}
	name := ck.Value

	ck, err = c.Ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return name == beego.AppConfig.String("account") && pwd == beego.AppConfig.String("password")
}
