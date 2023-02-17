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
	BOOKID int    `json:bookid`
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
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/getBooks", getBooksFullList)
	http.HandleFunc("/getBookByName/", getBookByName)
	http.HandleFunc("/getBook/", getBook)
	http.HandleFunc("/addBook", addPostBooks)
	http.HandleFunc("/updateBook/", updateBooks)
	http.HandleFunc("/deleteBook/", removeBook)

	//Book_Routes()

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// ● GET /getBooks - to retrieve all books
func getBooksFullList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Books record")
	w.Header().Set("Content-Type", "application/json")
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var getBook BOOK_RECORD
		err := results.Scan(&getBook.BOOKID, &getBook.Title, &getBook.Author, &getBook.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		BOOKZ = append(BOOKZ, getBook)
	}
	json.NewEncoder(w).Encode(BOOKZ)
}

// ● GET /getBookByName/{bookName} - to retrieve a book by its name
func getBookByName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a Book record by name")
	//w.Header().Set("Content-Type", "application/json")
	bookName := r.URL.Query().Get("bookName")
	//:= r.URL.Path[len("/getBookByName/"):]
	var getBuk BOOK_RECORD
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where Title=?", bookName)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getBuk.BOOKID, &getBuk.Title, &getBuk.Author, &getBuk.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(w).Encode(getBuk)
}

// ● GET /getBook/{bookId} - to retrieve a book by its id
func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a Book record by Id")
	w.Header().Set("Content-Type", "application/json")
	//bookId, err := strconv.Atoi(r.URL.Query().Get("bookId"))
	bookIDStr := r.URL.Path[len("/getBook/"):]
	bookId, err := strconv.Atoi(bookIDStr)

	var getBookId BOOK_RECORD
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where BOOKID=? ", bookId)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(results)
	for results.Next() {
		err := results.Scan(&getBookId.BOOKID, &getBookId.Title, &getBookId.Author, &getBookId.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(w).Encode(getBookId)
}

// ● POST /addBook - to add a new book to the repository
func addPostBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Book in Books record ")
	w.Header().Set("Content-Type", "application/json")
	var postBook BOOK_RECORD
	json.NewDecoder(r.Body).Decode(&postBook)
	_, err := db.Query("INSERT INTO BookRecord (BOOKID ,Title ,Author ,Year) VALUES (?,?,?,?)", postBook.BOOKID, postBook.Title, postBook.Author, postBook.Year)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("Data added successfully..."))
}

// ● DELETE /delete/{bookId} - to delete a book record by its id
func removeBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remove Book from Books record ")
	w.Header().Set("Content-Type", "application/json")
	book_id, err := strconv.Atoi(r.URL.Path[len("/deleteBook/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteBookID, err := db.Query("delete from BookRecord where BOOKID=?", book_id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(deleteBookID)
	w.Write([]byte("Book Record DELETED successfully..."))
}

// ● PUT /update/{bookId} - to update a book record by its id
func updateBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Book ")
	book_id, err := strconv.Atoi(r.URL.Path[len("/updateBook/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	ress, err := db.Query("Update BookRecord set BOOKID=?,Title=? ,Author=? ,Year=? where BOOKID=?)", book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(ress)
	w.Write([]byte("Book Record UPDATED successfully..."))
}
