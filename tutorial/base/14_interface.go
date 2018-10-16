package base

// 定义接口
// type interface_name interface {
// 	method_1 [return_type]
// 	method_2 [return_type]
// 	method_3 [return_type]
// 	...
// 	method_4 [return_type]
// }

// 定义结构体
// type struct_name struct {
/* variables */
// }

// 实现接口方法
// func (struct_name_var struct_name) method_1() [return_type] {}

type Phone interface {
	call()
}

type NokiaPhone struct {
	phoneType string
}

func (np NokiaPhone) call() {
	println("i'm nokia phone")
}

type IPhone struct {
	phoneType string
}

func (ip IPhone) call() {
	println("i'm iphone")
}

func Run_interface() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
