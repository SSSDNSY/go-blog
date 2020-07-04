package test

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/rpc"
	"testing"
)

func TestRpcClient(t *testing.T) {
	//client, err := rpc.DialHTTP("tcp", "localhost:1234")
	client, err := rpc.Dial("tcp", "localhost:8899")
	if nil != err {
		logs.Error("dialing", err)
	}
	args := Args{18, 3}
	var reply int
	err = client.Call("Math.Multiply", &args, &reply)
	if err != nil {
		logs.Error("Math error", err)
	}
	fmt.Printf("\nMath:%dX%d=%d", args.A, args.B, reply)

	var q Quotient
	err = client.Call("Math.Divide", args, &q)
	if err != nil {
		logs.Error("Math error", err)
	}
	fmt.Printf("\nMath:%d/%d: quo=%d,rem=%d", args.A, args.B, q.Quo, q.Rem)
}
