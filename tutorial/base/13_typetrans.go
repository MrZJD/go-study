package base

func Run_typetrans() {
	// 强制类型转换

	var sum int = 10
	var mean float32

	println("Type Trans:", sum/3)

	mean = float32(sum)
	println("Type Trans:", mean/3)
}
