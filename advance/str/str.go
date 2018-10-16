package main

import (
	"fmt"
	"unicode/utf8"
)

// string 是字节切片

func printBytes(s string) { // 打印每一个字节
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("")
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // 以一个字节的方式打印
	}
	fmt.Println("")
}

func printByRune(s string) {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i]) // 以一个代码点的方式打印
	}
	fmt.Println("")
}

func main() {
	str := "Hello Go!" // s[i] -> 取字节
	printBytes(str)
	printChars(str)

	// 但是出现以两个字节表示的字符时 bytes方式会出现问题
	name := "Señor"
	printBytes(name)
	printChars(name) // S e Ã ± o r

	// 以rune方式打印 rune -> int32 表示一个代码点 无论代码点占用多少字节
	printByRune(name)

	for i, rune := range name {
		fmt.Printf("%c starts at byte %d and chartcode is %x", rune, i, rune)
		fmt.Println(" -> ", rune)
	}

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // byte切片 // Café
	fmt.Println(byteSlice)
	mstr := string(byteSlice)
	fmt.Println(mstr)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} // rune切片
	fmt.Println(runeSlice)
	rstr := string(runeSlice)
	fmt.Println(rstr)

	fmt.Printf("length of %s is %d\n", name, len(name)) // 6
	// 以rune方式获取字符串长度
	fmt.Printf("length of %s is %d\n", name, utf8.RuneCountInString(name)) // 5

	nameRune := []rune(name)
	fmt.Printf("length of %s is %d\n", string(nameRune), len(nameRune)) // 5

	// 字符串的不可变
	h := "hello"
	// h[0] = 'x' // error strike

	// -> 转为rune在改变
	hRune := []rune(h)
	hRune[0] = 'a'
	fmt.Println("mod string -> ", string(hRune))
}
