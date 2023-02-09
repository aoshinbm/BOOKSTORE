package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// model of book
// models will specify the attributes of the book record that will be added to the books database
type BOOK_RECORD struct {
	BOOKID int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

// slice of books is created to hold the book records
var BOOKZ []BOOK_RECORD
var db *sql.DB
var err error

func main() {
	// Open up our database connection.
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/BOOKSTORE")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	bk := BOOK_RECORD{BOOKID: 1, Title: "Harry Potter", Author: "J K Rowling", Year: "1997"}
	BOOKZ = append(BOOKZ, bk)

	bk1 := BOOK_RECORD{BOOKID: 2, Title: "Golang", Author: "Mr Google", Year: "2010"}
	BOOKZ = append(BOOKZ, bk1)

	bk2 := BOOK_RECORD{BOOKID: 3, Title: "Lord of Rings", Author: "J R Tolkein", Year: "2001"}
	BOOKZ = append(BOOKZ, bk2)

	bk3 := BOOK_RECORD{BOOKID: 4, Title: "Revolution 2020", Author: "Chetan Bhagat", Year: "2011"}
	BOOKZ = append(BOOKZ, bk3)

	bk4 := BOOK_RECORD{BOOKID: 5, Title: "Half Girlfriend", Author: "Chetan Bhagat", Year: "2014"}
	BOOKZ = append(BOOKZ, bk4)

	http.HandleFunc("/books", handleBookRecords)
	http.HandleFunc("/books/{id}", handleBookRecords)

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// handleStocks
// 2 parameters :- response is interface, request is struct
func handleBookRecords(w http.ResponseWriter, r *http.Request) {
	//r.method returns which method is the request calling
	//
	switch r.Method {
	case "GET":
		getBooksFullList(w, r)
	case "POST":
		addPostBooks(w, r)
	case "PUT":
		updatePutBooks(w, r)
	case "DELETE":
		removeBook(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getBooksFullList(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var getBook BOOK_RECORD
		// for each row, scan the result into our tag composite object
		err := results.Scan(&getBook.BOOKID, &getBook.Title, &getBook.Author, &getBook.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		BOOKZ = append(BOOKZ, getBook)
	}
	json.NewEncoder(w).Encode(BOOKZ)
}
func addPostBooks(w http.ResponseWriter, r *http.Request) {
	var postBook BOOK_RECORD
	json.NewDecoder(r.Body).Decode(&postBook)
	//read from the request
	// Execute the query
	_, err := db.Query("INSERT INTO BookRecord (BOOKID ,Title ,Author ,Year) VALUES (?,?,?,?)", postBook.BOOKID, postBook.Title, postBook.Author, postBook.Year)
	//Scan(&postBook.BOOKID, &postBook.Title, &postBook.Author, &postBook.Year)
	if err != nil {
		panic(err.Error())
	}
	//BOOKZ = append(BOOKZ, postBook)
	//json.NewEncoder(w).Encode(BOOKZ) //optional
	w.Write([]byte("Data added successfully..."))
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	fmt.Println(r.URL.Path[len("/books/"):])
	path := "/books/12"
	bukid := path[8:]
	fmt.Println(bukid)

	result, err := db.Query("delete from product where id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	//BOOKZ = append(BOOKZ, postBook)
	//json.NewEncoder(w).Encode(BOOKZ) //optional
	w.Write([]byte("Data deleted..."))

	for i, data_id := range BOOKZ {
		if data_id.BOOKID == id {
			BOOKZ = append(BOOKZ[:i], BOOKZ...)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
func updatePutBooks(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r.URL.Path[len("/books/"):])
	path := "/books/1"
	book_id := path[:]
	fmt.Println(book_id)

	for i, bookdata := range BOOKZ {
		if bookdata.BOOKID == id {
			var newBook BOOK_RECORD
			json.NewDecoder(r.Body).Decode(&newBook)
			newBook.BOOKID = id
			BOOKZ[i] = newBook
			json.NewEncoder(w).Encode(newBook)
		}
	}
}
