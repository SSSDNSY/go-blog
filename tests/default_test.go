package test

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"strings"
	"testing"
	"time"
)

//
//import (
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"runtime"
//	"path/filepath"
//	_ "app1/routers"
//
//	"github.com/astaxie/beego"
//)
//
//func init() {
//	_, file, _, _ := runtime.Caller(0)
//	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
//	beego.TestBeegoInit(apppath)
//}
//
//
//// TestBeego is a sample to run an endpoint test
//func TestBeego(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
//	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
//
//	Convey("Subject: Test Station Endpoint\n", t, func() {
//	        Convey("Status Code Should Be 200", func() {
//	                So(w.Code, ShouldEqual, 200)
//	        })
//	        Convey("The Result Should Not Be Empty", func() {
//	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
//	        })
//	})
//}
//
func TestGolang(t *testing.T) {
	log.Error(strings.Split("abc", "a")[1])
	log.Error(strings.Split("2019-12-21 10:28:16.1517132 +0000 UTC", ".")[0])
}

func TestStrings(t *testing.T) {
	fmt.Println(strings.Split("2019-12-19 17:31:47.6821378 +0000 UTC", ".")[0])
}

func TestTime(t *testing.T) {
	time.Now().Local()
}

func TestSefFunc(t *testing.T) {
	(func() {
		Ca := make(map[string]interface{})
		fmt.Println(Ca)
	})()
}
