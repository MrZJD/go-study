package main

import (
	"crypto/sha1"
	"fmt"
)

// SHA1: 散列 经常用来计算二进制或者大文本数据的短标识值

func main() {
	s := "sha1 this string"

	h := sha1.New() // 1

	h.Write([]byte(s)) // 2

	bs := h.Sum(nil) // 3. 计算 参数也可以是追加额外的字节

	fmt.Println(s)

	fmt.Printf("%x\n", bs)
}
