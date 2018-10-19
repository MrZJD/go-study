package main

import (
	"fmt"
	"runtime/debug"
)

/* panic & recover */

// 当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
// 这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪（Stack Trace），最后程序终止.

// 1. 尽可能是有error 只有当程序无法运行时再使用panic和recover
// 1.1 -> 发生了一个不能恢复的错误 程序不能继续运行
// 1.2 -> 发生了一个编程上的错误
// 2. defer 堆栈依然会执行

// 3. recover 是一个内建函数，用于重新获得 panic 协程的控制
// 只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，并且停止 panic 续发事件

// 4. recover 只有在相同的协程中才管用

// 5. 运行时panic (如数组越界)

// func panic(interface {})

// func recover() interface {}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recover from", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("defered call in fullName")
	defer recoverName()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}

	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)

		// 5.1 获取堆栈
		debug.PrintStack()
	}
}

func a() {
	defer r()
	n := []int{5, 4, 7}
	fmt.Println(n[3]) // runtime panic
	fmt.Println("normally returned from a")
}

func main() {
	// defer fmt.Println("deferred call in main")

	// firstName := "mrzjd"
	// fullName(&firstName, nil)
	// fmt.Println("returned from main")

	a()
	fmt.Println("normally returned from main")
}
