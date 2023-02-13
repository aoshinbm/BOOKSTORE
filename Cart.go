/*
CartController
● GET /api/cart/decreaseQuantity/{cartID}: This API is used to decrease the quantity of a
cart record. It takes the cartID in the URL parameter and returns the ResponseDTO of
updated cart record with decreased quantity.
● GET /api/cart/increaseQuantity/{cartID}: This API is used to increase the quantity of a
cart record. It takes the cartID in the URL parameter and returns the ResponseDTO of
updated cart record with increased quantity.
*/

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

var CART []CART_DETAILS
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

	c1 := CART_DETAILS{CARTID: 1, ITEM_NAME: "Harry Potter Series", Quantity: 2, TOTAL_AMOUNT: 53000}
	CART = append(CART, c1)

	http.HandleFunc("/api/cart/getAll", getCartList)
	http.HandleFunc("/api/cart/getById/{cartId}", getCartItemById)
	http.HandleFunc("/api/cart/create", addPostCart)
	http.HandleFunc("/api/cart/updateById/{cartId}", updatePutCart)
	http.HandleFunc("/api/cart/delete/{cartId}", removeCartItemById)
	http.HandleFunc("/api/cart/deleteall", removeCart)
	//http.HandleFunc("/cart/{id}")

	//Cart_Routes()

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

/*
func handleCartDetails(w http.ResponseWriter, r *http.Request) {
	//OrderController
	//r.method returns which method is the request calling
	switch r.Method {
	case "GET":
	case "PUT":
		updatePutCart(w, r)
	case "POST":
		addPostCart(w, r)
	case "DELETE":
		removeCart(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
*/

// ● GET /api/cart/getAll: This API is used to retrieve all cart records.
// It does not take any parameters and returns the ResponseDTO of all cart details.
func getCartList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getCart CART_DETAILS
	results, err := db.Query("SELECT * FROM CARTDETAILS")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getCart.CARTID, &getCart.ITEM_NAME, &getCart.Quantity, &getCart.TOTAL_AMOUNT)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CART = append(CART, getCart)
	}
	json.NewEncoder(w).Encode(CART)
}

// ● GET /api/cart/getById/{cartId}: This API is used to retrieve cart record by cartId.
// It takes the cartId in the URL parameter and returns the ResponseDTO of specific cart details.
func getCartItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cart_id, err := strconv.Atoi(r.URL.Path[len("/api/cart/getById/"):])
	var getCartId CART_DETAILS
	results, err := db.Query("SELECT * FROM CARTDETAILS where CARTID=?", cart_id)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getCartId.CARTID, &getCartId.ITEM_NAME, &getCartId.Quantity, &getCartId.TOTAL_AMOUNT)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
	}
	json.NewEncoder(w).Encode(getCartId)
}

// ● POST /api/cart/create: This API is used to insert items in the cart.
// It accepts the token and CartDTO in the request body and returns the ResponseDTO of inserted item in the cart.
func addPostCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postCart CART_DETAILS
	_, err := db.Query("INSERT INTO CARTDETAILS (CARTID ,ITEM_NAME,Quantity,TOTAL_AMOUNT) VALUES (?,?,?,?)", postCart.CARTID, postCart.ITEM_NAME, postCart.Quantity, postCart.TOTAL_AMOUNT)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("Cart item added successfully..."))
}

// ● PUT /api/cart/updateById/{cartId}: This API is used to update cart by id.
// It takes the cartId in the URL parameter and the updated CartDTO in the request body and returns the ResponseDTO of updated cart record.
func updatePutCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cart_id, err := strconv.Atoi(r.URL.Path[len("/cart/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	ress, err := db.Exec("Update CARTDETAILS set ITEM_NAME=?,ADDRESS=?,AMOUNT=? where ORDERID=?)", cart_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cartUpdated, err := ress.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(cartUpdated)
	w.Write([]byte("CART UPDATED successfully..."))
}

// ● DELETE /api/cart/delete/{cartId}: This API is used to delete cart by id.
// It takes the cartId in the URL parameter and returns the ResponseDTO of deleted cart record.
func removeCartItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cart_id, err := strconv.Atoi(r.URL.Path[len("/cart/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteCartId, err := db.Exec("delete from CARTDETAILS where CARTID=?", cart_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	del_Cart, err := deleteCartId.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(del_Cart)
	w.Write([]byte("CART ITEM DELETED successfully..."))
}

// ● DELETE /api/cart/deleteall: This API is used to delete all items from the cart.
// It does not take any parameters and returns the list of deleted cart records.
func removeCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := strconv.Atoi(r.URL.Path[len("/cart"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteCart, err := db.Exec("delete from CARTDETAILS ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	del_Cart, err := deleteCart.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(del_Cart)
	w.Write([]byte("CART DELETED successfully..."))
}
