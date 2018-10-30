package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
	muxEntry map[string]func()
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloWeb2(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayHelloWeb2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello My Go World!")
}

func main() {
	mux := &MyMux{}

	http.ListenAndServe(":9090", mux)
}
