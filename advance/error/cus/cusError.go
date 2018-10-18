package main

import (
	"fmt"
	"math"
)

/* 自定义错误 */

// 1. 使用New函数创建自定义错误
// 2. fmt.Errorf
// 3. struct Error
// 4. struct Error Methods

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %.2f: %s", e.radius, e.err)
}

func (e *areaError) radiusNegative() bool {
	return e.radius < 0
}

func circleArea(radius float64) (float64, error) {
	// if radius < 0 { // 1. 自定义错误
	// 	return 0, errors.New("Area calculation failed, Radius can not less than zero!")
	// }

	if radius < 0 { // 2. Errorf
		// return 0, fmt.Errorf("Area calculation failed, Radius %.2f can not less than zero!\n", radius)
		return 0, &areaError{"Rsadius is less than zero", radius} // 3. struct error
	}

	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {

		// if err, ok := err.(*areaError); ok {
		// 	fmt.Printf("Radius %.2f is less than zero\n", err.radius)
		// 	return
		// }

		if err, ok := err.(*areaError); ok {
			if err.radiusNegative() {
				fmt.Printf("error: Radius %.2f is less than zero\n", err.radius)
			}
			return
		}

		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %.2f\n", area)
}
