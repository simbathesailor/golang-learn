package main

import (
	"fmt"
	"log"
	"net/http"
	"sequoia-backend-assignment/packages/trelloboard/cfg"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")

}

func main() {

	t := "Hello World"
	fmt.Println("t is", t)

	db, err := cfg.IntializeDatabase()

	fmt.Println(db, err)
	// Handle the /hello request get request
	// func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello!")
	// }

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
