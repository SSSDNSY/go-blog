package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/models"
	"go-blog/util"
	"strconv"
)

type Home struct {
	beego.Controller
}

func (this *Home) Get() {
	this.Data["IsHome"] = true
	this.Data["title"] = "Pengzh"
	this.Data["image"] = "https://open.saintic.com/api/bingPic/"
	this.Data["subtitle"] = " less is more "
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.Data["leftDis"] = "none"
	this.Data["RightDis"] = ""

	var topics []*models.Topic
	var err error
	var pageNumb int64

	lpage := this.Input().Get("lPageNumber")
	rpage := this.Input().Get("rPageNumber")
	if len(rpage) > 0 {
		pageNumb, _ = strconv.ParseInt(rpage, 10, 64)
		this.Data["lPageNumber"] = pageNumb - 1
		this.Data["rPageNumber"] = pageNumb + 1
		topics, err = models.GetPageTopic(pageNumb, util.PageSize)
		if len(topics) < 1 && nil == err {
			this.Data["leftDis"] = ""
			this.Data["rightDis"] = "none"
		} else {
			this.Data["leftDis"] = ""
			this.Data["rightDis"] = ""
		}
	} else if len(lpage) > 0 {
		pageNumb, _ = strconv.ParseInt(lpage, 10, 64)
		this.Data["lPageNumber"] = pageNumb - 1
		this.Data["PageNumber"] = pageNumb + 1
		if pageNumb <= 1 {
			this.Data["leftDis"] = "none"
			this.Data["rightDis"] = ""
		} else {
			this.Data["leftDis"] = ""
			this.Data["rightDis"] = ""
		}
		topics, err = models.GetPageTopic(pageNumb, util.PageSize)
	} else {
		topics, err = models.GetPageTopic(1, util.PageSize)
		this.Data["lPageNumber"] = 1
		this.Data["rPageNumber"] = 2
	}

	if nil != err {
		logs.Error(err.Error())
	} else {
		this.Data["Topics"] = topics
	}
	categories, err := models.GetAllCateGories()
	if nil != err {
		logs.Error(err.Error())
	}
	this.Data["Categories"] = categories
}
