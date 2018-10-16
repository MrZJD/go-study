package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) { // 注意这里为指针传递
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	// 1. waitGroup // 用于 程序阻塞 等待一批Go协程执行结束
	no := 3
	var wg sync.WaitGroup // 结构体
	for i := 0; i < no; i++ {
		wg.Add(1) // 计数器
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
