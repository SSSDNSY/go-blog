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

	//i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini");
	time.Local = time.FixedZone("CST", 3600*8)
	//timelocal = time.LoadLocation("Asia/Chongqing") //方法2
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>", time.Now().Local())
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	routers.Init()
	beego.Run()
}
