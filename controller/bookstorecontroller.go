package controller

import (
	"encoding/json"
	"net/http"

	repository "example.com/BookStoreProj/storage"
)

type BookStoreController struct {
	Repository *repository.BookStoreRepository
}

// GetBooks returns the list of all books
func (c *BookStoreController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Repository.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}
func (c *BookStoreController) GetBookByName(w http.ResponseWriter, r *http.Request) {
	bookName := r.URL.Query().Get("bookName")
	bookNamee, err := c.Repository.GetBookByName()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bookNamee)

}
