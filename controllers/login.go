package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"go-blog/controllers/e"
	"go-blog/util"
)

type LoginController struct {
	beego.Controller
}
type auth struct {
	Account  string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	account := this.GetString("account")
	password := this.GetString("password")

	valid := validation.Validation{}
	a := auth{Account: account, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := account == beego.AppConfig.String("adminU") && password == beego.AppConfig.String("adminP")
		if isExist {
			token, err := util.GenerateToken(account, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logs.Info(err.Key, err.Message)
		}
	}
	this.Data["json"] = e.SetResult(code, e.GetMsg(code), data)
	this.ServeJSON()

}
