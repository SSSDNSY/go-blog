package controllers

import "github.com/astaxie/beego"

type WeChatController struct {
	beego.Controller
}

func (this *WeChatController) Get() {
	m := make(map[string]string)
	paramStr := this.Input().Get("param")
	m["param"] = paramStr
	beego.Info("微信调用:", paramStr)
	this.Data["json"] = m["param"]
	this.ServeJSON()
}
