package base

func Run_arr() {
	// var arr_name [SIZE] var_type

	var balance [10]float32

	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0} // {}中个数不能超过SIZE

	var balance3 = [...]float32{1.1, 2.2, 3.3} // [...]表示根据后面{}的个数进行判断

	println("Array 1 ->", len(balance))
	println("Array 2 ->", len(balance2))
	println("Array 3 ->", len(balance3))

	// 多维数组
	// var arr_name [SIZE][SIZE][SIZE] var_type

	var a = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	println("Multiply Di Arr: ->", len(a))

	// 数组的传参
	var tArr = func(arr []int) {
		println("func Arr without spec length ->", len(arr))
	}

	var tArr2 = func(arr [2]int) {
		println("func Arr with spec length ->", len(arr))
	}

	tArr([]int{1, 2, 3})
	tArr2([2]int{1, 2})
}
