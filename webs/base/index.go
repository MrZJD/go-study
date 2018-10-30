package main

import (
	"fmt"
	"log"
	"net/http"
)

// 内部原理
// 1. -> net.Listen("tcp", port) // tcp监听端口
// 2. -> ser.Serve(I net.Listener) // 开启服务
//    -> for { rw := I.Accept(); c := ser.NewConn(); go c.serve() } // 循环接收请求
// 3. -> go c.serve() 开启协程 处理响应
//    -> c.readRequest() // response, request
//    -> handler.ServeHTTP(response, request) // handler -> func / DefaultServeMux

func sayHelloWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", v)
	}

	fmt.Fprintf(w, "Hello Go Web")
}

func main() {

	http.HandleFunc("/", sayHelloWeb) // 设置路由

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
