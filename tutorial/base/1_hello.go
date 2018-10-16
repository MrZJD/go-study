package base // 1. 当前包声明 一个程序应有一个main主包

import "fmt"

func Run_hello() { // 2. 主函数 init() -> main() // 3. 左花括号不能换行 (运行时错误)
	fmt.Println("Hello Go!") // 3. ; 语句结束后分号 可写也可不写 (编译时添加)
}
