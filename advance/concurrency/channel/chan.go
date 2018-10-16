package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello Channel in Goroutine[hello]")
	time.Sleep(1 * time.Second)
	done <- true
	fmt.Println("*** Test ***")
}

func calcS(num int, sres chan int) { // 计算每一位数的平方和
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	sres <- sum
}

func calcC(num int, cres chan int) { // 计算每一位数的立方和
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cres <- sum
}

func sendData(sendch chan<- int) { // 5. 单向可写信道
	sendch <- 10
}

func producer(chnl chan int) { // 6. 关闭信道 以及使用 range循环获取信道数据
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	var a chan int // 1. 信道关联了一个类型 并且只能运输这个类型的数据 // chan T
	if a == nil {
		fmt.Println("chan a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a) // chan int
	}

	// 2. 信道的发送和接收
	// 信道的发送和接收过程是阻塞的
	// data := <- a // 接收 读取
	// a <- data // 发送 写入

	done := make(chan bool)
	go hello(done)
	<-done // 读取数据会阻塞该协程 直到channal中有数据写入

	// time.Sleep(10)
	fmt.Println("We are done in main")

	// 3. channel数据传输示例
	num := 123
	sres := make(chan int)
	cres := make(chan int)
	go calcS(num, sres)
	go calcC(num, cres)

	ss, cc := <-sres, <-cres
	fmt.Println("Final Output: ", ss, cc)

	// 4. 死锁
	// chanA <- data // 向信道写入数据 没有协程读取时
	// <- chanB // 从信道读入数据 没有写成写入时
	// dChan := make(chan int)
	// dChan <- 5
	// <-dChan

	// 5. 单向信道
	// 创建一个单向信道没有意义，通常用于约束协程的信道通信方式
	schan := make(chan int)
	go sendData(schan)
	fmt.Println(<-schan)

	// 6. 关闭信道 以及使用 range循环获取信道数据
	schl := make(chan int)
	go producer(schl)
	// for {
	// 	v, ok := <-schl
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println("Recivied data", v, ok)
	// }
	for v := range schl {
		fmt.Println("Reviced data by range", v)
	}
}
