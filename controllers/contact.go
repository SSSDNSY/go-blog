package controllers

import (
	"github.com/astaxie/beego"
)

type Contact struct {
	beego.Controller
}

func (this *Contact) Get() {
	this.Data["title"] = "Contact me"
	this.Data["image"] = "static/img/contact-bg.jpg"
	this.Data["about.go"] = " Have questions? I have answers. "
	this.TplName = "contact.html"
}
