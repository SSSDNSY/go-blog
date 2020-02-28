package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCateGory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}

	this.TplName = "category.html"
	this.Data["IsCategory"] = true

	var err error
	this.Data["Categories"], err = models.GetAllCateGories()
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController) Post() {

}
