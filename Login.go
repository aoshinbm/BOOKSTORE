package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type LOGIN struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB
var PERSON []LOGIN
var err error

func main() {
	// Open up our database connection.
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/login/{id}", handleLogin)

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))

}
func handleLogin(w http.ResponseWriter, r *http.Request) {
	//r.method returns which method is the request calling
	switch r.Method {
	case "GET":
		getLoginList(w, r)
	//case "POST":
	//	addPostLoginUsers(w, r)
	//case "PUT":
	//updatePostUsers(w, r)
	//case "DELETE":
	//	removeUser(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func getLoginList(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	results, err := db.Query("SELECT Username ,Password FROM UserRegistration")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var resul USERREGIST
		// for each row, scan the result into our tag composite object
		err := results.Scan(&resul.Username, &resul.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		PERSON = append(PERSON, resul)
	}
	json.NewEncoder(w).Encode(PERSON)
}
