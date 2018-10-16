package main

import (
	"fmt"
)

/* 展示接口的实际使用场景 */

type Calculator interface {
	CalcSalary() int
}

type Permanent struct {
	eid int
	pay int
}

type Contact struct {
	eid      int
	pay      int
	basicpay int
}

func (p Permanent) CalcSalary() int { // 不需要implement 包含interface中的方法即隐性实现了接口
	return p.pay
}

func (c Contact) CalcSalary() int {
	return c.pay + c.basicpay
}

func totalExpense(s []Calculator) int {
	total := 0
	for _, v := range s {
		total += v.CalcSalary()
	}
	return total
}

func main() {
	var p1 = Permanent{eid: 1, pay: 1000}
	var p2 = Permanent{eid: 2, pay: 2000}
	var c1 = Contact{eid: 3, pay: 1500, basicpay: 800}

	fmt.Println("Total Expense is: ", totalExpense([]Calculator{p1, p2, c1}))

	// interface type
	var calc Calculator
	calc = p1
	fmt.Printf("interface type is %T, value is %v \n", calc, calc) // main.Permanent {1, 1000}

	// 1. 空接口
	// i interface {} // 空接口 -> 所有类型都实现了空接口

	// 2. 类型断言
	// i.(type) // i.(int) 断言接口为int类型 用于获取对应的接口
	var s interface{} = 56
	fmt.Println("What is s", s)
	i := s.(int)
	fmt.Println("What is s", i)
	// j := s.(string) // 会出现runtime error运行时错误
	// fmt.Println("Can Assert Wrong type", j)

	v, ok := s.(string) // 当出现不匹配类型时，v=该类型默认值，ok=false
	fmt.Println("Can Assert Wrong type", v, ok)

	// 3. 断言选择
	switch mm := s.(type) {
	case int:
		fmt.Println("mm is int", mm)
	case string:
		fmt.Println("mm is string", mm)
	default:
		fmt.Println("unkown type")
	}
}
