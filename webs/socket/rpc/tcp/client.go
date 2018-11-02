package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8888") // just this line moded
	if err != nil {
		log.Fatal("Dialing: ", err)
	}

	// sync
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Println("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Println("Arith: %d/%d=%d reminder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
