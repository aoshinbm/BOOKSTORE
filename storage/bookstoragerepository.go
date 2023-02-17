package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"example.com/BookStoreProj/types"
)

type BookStoreRepository struct {
	DB *sql.DB
}

// funcs to handle queries to database
func (r *BookStoreRepository) GetBooks() ([]*types.BOOK_RECORD, error) {
	fmt.Println("Get All Books record")
	results, err := r.DB.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord")
	if err != nil {
		panic(err.Error())
	}
	//store
	var BOOKZ []*types.BOOK_RECORD
	for results.Next() {
		//fetch data
		getBooks := &types.BOOK_RECORD{}
		err := results.Scan(&getBooks.BOOKID, &getBooks.Title, &getBooks.Author, &getBooks.Year)
		if err != nil {
			return nil, err
		}
		BOOKZ = append(BOOKZ, getBooks)
	}
	return BOOKZ, nil
}

func (r *BookStoreRepository) GetBookByName() ([]*types.BOOK_RECORD, error) {
	fmt.Println("Get a Book record by name")
	results, err := r.DB.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where Title=?", bookName)
	if err != nil {
		panic(err.Error())
	}
	getBukbyName := &types.BOOK_RECORD{}
	for results.Next() {
		//fetch data
		err := results.Scan(&getBukbyName.BOOKID, &getBukbyName.Title, &getBukbyName.Author, &getBukbyName.Year)
		if err != nil {
			return nil, err
		}
	}
	return getBukbyName, nil
}

func (r *BookStoreRepository) GetBookById() ([]*types.BOOK_RECORD, error) {
	fmt.Println("Get a Book record by Id")
	bookIDStr := r.URL.Path[len("/getBook/"):]
	bookId, err := strconv.Atoi(bookIDStr)
	results, err := r.DB.Query("SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where BOOKID=? ", bookId)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		//fetch data
		getBookbyId := &types.BOOK_RECORD{}
		err := results.Scan(&getBookbyId.BOOKID, &getBookbyId.Title, &getBookbyId.Author, &getBookbyId.Year)
		if err != nil {
			return nil, err
		}
	}
	return getBookbyId, nil
}

func (r *BookStoreRepository) AddBook() ([]*types.BOOK_RECORD, error) {
	fmt.Println("Add Book in Books record ")
	var postBook *types.BOOK_RECORD
	resultt, err := r.DB.Query("INSERT INTO BookRecord (BOOKID ,Title ,Author ,Year) VALUES (?,?,?,?)", postBook.BOOKID, postBook.Title, postBook.Author, postBook.Year)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data added successfully...")
}

func (r *BookStoreRepository) RemoveBook() ([]*types.BOOK_RECORD, error) {
	fmt.Println("Remove Book from Books record ")
	book_id, err := strconv.Atoi(r.URL.Path[len("/deleteBook/"):])
	if err != nil {
		panic(err.Error())
	}
	deleteBookID, err := r.DB.Query("delete from BookRecord where BOOKID=?", book_id)
	if err != nil {
		panic(err.Error())
	}
	return deleteBookID, nil
	fmt.Println("Book Record DELETED successfully...")
}
func (r *BookStoreRepository) UpdateBook() ([]*types.BOOK_RECORD, error) {
	fmt.Println("Update Book ")
	book_id, err := strconv.Atoi(r.URL.Path[len("/updateBook/"):])
	if err != nil {
		panic(err.Error())
	}
	ress, err := r.DB.Query("Update BookRecord set BOOKID=?,Title=? ,Author=? ,Year=? where BOOKID=?", book_id)
	if err != nil {
		panic(err.Error())
	}
	return ress, nil
	fmt.Println("Book Record UPDATED successfully...")
}
