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

	var result *cfg.User

	json.NewDecoder(r.Body).Decode(&result)

	fmt.Println("result decoded", result)

	result.RoleID = nil

	db.Create(result)

	if err != nil {
		http.Error(w, "Not able to read he body !", 404)
	}

	fmt.Println("Reached the post handle")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not valid route !", 404)
	}

	defer r.Body.Close()
	// body, err := ioutil.ReadAll(r.Body)

	// fmt.Println(body)

	var result []cfg.User

	errDb := db.Raw(`
		SELECT * 
		FROM user
	`).Scan(&result).Error

	if errDb != nil {
		http.Error(w, "Not able to get the results !", 404)
	}

	json.NewEncoder(w).Encode(result)

	fmt.Println("Reached the get  handle")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Success")

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

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/create-user", createUser)
	http.HandleFunc("/get-users", getUsers)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
