package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) { // "/" default router
	r.ParseForm() // 解析传递的参数

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
		t, _ := template.ParseFiles("D:/repo/code/golang/src/learning/webs/forms/template/login.tpl") // 模板引擎
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm() // 1. 处理表单输入

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		fmt.Fprintf(w, "hello %v\n", r.Form["username"][0])
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		return
	}
	r.ParseForm()
	// 2. 过滤用户输入

	// 2.1 必填字段

	// 2.1.1 文本
	username := r.Form.Get("username")
	if len(username) == 0 { // 文本框 文件上传
		fmt.Fprintln(w, "you must input username without null")
		return
	}
	fmt.Fprintf(w, "hello %v\n", r.Form["username"][0])
	// 如果是复选框 单选框 为空值时 不会产生对应字段 需根据 r.From.Get() 来获取单个值 来判断是否存在 如果是多个值需通过map获取 r.Form["username"]

	// 2.1.2 数字
	age, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		fmt.Fprintln(w, "data: age should be number")
		return
	}

	// 2.1.3 数字也可以使用正则
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age2")); !m {
		fmt.Fprintln(w, "data: age2 should be number")
		return
	}

	// 2.1.4 中文
	// -> pkg:unicode / func Is(rangeTab *RangeTable, r rune) bool
	// -> regexp
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
		fmt.Fprintln(w, "data: realname can only be chinese")
		return
	}

	// 英文: regexp: "^[a-zA-Z]+$"
	// 电子邮件: regexp: `^([\w\.\_]{2, 10})@(\w{1,}).([a-z]{2, 4})$`
	// 手机号码: regexp: `^(1[3|4|5|8][0-9]\d{4,8})$`

	fmt.Println("age", age)
	fmt.Println("age2", r.Form.Get("age2"))

	// 2.2 下拉框 单选框
	target := []string{"apple", "pear", "banana"}
	v := r.Form.Get("fruits")
	flag := false
	for _, item := range target {
		if item == v {
			flag = true
			break
		}
	}
	if flag {
		fmt.Fprintf(w, "You Select is Accepted\n")
	} else {
		fmt.Fprintf(w, "Your Select can NOT be Accepted\n")
		return
	}

	// 2.3 复选框
	slice := []string{"football", "basketball", "tennis"}
	a := SliceDiff(r.Form["sports"], slice)
	if a == nil {
		fmt.Fprintf(w, "Your selectbox is Accepted")
	} else {
		fmt.Fprintf(w, "Your selectbox can NOT be Accepted")
	}

	// 2.4 others
	// times
	t := time.Date(2018, time.October, 31, 12, 29, 29, 0, time.Local)
	fmt.Printf("Go Launched at %s\n", t.Local())
	// id number
	// regexp: `^(\d{15})$`
	// regexp: `^(\d{17})([0-9]|X)$`

	// 3. xxs
	// 3.1 转义
	template.HTMLEscape(w, []byte(`<script>alert("xxs~")</script>`))
	// fmt.Println(template.HTMLEscapeString(`<script>alert("xxs~")</script>`))
	// fmt.Println(template.HTMLEscaper(`<script>alert("xxs~")</script>`))

	// 4. 多次提交form表单 -> input hidden token验证
}

func upload(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		currt := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currt, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("D:/repo/code/golang/src/learning/webs/forms/template/upload.tpl")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error1: ", err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile("./webs/forms/uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error2: ", err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

/* client upload file */
// requestUpload("./webs/forms/uploads/test.txt", "http://localhost:9090/upload")
func requestUpload(filename string, targetURL string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buf")
		return err
	}

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	res, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(res.Status)
	fmt.Println(string(resBody))
	return nil
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/getInfo", getInfo)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	} else {
		fmt.Println("We are listening at 9090")
	}
}

// InSlice target in s
func InSlice(target string, s []string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

// SliceDiff difference between s1 & s2
func SliceDiff(s1, s2 []string) (diffS []string) {
	for _, v := range s1 {
		if !InSlice(v, s2) {
			diffS = append(diffS, v)
		}
	}
	return
}
