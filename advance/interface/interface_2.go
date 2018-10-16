package main

import "fmt"

type Describer interface {
	Desc()
}

type Mail interface { // 2. 多接口
	Deliver()
}

type Man struct {
	name string
	age  int
}

type Address struct {
	province string
	country  string
}

func (man Man) Desc() { // 值接受者
	fmt.Printf("%s is %d years old\n", man.name, man.age)
}

func (a *Address) Desc() { // 指针接受者
	fmt.Printf("%s Pro. of %s Country.\n", a.province, a.country)
}

func (a *Address) Deliver() {
	fmt.Printf("We will mail somthing to Address: %s %s\n", a.province, a.country)
}

type Shoper interface { // 3. 接口的嵌套
	Describer
	Mail
}

func main() {
	// 1. 接口的值接受者 与 指针接受者
	var d Describer
	man := Man{
		name: "mrzjd",
		age:  24,
	}
	d = man
	d.Desc()

	address := Address{
		country:  "china",
		province: "hubei",
	}
	d = &address
	d.Desc()

	var mail Mail = &address // 2. 实现多个接口
	mail.Deliver()

	var shoper Shoper = &address // 3. 接口的嵌套
	shoper.Desc()
	shoper.Deliver()

	var ss Shoper // 4. 接口的零值 nil
	if ss == nil {
		fmt.Printf("ss is nil and has type %T, value is %v\n", ss, ss)
	}
}
