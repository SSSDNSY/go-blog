package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-blog/models"
	"go-blog/routers"
	"time"
)

func init() {
	models.RegisterDB()
}

func main() {

	time.Local = time.FixedZone("CST", 3600*8)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>博客服务启动", time.Now().Local())
	orm.Debug = true
	//beego.EnableAdmin = true
	//beego.AdminHttpAddr = "localhost"
	//beego.AdminHttpPort = 8088
	orm.RunSyncdb("default", false, true)
	routers.Init()
	beego.Run()

}
