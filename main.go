package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	t := "Hello World"
	fmt.Println("t is", t)

	// Handle the /hello request get request

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
