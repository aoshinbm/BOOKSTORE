package controller

import (
	"encoding/json"
	"net/http"

	repository "example.com/BookStoreProj/storage"
)

type CartController struct {
	Repository *repository.CartStoreRepository
}

func (c *CartController) GetC(w http.ResponseWriter, r *http.Request) {
	books, err := c.Repository.GetB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}
