package base

func Run_pointer() {
	var ptr *int // 默认为空指针

	if ptr == nil {
		println("ptr is nil")
	} else {
		println("ptr is not nil")
	}

	var a int = 20

	ptr = &a

	println("a is ->", a)
	println("address of a is ->", ptr)
	println("address of a is pointing ->", *ptr)

	// 指针数组
	var ptrArr [3]*int // 3个指向int的指针
	var i int
	for i = 0; i < 3; i++ {
		ptrArr[i] = &i
	}
	for i = 0; i < 3; i++ {
		println("Using pointer array -> ", i, *ptrArr[i])
	}

	// 指针的指针
	var ptrOfPtr **int

	ptrOfPtr = &ptr

	println("Using double pointer to point a ->", **ptrOfPtr)

	// 指针的传递
	var swap = func(x *int, y *int) {
		// var temp int
		// temp = *x
		// *x = *y
		// *y = temp

		*x, *y = *y, *x
	}

	var t1 int = 10
	var t2 int = 20

	println("before swaping value by pointer: ", t1, t2)

	swap(&t1, &t2)

	println("after swaping value by pointer: ", t1, t2)
}
