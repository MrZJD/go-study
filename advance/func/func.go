package main

import "fmt"

// 1. 匿名函数
// 2. 自执行 闭包
// 3. 自定义的函数类型
type add func(a int, b int) int

// 4. 函数作为参数 函数作为返回值
func simple(a func(value1, value2 int) int) {
	fmt.Println(a(60, 50))
}

func simple2() func(value1, value2 int) int {
	return func(a, b int) int {
		return a + b
	}
}

// 5. 函数式编程
func forEach(s []int, f func(val int) int) []int {
	var res []int
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

func main() {

	// 1
	a := func() {
		fmt.Println("hello world first class func")
	}

	a()
	fmt.Printf("%T\n", a) // func()

	// 2
	func(name string) {
		fmt.Println("--- closure ---", name)
	}("Mrzjd")

	// 3
	var b add = func(a int, b int) int {
		return a + b
	}
	s := b(5, 6)
	fmt.Printf("Type: %T, value: %v\n", b, s)

	// 4
	simple(func(a, b int) int {
		return a + b
	})

	fmt.Println(simple2()(11, 10))

	// 5
	v := forEach([]int{1, 2, 3, 4, 5}, func(val int) int {
		return val * 10
	})
	fmt.Printf("value of %v", v)
}
