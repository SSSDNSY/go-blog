package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"go-blog/models"
	"go-blog/routers"
	"time"
)

func init() {
	logs.Info("c===[]>>>>>>>>>>>>>>>>>初始化配置>>>>>>>>>>>>>>>>>>")
	models.RegisterDB()
}

func main() {

	time.Local = time.FixedZone("CST", 3600*8)
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	routers.Init()
	logs.Info("c===[]>>>>>>>>>>>>>>>>>GOBLOG服务已启动>>>>>>>>>>>>>>>>>>")
	beego.Run()

}
