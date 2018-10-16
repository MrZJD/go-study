package base

import "fmt"

// 结构体声明
// type struct_type struct {}
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// . 访问成员

func Run_struct() {
	var b1 = Books{
		"Golang",
		"somebody",
		"program lanuage",
		101,
	}

	var b2 = Books{
		title:   "不想学py, 不想写fe, what can i do",
		author:  "zzz",
		subject: "diary",
		book_id: 102,
	}

	var b3 = Books{
		title:   "Everything is nosense",
		book_id: 103,
	}

	fmt.Println("b1 ->", b1)
	fmt.Println("b2 ->", b2)
	fmt.Println("b3 ->", b3)

	getBookInfo(b1)
	getBookInfo(b2)
	getBookInfo(b3)

	// 结构体指针
	var ptrOfBook *Books

	ptrOfBook = &b1

	println("Using Struct Pointer:", ptrOfBook.title, ptrOfBook.author, ptrOfBook.subject, ptrOfBook.book_id)
}

func getBookInfo(book Books) {
	println("Book Title:", book.title)
	println("Book Author:", book.author)
	println("Book Subject:", book.subject)
	println("Book BookId:", book.book_id)
}
