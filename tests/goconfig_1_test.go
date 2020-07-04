package test

import (
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/logs"
	"log"
	"testing"
)

func TestConf(t *testing.T) {
	confFile := "../conf/conf.ini"
	cfg, err := goconfig.LoadConfigFile(confFile)
	if err != nil {
		logs.Error("无法获取配置文件", confFile, ": ", err)
	}
	//获取配置kv
	val, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	println(val)
	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "key_default", "这是新的值")
	println(isInsert)

	val, err = cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	println(val)

	//获取配置分区注释
	comment := cfg.GetSectionComments("super")
	//log.Println(comment)
	log.Println(comment)

	flag := cfg.SetSectionComments("super", "#这是新的分区注释")
	println(flag)
	//不同类型的值读取
	vInt, err := cfg.Int("must", "int")
	if err != nil {
		logs.Error("无法获取配置文件", confFile, ": ", err)
	}
	println("vInt=", vInt)

	vin1t := cfg.MustInt("must", "in1t")
	println("vin1t=", vin1t)

	goconfig.SaveConfigFile(cfg, confFile)
}
