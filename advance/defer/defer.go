package main

import (
	"fmt"
	"sync"
)

// defer: 含有defer语句的函数 会在该函数将要返回之前 调用另一个函数

func beforeReturned() {
	fmt.Println("[4. before returned]")
}

func calc() {
	defer beforeReturned() // 1. 延迟函数 方法也是可以的
	fmt.Println("[1. start calc]")
	fmt.Println("[2. calc value -> tttt]")
	fmt.Println("[3. finish calc]")
}

func printA(a int) {
	fmt.Println("value of a in defered func", a) // 10
}

// 4. 实际应用
type rect struct {
	width int
	len   int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done() // 4. wg.Done 与 defer 联合使用
	if r.len < 0 {
		fmt.Printf("rect %v len should be greater than zero\n", r)
		// wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v width should be greater than zero\n", r)
		// wg.Done()
		return
	}
	area := r.width * r.len
	fmt.Printf("rect %v's area is %d\n", r, area)
	// wg.Done()
}

func main() {
	calc()

	a := 10
	defer printA(a) // 2. 实参取值
	a = 20
	fmt.Println("value of a before defered func call", a) // 20

	// 3. defer 为堆栈调用 LIFO: Last in First out

	// 4. 实际应用
	var wg sync.WaitGroup
	r1 := rect{-10, 20}
	r2 := rect{10, -10}
	r3 := rect{4, 5}
	rects := []rect{r1, r2, r3}

	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished")

}
