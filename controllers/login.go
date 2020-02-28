package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	isExist := this.Input().Get("exit") == "true"
	if isExist {
		this.Ctx.SetCookie("account", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	account := this.Input().Get("account")
	password := this.Input().Get("password")
	autoLogin := this.Input().Get("autoLogin") == "on"
	if account == beego.AppConfig.String("adminU") && password == beego.AppConfig.String("adminP") {
		maxAge := 0
		if autoLogin == true {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("account", account, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}
	account := ck.Value
	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value

	return account == beego.AppConfig.String("adminU") && password == beego.AppConfig.String("adminP")
}
