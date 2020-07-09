package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (c* CategoryController) Get() {
	c.Data["Title"] = "分类"
	c.Data["HomeIndex"] = 1

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if name == "" {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			logs.Error(err.Error())
			break
		}
		c.Redirect("/category", 301)
		break
	case "del":
		id, err := strconv.ParseInt(c.Input().Get("id"), 10, 64)
		if err != nil {
			logs.Error(err.Error())
			break
		}
		err = models.DeleteCategory(id)
		if err != nil {
			logs.Error(err.Error())
			break
		}
		break
	}

	categories, _ := models.GetAllCategory()
	c.Data["Categories"] = categories
	c.TplName = "category.html"
}
