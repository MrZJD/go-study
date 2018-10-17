package main

import (
	"fmt"
	"sync"
)

// 除mutex外 利用chan处理竞态

var x = 0

func increByChan(wg *sync.WaitGroup, ch chan bool) {
	ch <- true // 如果ch中buffer区域已经满了 则会阻塞协程 利用此处理竞态
	x += 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increByChan(&w, ch)
	}
	w.Wait()
	fmt.Println("Final value of x by buffer chan(cap=1): ", x) // 1000
}
