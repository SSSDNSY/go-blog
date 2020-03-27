package controllers

import (
	"github.com/astaxie/beego"
	"go-blog/util"
	"log"
	"strings"
)

type BdjyController struct {
	beego.Controller
}

var lru map[string]string
var cacheArr []map[string]string

func (this *BdjyController) Get() {
	this.Data["IsBd"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["lru"] = lru
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

func getUUID(uuid string) string {
	if len(cacheArr) > 9999 {
		log.Fatal("uuid缓存数组超过极限值 %v =%d", cacheArr, len(cacheArr))
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
