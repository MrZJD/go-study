package rectangle // 同一个文件夹属于同一包 其内部文件的包声明 应该保持一致 (约束:与文件夹名相同)

import "math"

func init() {
	println("rectangle pkg init!")
}

// 函数名首字母大写 表示 该函数export
// 函数名小写 表示 该函数为内部函数
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
