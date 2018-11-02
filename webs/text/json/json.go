package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Servers struct {
	Servers []Server
}

type Server struct {
	ServerName string
	ServerIP   string
}

/* 已知数据结构类型 */
func ReadByStruct() {
	file, err := os.Open("D:/repo/code/golang/src/learning/webs/text/json/target.json")
	checkErr(err)

	data, err := ioutil.ReadAll(file)
	checkErr(err)

	var s Servers

	err = json.Unmarshal(data, &s)
	checkErr(err)

	fmt.Println(s)
}

/* 解析到interface */
func ReadByInterface() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{} // 通过interface获取未知结构的json

	err := json.Unmarshal(b, &f)
	checkErr(err)

	fmt.Println("[After Parse]: ", f)

	// 通过断言获取数据
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type i don't know how to handle")
		}
	}
}

/* 使用第三方库 */
// func useSimpleJson() {
// 	// https://github.com/bitly/go-simplejson
// 	js, err := NewJson([]byte(`{
// 		"test": {
// 			"array": [1, "2", 3],
// 			"int": 10,
// 			"float": 5.150,
// 			"bignum": 9223372036854775807,
// 			"string": "simplejson",
// 			"bool": true
// 		}
// 	}`))

// 	arr, _ := js.Get("test").Get("array").Array()
// 	i, _ := js.Get("test").Get("int").Int()
// 	ms := js.Get("test").Get("string").MustString()
// }

func generateJSON() {
	var s Servers
	s.Servers = append(s.Servers, Server{
		ServerName: "SHANGHAI_VPN",
		ServerIP:   "127.0.0.1",
	})
	s.Servers = append(s.Servers, Server{
		ServerName: "GUANGZHOU_VPN",
		ServerIP:   "127.0.0.2",
	})

	b, err := json.Marshal(s)
	checkErr(err)

	fmt.Println(string(b))

	// `json:"serverName"` // -> 指定key
	// `json:"serverName,string"` // -> 转为字符串再输出
}

func main() {
	// ReadByStruct()
	// ReadByInterface()
	generateJSON()
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
