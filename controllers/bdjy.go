package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/util"
	"strings"
)

type BdjyController struct {
	beego.Controller
}

var lru map[string]string
var cacheArr []map[string]string

func (this *BdjyController) Get() {
	this.Data["subtitle"] = "百度经验统计"
	this.Data["image"] = "static/img/about-bg.jpg"
	this.Data["title"] = " jingyan"
	this.TplName = "bdjy.html"
}

func (this *BdjyController) Person() {
	uuid := this.Ctx.Request.FormValue("uuid")
	uuid = getUUID(uuid)
	html, _ := util.GetPerson(uuid)
	bd, _ := util.ParsePerson(html)
	this.Data["json"] = &bd
	this.ServeJSON()
}

func (this *BdjyController) Static() {
	uuid := this.Ctx.Request.FormValue("uuid")
	uuid = getUUID(uuid)
	json := util.ParseExPublished(uuid)
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *BdjyController) Reward() {
	this.Data["json"] = util.ParRewardExp()
	this.ServeJSON()
}

func getUUID(uuid string) string {
	if len(cacheArr) > 9999 {
		logs.Error("uuid缓存数组超过极限值 %v =%d", cacheArr, len(cacheArr))
	}
	if nil == lru && len(lru) == 0 {
		lru = make(map[string]string)
		lru["uuid"] = uuid
		return uuid
	} else {
		if len(lru) > 0 && len(lru["uuid"]) > 0 {
			if strings.EqualFold(uuid, lru["uuid"]) {
				return uuid
			} else {
				cacheArr[len(cacheArr)] = lru
				lru["uuid"] = uuid
				return uuid
			}
		} else {
			lru["uuid"] = uuid
			return uuid
		}
	}
}
