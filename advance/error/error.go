package main

import (
	"fmt"
	"path/filepath"
)

// type error interface {
// 	Error() string
// }

// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

// func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

// type DNSError struct {
// 	...
// }

// func (e *DNSError) Error() string { ... }

// func (e *DNSError) Timeout() bool { ... }

// func (e *DNSError) Temporary() bool { ... }

// var ErrBadPattern = errors.New("syntax error in pattern")

// 1. 如果一个函数 或方法 返回了错误，按照惯例，错误会作为最后一个值返回
// 2. 按照 Go 的惯例，在处理错误时，通常都是将返回的错误与 nil 比较。nil 值表示了没有错误发生，而非 nil 值表示出现了错误。

// 3. 从错误获取更多信息的方法
// 3.1 利用底层结构体类型断言(属性 方法)
// 3.2 直接比较

// 4. 不要忽略错误

func main() {
	// f, err := os.Open("/txt.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(f.Name(), "opened successfully")

	// 3.1 利用断言获取属性
	// if err, ok := err.(*os.PathError); ok {
	// 	fmt.Println("File at path", err.Path, "failed to open")
	// 	return
	// }
	// fmt.Println(f.Name(), "opened successfully")

	// 3.2 利用断言获取方法
	// addr, err := net.LookupHost("qwas123.com")
	// if err, ok := err.(*net.DNSError); ok {
	// 	if err.Timeout() {
	// 		fmt.Println("opration timed out!")
	// 	} else if err.Temporary() {
	// 		fmt.Println("temporary error!")
	// 	} else {
	// 		fmt.Println("generic error:", err)
	// 	}
	// 	return
	// }
	// fmt.Println(addr)

	// 3.3 直接比较
	// files, err := filepath.Glob("[")
	// if err != nil && err == filepath.ErrBadPattern {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("matched files", files)

	// 4. 不要忽略错误
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files) // 这里依然可以输出
}
