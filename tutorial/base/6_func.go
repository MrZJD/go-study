package base

import "math"

// 函数声明
// func funcName ([param list]) [return_type] {}

func Run_fn() {
	// 匿名函数
	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}
	println(getSqrt(4))

	// 函数支持当做值来使用 高阶函数
	// 闭包
	getSequence := func() func() int {
		i := 0
		return func() int {
			i += 1
			return i
		}
	}

	nextNum := getSequence()

	println("nextNum -> ", nextNum())
	println("nextNum -> ", nextNum())
	println("nextNum -> ", nextNum())

	// 方法
	var c1 Circle
	c1.radius = 10
	println("Area of Circle(r is 10): -> ", c1.getArea())

	// 递归调用
	println("Fib Arr of 10 is", fib(10))
}

// 方法声明
// func (variable vari_type) func_name() [return_type] {}

type Circle struct {
	radius float64
}

// 方法 -> 定义在 [结构体 或者 命名类型] 上
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

// 支持递归
func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
