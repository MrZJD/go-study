package base

func Run_opt() {
	// 基础运算符与c语言一致 不做特殊说明

	// 1. 指针运算
	var a int = 32
	var pa *int

	pa = &a // 取地址

	println("value of a: ", a)
	println("using pointer to a: ", *pa) // 寻址
}
