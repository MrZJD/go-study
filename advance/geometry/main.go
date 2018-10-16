package main

import ( // 1. init import pkg 初始化导入的包
	"fmt"
	"learning/advance/geometry/rectangle"
	"log"
)

// 包级别变量
var rectlen, rectwid float64 = 6, 7 // 2. 加载包级别变量 -> init()

// 用 _ (空白标识符) 来屏蔽 导入未使用的错误
// var _ = rectangle.Area

func init() { // 3. init -> 包被加载时调用，有且只调用一次，不参数及返回值，不能被显示调用
	fmt.Println("main pkg init!")

	// -> 可以用来检查包级别变量的正确性
	if rectlen < 0 {
		log.Fatal("length less than zero")
	}
	if rectwid < 0 {
		log.Fatal("width less than zero")
	}
}

func main() { // 4. main
	fmt.Println("Geometrical shape properties")

	fmt.Printf("area of rectangle is %.2f\n", rectangle.Area(rectlen, rectwid))
	fmt.Printf("diagonal of rectangle if %.2f\n", rectangle.Diagonal(rectlen, rectwid))
}
