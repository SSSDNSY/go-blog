package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Get() {
	this.TplName = "topic.html"
	topics, err := models.GetAllTopic("", "", false)
	if nil != err {
		beego.Error(err.Error())
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *ReplyController) Add() {

	tid := this.Input().Get("tid")
	err := models.AddReply(tid, this.Input().Get("nickName"),
		this.Input().Get("content"))
	if nil != err {
		beego.Error(err.Error())
	}
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {

	tid := this.Input().Get("tid")
	err := models.DelReply(this.Input().Get("rid"))
	if nil != err {
		beego.Error(err.Error())
	}
	this.Redirect("/topic/view/"+tid, 302)
}
