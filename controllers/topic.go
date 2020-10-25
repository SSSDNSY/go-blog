package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/controllers/e"
	"go-blog/models"
	"path"
	"strconv"
	"strings"
	"time"
)

type TopicController struct {
	beego.Controller
}

const TIME_LAYOUT = "2006-01-02 15:04:05"

func (this *TopicController) Blog() {
	id := this.Ctx.Input.Params()["0"]
	this.TplName = "topic.html"
	_, err := strconv.Atoi(id)
	if err == nil {
		topic, err1 := models.GetTopic(id)
		if nil != err1 {
			logs.Error(err1.Error())
		} else {
			topic.ReplyTime, _ = time.Parse(TIME_LAYOUT, topic.ReplyTime.Format(TIME_LAYOUT))
			topic.Created, _ = time.Parse(TIME_LAYOUT, topic.ReplyTime.Format(TIME_LAYOUT))
			topic.Updated, _ = time.Parse(TIME_LAYOUT, topic.ReplyTime.Format(TIME_LAYOUT))
			this.Data["Topic"] = topic
		}
	}
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {

	this.TplName = "topic_view.html"
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if nil != err {
		logs.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(strings.Trim(topic.Labels, " "), " ")
	replies, err := models.GetAllReplies(this.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
		return
	}
	this.Data["Replies"] = replies
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}

func (this *TopicController) Edit() {
	tid := this.Ctx.Input.Params()["0"]
	this.TplName = "topicEdit.html"

	topic, err := models.GetTopic(tid)
	if nil != err {
		logs.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["blogNum"], _ = models.GetBlogCount()
	this.Data["cateNum"], _ = models.GetCateCount()
	this.Data["Tid"] = tid
	this.Data["Categories"], _ = models.GetAllCateGories()

}

func (this *TopicController) Delete() {
	err := models.DelTopic(this.Ctx.Input.Param("0"))
	if nil != err {
		logs.Error(err)
	}
	this.Redirect("/topic", 302)
	return
}

//添加 修改提交到这个post方法里面
func (this *TopicController) Post() {
	//if !checkAccount(this.Ctx) {
	//	this.Redirect("/login", 302)
	//	return
	//}
	//解析表单
	tid := this.Input().Get("tid")
	title := this.GetString("title")
	content := this.GetString("content")
	category := this.GetString("category")
	label := this.Input().Get("label")
	//参数校验
	code := e.SUCCESS
	if title == "" || category == "" || content == "" {
		code = e.INVALID_PARAMS
	}
	//获取附件
	_, fileHeader, err := this.GetFile("attachment")
	if nil != err {
		logs.Error(err.Error())
	}
	var attachmentFileName string
	if fileHeader != nil {
		//保存附件
		attachmentFileName = fileHeader.Filename
		logs.Info("》》》》》 上传附件" + attachmentFileName)
		err = this.SaveToFile("attachment", path.Join("attachment", attachmentFileName))
		if nil != err {
			logs.Error(err)
		}
	}

	if len(tid) == 0 {
		id, err := models.AddTopic(title, content, category, label, attachmentFileName)
		if nil != err {
			logs.Error(err)
		} else {
			this.Data["json"] = e.SetResult(code, e.GetMsg(code), id)
			this.ServeJSON()
			return
		}
	} else {
		err = models.EditTopic(tid, title, content, category, label, attachmentFileName)
	}
	if nil != err {
		logs.Error(err)
	}
	this.Data["json"] = e.SetResult(code, e.GetMsg(code), nil)
	this.ServeJSON()
}
