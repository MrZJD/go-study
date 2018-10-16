package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully Write", i, "to Buffer Channal")
	}
	close(ch)
}

func main() {
	// 缓冲信道: 相对于普通信道 只有在缓冲空间填满的时候阻塞信道

	// ch0 := make(chan string, 0)
	// ch0 <- "none" // 死锁

	// 1. capacity 表示 缓冲空间大小
	// ch := make(chan type, capacity)
	ch := make(chan string, 2)
	ch <- "bill"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 2. buffer channel 阻塞逻辑
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second) // write 0 1
	for chValue := range ch2 {
		fmt.Println("Read Value from ch2: ", chValue) // read 0 -> write 2 -> read 1 -> write 3 -> read 2 -> write 4 -> read 3 -> read 4
		time.Sleep(2 * time.Second)
	}

	// 3. buffer channel deadlock
	// 当写入数据超出capacity 发生阻塞 没有操作读取数据时

	// 4. len & cap
	// len(ch2) // 当前排队的元素个数
	// cap(ch2) // 信道可以存储的值的数量
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 4
	fmt.Println("Curr Channel 3: len:", len(ch3), "capacity:", cap(ch3))
}
