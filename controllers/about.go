package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/util"
	"io/ioutil"
	"net/http"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Get() {
	this.Data["title"] = "AboutController me"
	this.Data["image"] = "static/img/home-bg.jpg"
	this.TplName = "about.html"
}

func (this *AboutController) GetWord() {
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
		logs.Error("AboutController getWord err : ", err)
	}
	word := string(body)
	myMap := make(map[string]string)
	json.Unmarshal([]byte(word), &myMap)
	return myMap["hitokoto"]
}
