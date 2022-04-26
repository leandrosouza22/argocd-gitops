package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello DevSecOps!!!</h1>"))
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("ec3ca5d1-02dc-4a61-ba2f-f7fc109683b0")
}
