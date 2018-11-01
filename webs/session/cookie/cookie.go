package main

// 1. http.SetCookie(w ResponseWriter, cookie *Cookie)

// type Cookie struct {
// 	Name       string
// 	Value      string
// 	Path       string
// 	Domain     string
// 	Expires    time.Time
// 	RawExpires string
// 	MaxAge     int
// 	Secure     bool
// 	HttpOnly   bool
// 	Raw        string
// 	Unparsed   []string
// }

// func setCookie () {
// 	expire := time.Now()
// 	expire = expire.AddDate(1, 0, 0)

// 	cookie := http.Cookie{
// 		Name: "username",
// 		Value: "mrzjd",
// 		Expires: expire
// 	}

// 	http.SetCookie(w, cookie)
// }

// func readCookie() {
// 	cookie, _ := r.Cookie("username")

// 	for _, cookie := range r.Cookie {
// 		fmt.Println(cookie.Name)
// 	}
// }

func main() {}
