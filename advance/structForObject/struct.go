package main

import (
	"fmt"
	"learning/advance/structForObject/employee"
	"learning/advance/structForObject/users"
)

/* 使用结构体 替代 类 */

// 1. 使用New函数 而非构造器 // 用于初始化变量等操作

// 2. 组合取代继承

// 3. 使用接口实现多态

func main() {
	// e := employee.Employee{
	// 	FirstName: "Ed",
	// 	LastName:  "Sherren",
	// 	Male:      "male",
	// 	Age:       28,
	// 	Job:       "singer",
	// }
	e := employee.New("Ed", "Sherren", "male", 28, "singer") // 1

	e.GetInfo()

	u := users.New("YH", "Ed123", "Ed", "male", 28)

	canlog := u.Login("YH", "Ed12")
	fmt.Println("YH User canlog: ", canlog)
	fmt.Println("GetBio:", u.GetBio())
}
