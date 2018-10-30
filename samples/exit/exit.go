package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("defer in main func!")

	os.Exit(3) // 程序立即退出 不会执行defer等后续所有行为 // windows: exit status 3
}
