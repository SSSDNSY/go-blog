package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	var str = "W:\\e\\algorithm\\gh\\src\\main\\java\\algo"
	fileInfos, _ := ioutil.ReadDir(str)

	num := 0
	var totalNum int
	for i, info := range fileInfos {
		vName := info.Name()
		if strings.Contains(vName, "(.gitkeep") {
			fmt.Println(vName)
			num++
			os.Remove(str + "\\" + vName)
		}
		totalNum = i
	}
	fmt.Println("i=", totalNum, " num=", num)
}
