package main

import (
	"log"
	"net/http"
)

func main() {
	// github.com/julienschmidt/httprouter

	router := httprouter.New()

	router.GET("/", Index)
	router.POST("/hello/:name", Hello)

	router.GET("/user/:uid", getUser)
	router.POST("/adduser/:uid", addUser)
	router.DELETE("/deleteuser/:uid", deleteUser)
	router.PUT("/moduser/:uid", modUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
