package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["Title"] = "文章"
	c.Data["HomeIndex"] = 2
	topics, err :=  models.GetAllTopic(false, "")
	if err != nil {
		println(err.Error())
	}
	c.Data["Topics"] = topics
	c.TplName = "topic.html"
}

func (c *TopicController) Post() {
	name := c.Ctx.GetCookie("name")
	if name == "" {
		c.Redirect("/login", 302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")
	cid := c.Input().Get("cid")
	id := c.Input().Get("id")

	var err error
	if id != "" {
		err = models.ModifyTopic(id, title, cid, content)
	} else {
		err = models.AddTopic(id, title, cid, content)
	}
	if err != nil {
		println(err.Error())
	}
	c.Redirect("/topic", 302)
}

func (c *TopicController) Add() {
	categories, err := models.GetAllCategory()
	if err != nil {
		println(err.Error())
	}
	c.Data["HomeIndex"] = 2
	c.Data["Categories"] = categories
	c.TplName = "topic_add.html"
}

func (c *TopicController) View() {
	c.Data["HomeIndex"] = 2
	id := c.Ctx.Input.Param("0")
	topic, err := models.GetTopicById(id)
	if err != nil {
		println(err.Error())
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Id"] = id
	c.TplName = "topic_view.html"
}

func (c *TopicController) Modify() {
	c.Data["HomeIndex"] = 2
	id := c.Ctx.Input.Param("0")
	topic, err := models.GetTopicById(id)
	if err != nil {
		println(err.Error())
		c.Redirect("/", 302)
		return
	}
	categories, err := models.GetAllCategory()
	if err != nil {
		println(err)
	}
	c.Data["Categories"] = categories
	c.Data["Topic"] = topic
	c.TplName = "topic_modify.html"
}

func (c *TopicController) Delete() {
	c.Data["HomeIndex"] = 2
	id := c.Ctx.Input.Param("0")
	err := models.DeleteTopic(id)
	if err != nil {
		println(err.Error())
	}
	c.Redirect("/topic", 302)
}