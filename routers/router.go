package routers

import (
	"github.com/astaxie/beego"
	"go-blog/controllers"
)

func Init() {
	//注册页面路由
	beego.Router("/", &controllers.Home{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/file", &controllers.AttachmentController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	//自动路由
	beego.AutoRouter(&controllers.TopicController{})
	//静态文件处理（附件）
	//os.Mkdir("attachment",os.ModePerm)
	//控制器处理
	beego.Router("/attachment/:all", &controllers.AttachmentController{})
	beego.SetStaticPath("/attachment", "attachment")

}
