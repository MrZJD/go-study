package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 100
	var b = 200

	var val = Add(a, b)

	if val != a+b {
		t.Error("Test Case[", "TestAdd", "] failed!")
	}
}

func TestSubtract(t *testing.T) {
	a := 100
	b := 200

	val := Subtract(a, b)

	if val != a-b {
		t.Error("Test Case[", "TestSubtract", "] failed!")
	}
}

func TestMultiply(t *testing.T) {
	a := 100
	b := 200

	val := Multiply(a, b)

	if val != a*b {
		t.Error("Test Case[", "TestMultiply", "] failed!")
	}
}

func TestDivide(t *testing.T) {
	a := 100
	b := 200
	// c := 0

	val, _ := Divide(a, b)

	if val != a/b {
		t.Error("Test Case[", "TestDivide", "] failed!")
	}
}

func TestDivideZero(t *testing.T) {
	a := 100
	b := 0

	_, err := Divide(a, b)

	if err == nil {
		t.Error("Test Case[", "TestDivideZero", "] failed!")
	}
}

// 压力测试
func BenchmarkAdd(b *testing.B) {
	b.StopTimer()
	// .. 中间用来做一些初始化工作
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Add(4, 5)
	}
}
