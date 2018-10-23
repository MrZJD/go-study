package main

import (
	"fmt"
	"reflect"
)

// 1. Reflection in Go
// -> 在运行阶段检查变量的类型

// 2. reflect.TypeOf() // 获取实际类型 main.order
//    reflect.ValueOf() // 取值
//    reflect.TypeOf().Kind() // struct
//    reflect.ValueOf(q).NumFileds() // 字段数量
//    reflect.ValueOf().Field(i) // 取字段
//    reflect.ValueOf().Field(i).Kind() // 取字段类型
//    reflect.ValueOf().Field(i).String() // 取值为String

type order struct {
	oid        int
	customerid int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	fmt.Println("Type ", t)  // main.order
	fmt.Println("Value ", v) // {oid customerid}

	k := t.Kind()
	fmt.Println("Kind ", k) // struct
}

func query(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s value(", t)

		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Unsupported type")
}

func main() {
	o := order{
		oid:        456,
		customerid: 901,
	}
	createQuery(o)

	query(o)

	e := employee{
		name:    "mrzjd",
		id:      101,
		address: "wuhan",
		salary:  0,
		country: "China",
	}

	query(e)
}
