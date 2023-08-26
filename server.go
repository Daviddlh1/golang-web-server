package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Daviddlh1/golang-web-server/handlers"
)

func main() {
	fmt.Printf("Starting server on port 8080\n")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/form", handlers.FormHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
