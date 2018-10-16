package main

import (
	"fmt"
)

func find(num int, nums ...int) { // -> nums := []int{} // 将多参数转为切片
	fmt.Printf("typeof nums is %T\n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in", nums)
	}

	fmt.Printf("\n")
}

func changeSlice(s ...string) {
	s[0] = "Go"
}

func changeSlice2(s ...string) {
	s[0] = "GoWow"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	find(1, 0, 1, 2, 3, 4, 5)

	find(-1, 0, 1, 2, 3, 4, 5)

	find(123)

	params := []int{0, 1, 10, 20, 30}
	find(10, params...) // ... 表示将切片直接作为参数 传入 不再进行多参数转换的语法糖

	str := []string{"World", "Hello"}
	fmt.Println(str) // World Hello
	changeSlice(str...)
	fmt.Println(str) // Go Hello

	changeSlice2(str...)
	fmt.Println(str) // GoWow Hello
}
