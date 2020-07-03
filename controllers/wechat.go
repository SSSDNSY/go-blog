package controllers

import "github.com/astaxie/beego"

type WeChatControllers struct {
	beego.Controller
}

func (this *WeChatControllers) Get() {
	m := make(map[string]string)
	m["param"] = this.Input().Get("param")
	this.Data["param"] = m["param"]
}
