package main

import (
	"fmt"
	"time"
)

// 协程的优势:
// 1. 无需指定固定的堆栈大小 (可根据应用需求进行增减)
// 2. 复用os线程
// 3. channel通信防止竞态

func hello() {
	fmt.Println("Hello Goroutine!")
}

func countNumbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func countChars() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go hello()
	// 1. 启动一个协程后 调用会立即返回 不会等待协程的执行
	// 2. 如果主协程终止 则程序终止 其他协程也不会运行

	time.Sleep(1 * time.Second) // 加上这一行之后才会执行到hello()

	fmt.Println("Hello MainGo! It's Main Goroutine")

	go countNumbers()
	go countChars()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated!")
}
