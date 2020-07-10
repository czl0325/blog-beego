package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (c* CategoryController) Get() {
	c.Data["Title"] = "分类"
	c.Data["HomeIndex"] = 1
	categories, _ := models.GetAllCategory()
	c.Data["Categories"] = categories
	c.TplName = "category.html"
}

func (c* CategoryController) Add() {
	name := c.Input().Get("name")
	if name == "" {
		println("缺少参数name")
	} else {
		err := models.AddCategory(name)
		if err != nil {
			println(err.Error())
		}
	}
	c.Redirect("/category", 302)
}

func (c* CategoryController) Del() {
	id, err := strconv.ParseInt(c.Input().Get("id"), 10, 64)
	if err != nil {
		println(err.Error())
		return
	}
	err = models.DeleteCategory(id)
	if err != nil {
		println(err.Error())
	}
	c.Redirect("/category", 302)
}