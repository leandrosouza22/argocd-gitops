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
	fmt.Println("a4db08b7-5729-4ba9-8c08-f2df493465a1")
}
