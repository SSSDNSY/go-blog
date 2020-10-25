package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/controllers/e"
	"go-blog/models"
)

type CategoryController struct {
	beego.Controller
}

//查
func (this *CategoryController) Get() {
	this.Data["json"], _ = models.GetAllCateGories()
	this.ServeJSON()
}

//增、改
func (this *CategoryController) Post() {
	op := this.GetString("op")
	cid := this.GetString("cid")
	title := this.GetString("title")
	switch op {
	case "add":
		if len(title) == 0 {
			this.Data["json"] = e.SetResult(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), nil)
			this.ServeJSON()
		}
		err := models.AddCateGory(title)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = e.SetResult(e.ERROR, e.GetMsg(e.ERROR), nil)
		}
		this.Data["json"] = e.SetResult(e.SUCCESS, e.GetMsg(e.SUCCESS), nil)
		this.ServeJSON()
	case "del":
		if len(cid) == 0 {
			this.Data["json"] = e.SetResult(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), nil)
			this.ServeJSON()
		}
		err := models.DelCategory(cid)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = e.SetResult(e.ERROR, e.GetMsg(e.ERROR), nil)
		}
		this.Data["json"] = e.SetResult(e.SUCCESS, e.GetMsg(e.SUCCESS), nil)
		this.ServeJSON()
	case "upd":
		if len(cid) == 0 {
			this.Data["json"] = e.SetResult(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), nil)
			this.ServeJSON()
		}
		if len(title) == 0 {
			this.Data["json"] = e.SetResult(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), nil)
			this.ServeJSON()
		}
		err := models.UpdCategory(cid, title)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = e.SetResult(e.ERROR, e.GetMsg(e.ERROR), nil)
		}
		this.Data["json"] = e.SetResult(e.SUCCESS, e.GetMsg(e.SUCCESS), nil)
		this.ServeJSON()
	}
}
