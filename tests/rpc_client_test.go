package test

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func TestRpcClient(t *testing.T) {
	//client, err := rpc.DialHTTP("tcp", "localhost:1234")
	client, err := rpc.Dial("tcp", "localhost:8899")
	if nil != err {
		log.Fatal("dialing", err)
	}
	args := Args{18, 3}
	var reply int
	err = client.Call("Math.Multiply", &args, &reply)
	if err != nil {
		log.Fatal("Math error", err)
	}
	fmt.Printf("\nMath:%dX%d=%d", args.A, args.B, reply)

	var q Quotient
	err = client.Call("Math.Divide", args, &q)
	if err != nil {
		log.Fatal("Math error", err)
	}
	fmt.Printf("\nMath:%d/%d: quo=%d,rem=%d", args.A, args.B, q.Quo, q.Rem)
}
