/* 变量与数据结构 */
package base

import "unsafe"

// 1. variable statement

// var v_name v_type // i. 声明变量及类型 (不赋值 则填充默认值)
// var v_name = value // ii. 不声明类型 根据赋值自动判断
// v_name := value // iii. 省略var (只能在函数体中声明)

// 同理多个 变量声明规则同上
// var (
// 	  v_name_1 v_type_1
// 	  v_name_2 v_type_2
// ) // 多个全局变量声明

func use_var() {
	var a int = 10 // 局部变量 不允许声明了 不使用 (全局变量可以)
	var b = 10
	c := 10

	println(a, b, c)

	var a1 string = "hello go!"
	var b1 = "i'm mrzjd"
	var c1 bool

	println(a1, b1, c1)
}

func use_const() {
	// const c_name c_type = value
	// const c_name = value
	// 常量类型可以是 bool 数字型 字符串型

	const LENGTH int = 10
	const WIDTH int = 20

	println(LENGTH, WIDTH)

	const (
		a = iota   // 0
		b          // 1 (= ++iota)
		c          // 2 (= ++iota)
		d = "haha" // 3 ( ++iota )
		e          // "haha" ( ++iota )
		f = iota   // 5 (= ++iota) // 恢复计数
		g          // 6
	)

	println(a, b, c, d, e, f, g)

	// len() cap() unsafe.Sizeof()
	println(unsafe.Sizeof(WIDTH))
}

func Run_var() {
	use_var()

	use_const()
}
