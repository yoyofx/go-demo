package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	ht := http.HandlerFunc(helloHandler)
	if ht != nil {
		http.Handle("/hello", ht)
	}
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

var iCnt int = 0

func helloHandler(w http.ResponseWriter, r *http.Request) {
	iCnt++
	str := "Hello world ! friend(" + strconv.Itoa(iCnt) + ")"
	io.WriteString(w, str)
	fmt.Println(str)
}
