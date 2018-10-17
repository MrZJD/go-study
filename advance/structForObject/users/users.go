package users

import (
	"fmt"
	"time"
)

type man struct {
	name string
	male string
	age  int
}

type user struct {
	uid      int64
	nickname string
	pswd     string
	man      /* 继承 */
}

func (a man) getBio() string {
	return fmt.Sprintf("%s %s %d", a.name, a.male, a.age)
}

func (u user) GetBio() string {
	return u.man.getBio()
}

func (u user) Login(nname string, pswd string) bool {
	return u.nickname == nname && u.pswd == pswd
}

func New(nickname string, pswd string, name string, male string, age int) user {
	m := man{
		name: name,
		male: male,
		age:  age,
	}
	return user{
		uid:      time.Now().Unix(),
		nickname: nickname,
		pswd:     pswd,
		man:      m,
	}
}

type service struct {
	users []user /* 结构体切片嵌套 */
}
