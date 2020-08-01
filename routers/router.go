package routers

import (
	"github.com/astaxie/beego"
	"go-blog/controllers"
)

func Init() {
	//注册页面路由
	beego.Router("/", &controllers.Home{})
	beego.Router("/about", &controllers.About{})
	beego.Router("/about/getWord", &controllers.About{}, "get:GetWord")
	beego.Router("/contact", &controllers.Contact{})
	beego.Router("/login", &controllers.About{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/file", &controllers.AttachmentController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.Router("/bdjy", &controllers.BdjyController{})
	beego.Router("/wc", &controllers.WeChatController{})
	//自动路由
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.BdjyController{})
	//静态文件处理（附件）
	//os.Mkdir("attachment",os.ModePerm)
	//控制器处理
	beego.Router("/attachment/:all", &controllers.AttachmentController{})
	beego.SetStaticPath("/attachment", "attachment")

}
