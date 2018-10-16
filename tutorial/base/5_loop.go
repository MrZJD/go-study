package base

func Run_loop() {
	var a int

	for a = 0; a < 2; a++ {
		println("a -> ", a)
	}

	var b int = 4
	for a < b {
		println("a -> ", a)
		a++
	}

	var nums = [6]int{1, 2, 3, 4}
	for ind, val := range nums {
		println("range nums -> ", ind, val)
	}
}
