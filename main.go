package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"go-blog/models"
	"go-blog/routers"
)

func init() {
	logs.Info("c===[]>>>>>>>>>>>>>>>>>初始化配置>>>>>>>>>>>>>>>>>>")
	models.RegisterDB()
	orm.RunSyncdb("default", false, true)
}

func main() {

	routers.Init()
	logs.Info("c===[]>>>>>>>>>>>>>>>>>BLOG服务启动>>>>>>>>>>>>>>>>")
	beego.Run()

}
