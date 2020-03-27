package test

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"strings"
	"testing"
	"time"
)

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

func TestSplit(t *testing.T) {
	fmt.Println(strings.SplitN(BDID, ";", 3))
}

func TestEqt(t *testing.T) {
	fmt.Println(strings.EqualFold("123", "12"))
	fmt.Println(strings.EqualFold("123", strings.Trim("   123 ", " ")))
}
