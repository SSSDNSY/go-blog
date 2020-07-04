package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type WeChatController struct {
	beego.Controller
}

func (this *WeChatController) Prepare() {
	beego.Debug("接口过滤:", time.Now())
}

func (this *WeChatController) Get() {
	m := make(map[string]string)
	paramStr := this.Input().Get("param")
	m["param"] = paramStr
	beego.Info("微信调用:", paramStr)
	this.Data["json"] = m["param"]
	this.ServeJSON()
}

func (c *WeChatController) Finish() {

}
