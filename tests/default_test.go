package test

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go-blog/controllers"
	"go-blog/util"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGolang(t *testing.T) {
	fmt.Println(strings.Split("abc", "a")[1])
	fmt.Println(strings.Split("2019-12-21 10:28:16.1517132 +0000 UTC", ".")[0])
}

func TestStrings(t *testing.T) {
	fmt.Println(strings.Split("2019-12-19 17:31:47.6821378 +0000 UTC", ".")[0])
}

func TestTime(t *testing.T) {
	t1 := time.Now().Local()
	fmt.Println(t1)
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())

	fmt.Println("格式化日期" + util.PaserDateTime(t1))
}

func TestSefFunc(t *testing.T) {
	(func() {
		Ca := make(map[string]interface{})
		fmt.Println(Ca)
	})()
}

func TestSplit(t *testing.T) {
	fmt.Println(strings.SplitN(BDID, ";", 3))
}

func TestEqt(t *testing.T) {
	fmt.Println(strings.EqualFold("123", "12"))
	fmt.Println(strings.EqualFold("123", strings.Trim("   123 ", " ")))
}
func TestNetHttp(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		logs.Error(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func TestGetMD5Hash(t *testing.T) {

	logs.Info("sign", util.GetMD5Hash(controllers.GET_ALL+util.ApiAppSecrect))
	logs.Info("sign", util.GetMD5Hash(controllers.ADD+util.ApiAppSecrect))
	//logs.Info("sign",util.GetMD5Hash(controllers.GET_ONE+util.ApiAppSecrect))
	//logs.Info(util.RandomStr(88)) key
	//logs.Info(util.RandomStr(36)) secret
}
func TestWord(t *testing.T) {
	resp, err := http.Get("https://pyq.shadiao.app/api.php")
	if err != nil {
		logs.Error(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
func TestPagination(t *testing.T) {
	//topics, err := models.GetPageTopic(1, 2, "", "", true)
	//if err != nil {
	//	logs.Info(topics)
	//}
}
