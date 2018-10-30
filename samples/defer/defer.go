package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("D:\\repo\\code\\golang\\src\\learning\\samples\\defer\\defer.txt")

	defer closeFile(f)

	writeFile(f, "hello go")
}

func createFile(path string) *os.File {
	fmt.Println("Creating Files...")

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File, data string) {
	fmt.Println("Writing Files...")

	fmt.Fprintln(f, data)
}

func closeFile(f *os.File) {
	fmt.Println("Closing Files...")
	f.Close()
}
