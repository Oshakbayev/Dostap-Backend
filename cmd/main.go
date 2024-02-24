package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "80"
	router := http.NewServeMux()
	router.HandleFunc("/", MainPage)

	fmt.Println("listening on: http://localhost:" + port + "/")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error: several connecrtions")
		return
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
