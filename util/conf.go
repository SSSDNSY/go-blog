package util

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	Port int

	RunMode   string
	DriveName string
	DevDB     string
	ProdDB    string
	DbConn    string

	OrmDebug bool
)

func Init() {
	var err error
	Cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		logs.Error("Parse conf.ini failed: %v", err)
	}
	loadApp()
	loadServer()
}

func loadServer() {
	section, err := Cfg.GetSection("server")
	if err != nil {
		logs.Error("Parse conf.ini section [app] failed: %v", err)
	}
	//Port = 9002
	Port = section.Key("Port").MustInt(9002)
}

func loadApp() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		logs.Error("Parse conf.ini section [app] failed: %v", err)
	}
	//runMode = dev
	//devDB = "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8"
	//ProdDB = "root:root@tcp(111.229.192.247:3306)/goblog?charset=utf8"
	RunMode = section.Key("RunMode").MustString("dev")
	OrmDebug = section.Key("OrmDebug").MustBool(true)
	DriveName = section.Key("DriveName").MustString("mysql")
	DevDB = section.Key("DevDB").MustString("root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8")
	ProdDB = section.Key("ProdDB").MustString("root:root@tcp(111.229.192.247:3306)/goblog?charset=utf8")
	if RunMode == "dev" {
		DbConn = DevDB
	} else {
		DbConn = ProdDB
	}
}
