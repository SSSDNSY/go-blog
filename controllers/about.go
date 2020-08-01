package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/util"
	"io/ioutil"
	"net/http"
)

type About struct {
	beego.Controller
}

func (this *About) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["title"] = getUrl(util.Api4)
	this.Data["image"] = "static/img/home-bg.jpg"
	this.Data["about.go"] = " less is more "
	this.TplName = "about.html"
}

func (this *About) GetWord() {
	this.Data["json"] = getUrl(util.Api4)
	this.ServeJSON()
}

func getUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("About getWord err : ", err)
	}
	word := string(body)
	myMap := make(map[string]string)
	json.Unmarshal([]byte(word), &myMap)
	return myMap["hitokoto"]
}
