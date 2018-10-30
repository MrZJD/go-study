package math

import "errors"

// Add 加法运算
func Add(a, b int) int {
	return a + b
}

// Subtract 减法运算
func Subtract(a, b int) int {
	return a - b
}

// Multiply 乘法运算
func Multiply(a, b int) int {
	return a * b
}

// Divide 除法运算
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}
