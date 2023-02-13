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
	/*
		bk := BOOK_RECORD{BOOKID: 1, Title: "Harry Potter", Author: "J K Rowling", Year: "1997"}
		BOOKZ = append(BOOKZ, bk)

		bk1 := BOOK_RECORD{BOOKID: 2, Title: "Golang", Author: "Mr Google", Year: "2010"}
		BOOKZ = append(BOOKZ, bk1)

		bk2 := BOOK_RECORD{BOOKID: 3, Title: "Lord of Rings", Author: "J R Tolkein", Year: "2001"}
		BOOKZ = append(BOOKZ, bk2)

			BookDetailsController
		● GET /getBooks - to retrieve all books
		● GET /getBookByName/{bookName} - to retrieve a book by its name
		● GET /getBook/{bookId} - to retrieve a book by its id
		● POST /addBook - to add a new book to the repository
		● PUT /update/{bookId} - to update a book record by its id
		● DELETE /delete/{bookId} - to delete a book record by its id
	*/
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

/*
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
}*/

// ● GET /getBooks - to retrieve all books
func getBooksFullList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Execute the query
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord")
	if err != nil {
		panic(err.Error())
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

// ● GET /getBookByName/{bookName} - to retrieve a book by its name
func getBookByName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get book by name")
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
	fmt.Println("calling get func")
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
	w.Header().Set("Content-Type", "application/json")
	book_id, err := strconv.Atoi(r.URL.Path[len("/deleteBook/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteBookID, err := db.Exec("delete from BookRecord where BOOKID=?", book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	del_Book, err := deleteBookID.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(del_Book)
	w.Write([]byte("Book Record DELETED successfully..."))
}

// ● PUT /update/{bookId} - to update a book record by its id
func updateBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Book")
	//w.Header().Set("Content-Type", "application/json")
	book_id, err := strconv.Atoi(r.URL.Path[len("/updateBook/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	ress, err := db.Query("Update BookRecord set BOOKID=?,Title=? ,Author=? ,Year=? where BOOKID=?)", book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte("Book Record UPDATED successfully..."))
}
