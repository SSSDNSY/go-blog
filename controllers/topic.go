package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
	"path"
	"strings"
	"time"
)

type TopicController struct {
	beego.Controller
}

const TIME_LAYOUT = "2006-01-02 15:04:05"

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	topics, err := models.GetAllTopic("", "", false)
	if nil != err {
		beego.Error(err.Error())
	} else {
		for t, _ := range topics {
			topics[t].ReplyTime, _ = time.Parse(TIME_LAYOUT, topics[t].ReplyTime.Format(TIME_LAYOUT))
			topics[t].Created, _ = time.Parse(TIME_LAYOUT, topics[t].ReplyTime.Format(TIME_LAYOUT))
			topics[t].Updated, _ = time.Parse(TIME_LAYOUT, topics[t].ReplyTime.Format(TIME_LAYOUT))
		}
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.TplName = "topic_view.html"
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if nil != err {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(strings.Trim(topic.Labels, " "), " ")
	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Replies"] = replies
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}

func (this *TopicController) Edit() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.TplName = "topic_edit.html"
	tid := this.Input().Get("tid")

	topic, err := models.GetTopic(tid)
	if nil != err {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	err := models.DelTopic(this.Ctx.Input.Param("0"))
	if nil != err {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return
}

//添加 修改提交到这个post方法里面
func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	//解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	label := this.Input().Get("label")
	//获取附件
	_, fileHeader, err := this.GetFile("attachment")
	if nil != err {
		beego.Error(err.Error())
	}
	var attachmentFileName string
	if fileHeader != nil {
		//保存附件
		attachmentFileName = fileHeader.Filename
		beego.Info("》》》》》 上传附件" + attachmentFileName)
		err = this.SaveToFile("attachment", path.Join("attachment", attachmentFileName))
		if nil != err {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, content, category, label, attachmentFileName)
	} else {
		err = models.EditTopic(tid, title, content, category, label, attachmentFileName)
	}
	if nil != err {
		beego.Error(err)
	}
	this.Redirect("/topic", 301)
}
