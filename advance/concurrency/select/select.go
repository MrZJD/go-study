package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successfully"
}

func main() {
	// 1. select 语句用于在多个发送/接收channel操作中进行选择

	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)

	select { // 实际场景: 快速响应先返回的数据
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2) // 执行后跳出阻塞
	}

	// 2. select default
	ch3 := make(chan string)
	go process(ch3)

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch3:
			fmt.Println("Received value:", v)
			return
		default: // 没有信道准备就绪时调用
			fmt.Println("No Value Received")
		}
	}

	// 3. deadlock // select没有default 且ch为nil或者没有写入操作时 -> 通过default可以避免

	// 4. random 同时有多个channel准备就绪时 随机选取执行

	// 5. 空select
	// select {} // -> panic deadlock 会一直阻塞
}
