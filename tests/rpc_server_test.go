package test

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"testing"
)

type Args struct {
	A, B int
}
type Math int

func (m *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Quotient struct {
	Quo, Rem int
}

func (m *Math) Divide(args *Args, Q *Quotient) error {
	if args.B == 0 {
		return errors.New("divided by 0")
	}
	Q.Quo = args.A / args.B
	Q.Rem = args.A % args.B
	return nil
}

func TestRpcServe(t *testing.T) {
	math := new(Math)
	rpc.Register(math)
	//1、http rpc方式
	//rpc.HandleHTTP()
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8899")
	if err != nil {
		println(err.Error())
		os.Exit(2)
	}
	//2、tcp rpc方式
	listener, err := net.ListenTCP("tcp", tcpAddr)
	err = http.ListenAndServe(":8899", nil)
	if err != nil {
		println(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			println("conn error", err)
			continue
		}
		//rpc.ServeConn(conn)
		//3、json rpc方式
		jsonrpc.ServeConn(conn)
	}
	if nil != err {
		println(err.Error())
	}

}
