package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/util"
)

type BdjyController struct {
	beego.Controller
}

var lru map[string]string

func (this *BdjyController) Get() {
	this.Data["IsBd"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["lru"] = lru
	this.TplName = "bdjy.html"
}

func (this *BdjyController) Api() {
	uuid := this.Ctx.Request.FormValue("uuid")
	html, _ := util.Get(uuid)
	bd, _ := util.ParseBDParam(html)
	if nil == lru && len(lru) == 0 {
		lru = make(map[string]string)
		lru["uuid"] = uuid
	} else {
		lru["uuid"] = uuid
	}
	this.Data["json"] = &bd
	this.ServeJSON()
}
