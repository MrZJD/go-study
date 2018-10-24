package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// 1. fs 读取文件到内存

// 2.1 绝对路径来标记文件
// 2.2 命令行来标记文件
// 2.3 文件绑定到二进制文件中 packr

// 3. 分块读取文件
func readByBuf() {
	fptr := "./txt.txt"

	f, err := os.Open(fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f) // buffered reader
	b := make([]byte, 3)    // 长度为3的字节切片
	for {
		_, err := r.Read(b) // len(b) 打到 3 返回所读取的字节
		if err != nil {
			fmt.Println("Error reading file: ", err) // 读到文件的末尾会返回一个EOF Error
			break
		}
		fmt.Println(string(b))
	}
}

// 4. 逐行读取文件
func readByLine() {
	fptr := "./doc/line.txt"

	f, err := os.Open(fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 1.
	data, err := ioutil.ReadFile("./txt.txt") // data -> binary slice 字节切片
	if err != nil {
		fmt.Println("File reading error: ", err)
		return
	}
	fmt.Println("Content of file: ", string(data)) // go install 这个程序后如果file与exe相对路径发生改变则会报错

	// 2.1 绝对路径
	file2, err2 := ioutil.ReadFile("D:/repo/code/golang/src/learning/apis/files/txt.txt")
	if err2 != nil {
		fmt.Println("File2 reading error2: ", err2)
		return
	}
	fmt.Println("Content of file2: ", string(file2))

	// 2.2 命令行
	// fptr := flag.String("fpath", "test.txt", "file path to read from")
	// flag.Parse()
	// fmt.Println("value of fpath is ", *fptr) // go run file.go -fpath=D:/repo/code/golang/src/learning/apis/files/txt.txt

	// file3, err3 := ioutil.ReadFile(*fptr)
	// if err3 != nil {
	// 	fmt.Println("File reading error3: ", err3)
	// 	return
	// }
	// fmt.Println("Content of file3: ", string(file3))

	// // 2.3 将文件绑定到二进制文件之中

	// "github.com/gobuffalo/packr"
	// box := packr.NewBox("./doc") // 2.3.1 -> 实际转为了绝对路径
	// data4 := box.String("test.txt")
	// fmt.Println("Contents of file4: ", data4) // 2.3.2 packr install -v ${GOPATH}/learing/apis/files -> 把文件打包进入二进制exe之中

	// 3. by Chunk
	readByBuf()

	// 4. by Line
	readByLine()
}
