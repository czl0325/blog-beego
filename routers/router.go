package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/category", &controllers.CategoryController{})
    beego.AutoRouter(&controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
    beego.AutoRouter(&controllers.TopicController{})
    beego.Router("/comment", &controllers.CommentController{})
    beego.AutoRouter(&controllers.CommentController{})
}
