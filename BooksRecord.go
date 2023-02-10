package main

import (
	"database/sql"
	"encoding/json"
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
	DatabaseConnection()

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
	StartServer()
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
	w.Header().Set("Content-Type", "application/json")
	var getBuk BOOK_RECORD
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where Title=?", getBuk.Title)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getBuk.BOOKID, &getBuk.Title, &getBuk.Author, &getBuk.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		BOOKZ = append(BOOKZ, getBuk)
	}
	json.NewEncoder(w).Encode(BOOKZ)
}

// ● GET /getBook/{bookId} - to retrieve a book by its id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getBookId BOOK_RECORD
	results, err := db.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where BOOKID=?", getBookId.BOOKID)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getBookId.BOOKID, &getBookId.Title, &getBookId.Author, &getBookId.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		BOOKZ = append(BOOKZ, getBookId)
	}
	json.NewEncoder(w).Encode(BOOKZ)
}

// ● POST /addBook - to add a new book to the repository
func addPostBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

// ● DELETE /delete/{bookId} - to delete a book record by its id
func removeBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book_id, err := strconv.Atoi(r.URL.Path[len("/books/{id}"):])
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
func updatePutBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book_id, err := strconv.Atoi(r.URL.Path[len("/books/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	ress, err := db.Exec("Update BookRecord set BOOKID=?,Title=? ,Author=? ,Year=? where BOOKID=?)", book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	bookUpdated, err := ress.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(bookUpdated)
	w.Write([]byte("Book Record UPDATED successfully..."))
}
