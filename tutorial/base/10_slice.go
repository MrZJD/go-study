package base

import "fmt"

// 切片 (动态数组)

func Run_slice() {
	// 1. 声明
	// var var_name []type
	// var slice_1 []type = make([]type, len)
	// slice_1 := make([]type, len)
	// make([]T, length, capacity) // length数组长度 切片初始长度, capacity切片容量

	s := []int{1, 2, 3} // cap=len=3
	println("length s: ", len(s), "capacity s: ", cap(s))

	var arr = [10]int{}
	var startI int = 1
	var endI int = 3

	s_1 := arr[:]           // 切片引用数组
	s_2 := arr[startI:endI] // startI -> endI-1
	s_3 := arr[startI:]     // startI -> len-1
	s_4 := arr[:endI]       // 0 -> endI-1

	println("length s_1: ", len(s_1), "capacity s_1: ", cap(s_1))
	println("length s_2: ", len(s_2), "capacity s_2: ", cap(s_2))
	println("length s_3: ", len(s_3), "capacity s_3: ", cap(s_3))
	println("length s_4: ", len(s_4), "capacity s_4: ", cap(s_4))

	// var nilSlice []int // nil 空切片

	// append()
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 1, 2, 3, 4)
	printSlice(numbers)

	var numbersBig = make([]int, 10, 10)
	copy(numbersBig, numbers)
	printSlice(numbersBig)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
