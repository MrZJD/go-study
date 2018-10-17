package employee

import "fmt"

/* -> 使用struct 代替类 */

type employee struct { /* 改为小写 从而通过New进行调用 */
	firstName string
	lastName  string
	male      string
	age       int
	job       string
}

func (e employee) GetInfo() {
	fmt.Printf("%s %s, %s, %d, %s\n", e.firstName, e.lastName, e.male, e.age, e.job)
}

func New(firstName string, lastName string, male string, age int, job string) employee {
	e := employee{firstName, lastName, male, age, job}
	return e
}
