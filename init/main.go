package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("에러가 떠버렸는데요?")
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
