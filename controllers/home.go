package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
)

type Home struct {
	beego.Controller
}

func (this *Home) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopic(this.Input().Get("cate"), this.Input().Get("label"), true)
	if nil != err {
		beego.Error(err.Error())
	} else {
		this.Data["Topics"] = topics
	}
	categories, err := models.GetAllCateGories()
	if nil != err {
		beego.Error(err.Error())
	}
	this.Data["Categories"] = categories
}
