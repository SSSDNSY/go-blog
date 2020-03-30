package test

import (
	"fmt"
	"go-blog/util"
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
	t1 :=time.Now().Local()
	fmt.Println(t1)
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())

	fmt.Println("格式化日期"+util.PaserDateTime(t1))
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
