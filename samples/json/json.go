package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Response1 结构体1
type Response1 struct {
	Page   int
	Fruits []string
	belong string
}

// Response2 结构体2
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// 1. 基础类型的json编码
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.1415926)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("hello go~")
	fmt.Println(string(strB))

	// 2. slice & map
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// 3. 自定义结构体
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		belong: "mrzjd",
	}
	res1B, _ := json.Marshal(res1D) // 只会解析可导出成员 并且键名默认为成员名
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}, // 可以使用struct tag自定义编码后的json键名
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 4. 解码
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // 基础类型
	fmt.Printf("%T, %v\n", dat["num"], dat["num"])
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // 嵌套结构
	str1 := strs[0].(string)
	// fmt.Println(dat["strs"][0]) // interface{} does not supporting indexing [0]
	fmt.Println(str1)

	// 4.1 解码为自定义类型
	str := `{"Page":2,"Fruits":["banana","cherry","tomato"]}`
	res := &Response2{}

	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 5. 编码解码流 (stdout files http.reponse)
	enc := json.NewEncoder(os.Stdout)

	d := map[string]int{"apple": 5, "lettucr": 7}

	enc.Encode(d)
}
