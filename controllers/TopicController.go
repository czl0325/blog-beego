package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["Title"] = "文章"
	c.Data["HomeIndex"] = 2
	topics, err := models.GetAllTopic(false, "", "")
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
	labels := c.Input().Get("labels")
	cid := c.Input().Get("cid")
	id := c.Input().Get("id")

	var err error
	if id != "" {
		err = models.ModifyTopic(id, title, cid, labels, content)
	} else {
		err = models.AddTopic(id, title, cid, labels, content)
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
	topic.Comments, err = models.GetAllComment(id)
	if err != nil {
		println(err.Error())
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic
	labels := strings.Split(topic.Labels, "#")
	for index := range labels {
		labels[index] = strings.Replace(labels[index], "$", "", -1)
	}
	for  {
		num := 0
		for index := range labels {
			if labels[index] == "" {
				num = 1
				labels = append(labels[:index], labels[index+1:]...)
				break
			}
		}
		if num == 0 {
			break
		}
	}
	c.Data["Labels"] = labels
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
	labels := strings.Split(topic.Labels, "#")
	for index := range labels {
		labels[index] = strings.Replace(labels[index], "$", "", -1)
	}
	for  {
		num := 0
		for index := range labels {
			if labels[index] == "" {
				num = 1
				labels = append(labels[:index], labels[index+1:]...)
				break
			}
		}
		if num == 0 {
			break
		}
	}
	topic.Labels = strings.Join(labels, ",")
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
