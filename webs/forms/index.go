package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) { // "/" default router
	r.ParseForm() // 解析url传递的参数

	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	fmt.Fprintln(w, "hello go web!")
}

func login(w http.ResponseWriter, r *http.Request) { // "/login" router
	fmt.Println("method: ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("D:/repo/code/golang/src/learning/webs/forms/template/login.gtpl") // 模板引擎
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()

		username := r.Form["username"]
		fmt.Println("username", username)
		fmt.Println("password", r.Form["password"])

		fmt.Fprintf(w, "hello %v", username)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}

	fmt.Println("We are listening at 9090")
}
