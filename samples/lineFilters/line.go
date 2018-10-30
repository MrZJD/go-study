package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// line filters 行过滤器
// -> stdin 读取数据 -> line filters -> stdout 标准输出
// grep sed

func main() {

	// os.Stdin -> 无缓冲
	// bufio -> 缓冲scanner

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "errors: ", err)

		os.Exit(1)
	}

}
