package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func DatabaseConnection() {
	// Open up our database connection.
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
}

func StartServer() {
	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func InvalidRequest() {
	var w http.ResponseWriter
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
}

func Error() {
	if err != nil {
		panic(err.Error())
	}
}
