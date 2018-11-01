package main

import (
	"fmt"
	"html/template"
	"learning/webs/session/sess/msession"
	"log"
	"net/http"
	"strings"
	"time"
)

// 1. 生成全局唯一标识符 sessionid
// 2. 数据存储空间
// 3. sessionid 发送给客户端 (Set-Cookie / URL override)

// 1. 全局session管理器
// 2. sessionid唯一性
// 3. sessionid关联用户
// 4. 存储
// 5. 过期

// 1. 对于session会话劫持的问题
// a. cookie httpOnly
// b. 外加一个token (登陆成功后返回给客户端 客户端每次请求都带上该token)
// c. 间隔时间刷新sessionid

var globalSession *msession.Manager

func init() {
	globalSession, _ = msession.NewManager("memory", "GOSID", 3600)

	go globalSession.GC() // GC 协程
}

func main() {

	// var globalSession *Manager
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("We are listening at 9090")
	}

}

func login(w http.ResponseWriter, r *http.Request) {

	sess := globalSession.SessionStart(w, r)
	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("D:/repo/code/golang/src/learning/webs/forms/template/login.tpl")
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}

}

// session操作
func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)

	createtime := sess.Get("createtime")

	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSession.SessionDestory(w, r)
	}

	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int))+1)
	}

	fmt.Fprintf(w, "Your CountNum: %v", sess.Get("countnum"))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析传递的参数

	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	sess := globalSession.SessionStart(w, r)

	fmt.Println("[home]: ", sess.Get("username"), sess.SessionID())

	fmt.Fprintf(w, "hello go web! %s %s", sess.Get("username"), sess.SessionID())
}
