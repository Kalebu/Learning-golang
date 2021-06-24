package main

import (
	"fmt"
	"httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func spam(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "Hello Kalebu, You made your very first go route function !!\n")
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/spam", spam)
	fmt.Println("Server running at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
