package main

import (
	"fmt"
	"time"
)

//
//import (
//	"fmt"
//	"net/http"
//)

func hello(async chan bool) {
	fmt.Println("hello world goroutine")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")

	async <- true
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully write", i, "to ch .")
		time.Sleep(1 * time.Second)

	}
	close(ch)
}

func main() {
	//async := make(chan bool)
	//go hello(async)
	//<- async
	//fmt.Println("main func")

	ch := make(chan int, 2)
	go write(ch)
	//time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch .")
		//time.Sleep(2 * time.Second)
	}
	fmt.Println("complated")
	//http.HandleFunc("/zc", hello)
	//
	//http.ListenAndServe(":8080", nil)

}

//
//func hello(w http.ResponseWriter, r *http.Request) {
//
//	fmt.Fprintf(w, "Hello Docker Form Golang!")
//
//}
