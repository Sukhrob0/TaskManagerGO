package main

import (
	"fmt"
	"net/http"
)

func main() {

	n := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("мой первый сервер")
	}

	http.ListenAndServe(":5050", nil)

	http.HandleFunc("/start", n)

}
