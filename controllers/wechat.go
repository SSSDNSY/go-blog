package controllers

import "github.com/astaxie/beego"

type WeChatController struct {
	beego.Controller
}

func (this *WeChatController) Get() {
	m := make(map[string]string)
	m["param"] = this.Input().Get("param")
	this.Data["json"] = m["param"]
	this.ServeJSON()
}
