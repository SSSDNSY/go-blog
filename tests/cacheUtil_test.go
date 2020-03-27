package test

import (
	"fmt"
	"go-blog/util"
	"testing"
)

func TestCache(t *testing.T) {

	c := util.GetIns()
	c.PutVal("int", 1)
	c.PutVal("string", "234")
	c.PutVal("arr", []int{1, 34, 3})

	m1 := make(map[string]string)
	m1["k1"] = "val1"
	m1["k2"] = "val2"
	c.PutVal("map", m1)
	fmt.Println(c)
	fmt.Println(c.Empty())
	fmt.Println(c.Size())
	fmt.Println(c.GetVal("map"))
	fmt.Println(c.GetVal("int"))
	fmt.Printf("类型=%T ,值=%d\n", c.GetVal("int"), c.GetVal("int"))
	fmt.Printf("类型=%T ,值=%+v\n", c.GetVal("map"), c.GetVal("map"))
	fmt.Println("test contains ", c.Contain("in1t"))
}
