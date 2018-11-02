package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

// RPC in go:
// 标准包中的RPC只支持GO客户 服务 之间的交互 (因为内部使用Gob编码)

// RPC函数格式
// func (t *T) MethodName(argType T1, replyType *T2) error

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	arith := new(Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
