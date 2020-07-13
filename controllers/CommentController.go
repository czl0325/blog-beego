package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) Add() {
	tid := c.Input().Get("tid")
	name := c.Input().Get("name")
	content := c.Input().Get("content")

	err := models.AddComment(tid, name, content)
	if err != nil {
		println(err)
	}
	c.Redirect("/topic/view/"+tid, 302)
}
