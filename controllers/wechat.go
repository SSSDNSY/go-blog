package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-blog/models"
	"go-blog/util"
)

//s: App.Table.FreeQuery
//model_name: score_history
//fields: id,user_step,user_time,user_infos,add_time,update_time,ext_data
//field:
//where: [["user_infos","<>","computer"],["user_infos","<>","AI player"],["user_infos","NLIKE","未登录"]]
//perpage: 50
//app_key: A1MLDO616VDVBQHVQOKFBT6OQDJEJF280W5S
//order: ["user_step ASC", "user_time ASC"]
//sign: C63C5A88DA3113E8B97F10C293DF4662

type WeChatController struct {
	beego.Controller
}

const (
	GET_ALL = "App.Table.FreeQuery"
	ADD     = "App.Table.Create"
	//GET_ONE = "App.Table.FreeFindOne"
)

var AUTH_FLAG bool = false

func (this *WeChatController) Prepare() {
	s := this.GetString("s")
	appKey := this.GetString("app_key")
	sign := this.GetString("sign")

	var signServer string
	if s == GET_ALL {
		signServer = util.GetMD5Hash(GET_ALL + util.ApiAppSecrect)
	} else if s == ADD {
		signServer = util.GetMD5Hash(ADD + util.ApiAppSecrect)

	}

	if appKey != "" && sign != "" && appKey == util.ApiAppKey && sign == signServer {
		AUTH_FLAG = true
	}

	logs.Debug("【微信】接口安全预校验：\ns=", s, "\nappKey=", appKey, "\nsign=", sign, "\nsignServer=", signServer)
}

func (this *WeChatController) Get() {
	m := make(map[string]string)
	paramStr := this.Input().Get("param")
	m["param"] = paramStr
	logs.Info("【微信】微信GET调用:", this.Input())
	this.Data["json"] = m["param"]
	this.ServeJSON()
}

func (this *WeChatController) Post() {
	if !AUTH_FLAG {
		logs.Error("【微信】客户端非法访问!")
		this.Data["json"] = "error"
	} else { //校验通过
		paramStr := this.Input().Get("s")

		logs.Info("微信POST调用:", this.Input())

		if paramStr == GET_ALL {
			limit, err := this.GetInt("perpage")
			if err != nil {
				logs.Error("【微信】查询用户成绩失败：", err)
				this.Data["json"] = "error"
			} else {
				record, err := models.QueryScRecord(limit)
				if err != nil {
					logs.Error("【微信】查询用户成绩失败：", err)
					this.Data["json"] = "error"
				} else {
					this.Data["json"] = record
				}
			}
		} else if paramStr == ADD {
			userInfo := this.GetString("user_info")
			time, err := this.GetFloat("user_time")
			step, err := this.GetInt32("user_step")
			if err != nil {
				logs.Error("【微信】保存用户成绩失败1 ：", err)
				this.Data["json"] = "error"
			} else {
				err = models.AddScRecord(userInfo, step, time)
				if err != nil {
					logs.Error("【微信】保存用户成绩失败2 ：", err)
					this.Data["json"] = "error"
				} else {
					this.Data["json"] = "ok"
				}
			}
		} else { //待开发

		}
	}
	this.ServeJSON()
}

func (this *WeChatController) Finish() {
	if AUTH_FLAG {
		AUTH_FLAG = false
	}
}
