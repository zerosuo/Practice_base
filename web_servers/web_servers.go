package main

import (
	"net/http"
	"fmt"
	"log"
)

// 包http通过任何实现了http.Handler的值来响应HTTP请求

type Hello struct {

}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
