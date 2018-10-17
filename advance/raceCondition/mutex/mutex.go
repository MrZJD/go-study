package main

import (
	"fmt"
	"sync"
)

// 1. Race Condition
// 当多个协程同时操作同一块数据时，程序的输出由协程的执行顺序决定

// 2. Mutex // 加锁机制 // 确保某时刻只有一个协程在临界区运行
// mutex.Lock()
// x = x + 1
// mutex.Unlock()

var x = 0
var y = 0

func incre(wg *sync.WaitGroup) {
	x += 1
	wg.Done()
}

func increByMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y += 1
	m.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incre(&wg)
	}
	wg.Wait()
	fmt.Println("Final value of x is:", x) // value值不确定 基本都是小于1000的

	var wg2 sync.WaitGroup
	var m sync.Mutex
	for j := 0; j < 1000; j++ {
		wg2.Add(1)
		go increByMutex(&wg2, &m)
	}
	wg2.Wait()
	fmt.Println("Final value of x by mutex is:", y) // 1000
}
