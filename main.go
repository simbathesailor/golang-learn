package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sequoia-backend-assignment/packages/trelloboard/cfg"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not valid route !", 404)
	}

	defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)
	// fmt.Println(body)

	var result cfg.User

	json.NewDecoder(r.Body).Decode(&result)

	fmt.Println("result decoded", result)

	result.RoleID = nil

	db.Create(result)

	if err != nil {
		http.Error(w, "Not able to read he body !", 404)
	}

	fmt.Println("Reached the post handle")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	c := 1
	user := cfg.User{ID: 2, Email: "sudhir.chaudhary@sequoia.com", RoleID: &c}

	fmt.Println("The user is : -=>", user)

	db.Create(user)

	fmt.Fprintf(w, "Hello 3!")

}

func main() {

	t := "Hello World"
	fmt.Println("t is", t)

	db, err = cfg.IntializeDatabase()

	defer db.Close()

	fmt.Println(db, err)
	// Handle the /hello request get request
	// func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello!")
	// }

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/create-user", createUser)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
