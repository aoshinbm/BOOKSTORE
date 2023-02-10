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
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type CART_DETAILS struct {
	CARTID       int    `json:id`
	ITEM_NAME    string `json:order_name`
	Quantity     int    `json:quantity`
	STATUS       string `json:status`
	TOTAL_AMOUNT int    `json:total_amount`
}

var CART []ORDER_DETAILS
var db *sql.DB
var err error

func main() {
	DatabaseConnection()

	http.HandleFunc("/cart", handleOrderDetails)
	http.HandleFunc("/cart/{id}", handleOrderDetails)

	//start the server
	StartServer()
}

func handleCartDetails(w http.ResponseWriter, r *http.Request) {
	//OrderController
	//r.method returns which method is the request calling
	switch r.Method {
	case "GET":
		getCartList(w, r)
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

// ● GET /api/cart/getAll: This API is used to retrieve all cart records.
// It does not take any parameters and returns the ResponseDTO of all cart details.
func getCartList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	results, err := db.Query("SELECT CARTID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS FROM CARTDETAILS")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var getCart CART_DETAILS
		err := results.Scan(&getCart.CARTID, &getCart.ITEM_NAME, &getCart.Quantity, &getCart.TOTAL_AMOUNT, &getCart.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CART = append(CART, getCart)
	}
	json.NewEncoder(w).Encode(CART)
}

func getCartById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getCartbyID CART_DETAILS
	results, err := db.Query("SELECT CARTID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS FROM CARTDETAILS where CARTID=?", getCartbyID.CARTID)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getCartbyID.CARTID, &getCartbyID.ITEM_NAME, &getCartbyID.Quantity, &getCartbyID.TOTAL_AMOUNT, &getCartbyID.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CART = append(CART, getCartbyID)
	}
	json.NewEncoder(w).Encode(CART)
}

// ● GET /api/cart/getById/{cartId}: This API is used to retrieve cart record by cartId.
// It takes the cartId in the URL parameter and returns the ResponseDTO of specific cart details.
func getCartItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getCartId CART_DETAILS
	results, err := db.Query("SELECT CARTID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS FROM CARTDETAILS where CARTID=?", getCartId.CARTID)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getCartId.CARTID, &getCartId.ITEM_NAME, &getCartId.Quantity, &getCartId.TOTAL_AMOUNT, &getCartId.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		CART = append(CART, getCartId)
	}
	json.NewEncoder(w).Encode(CART)
}

// ● POST /api/cart/create: This API is used to insert items in the cart.
// It accepts the token and CartDTO in the request body and returns the ResponseDTO of inserted item in the cart.
func addPostCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postCart CART_DETAILS
	json.NewDecoder(r.Body).Decode(&postCart)
	//read from the request
	// Execute the query
	_, err := db.Query("INSERT INTO CARTDETAILS (CARTID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS) VALUES (?,?,?,?,?)")
	//.Scan(&postOrder.ORDERID, &postOrder.ORDER_NAME, &postOrder.Quantity, &postOrder.STATUS)
	if err != nil {
		panic(err.Error())
	}

	//ORDER = append(ORDER, postOrder)
	//json.NewEncoder(w).Encode(ORDER) //optional
	w.Write([]byte("Data added successfully..."))
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
	ress, err := db.Exec("Update CARTDETAILS set ITEM_NAME=?,ADDRESS=?,AMOUNT=?,STATUS=? where ORDERID=?)", cart_id)
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
	cart_id, err := strconv.Atoi(r.URL.Path[len("/order/{id}"):])
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
	cart_id, err := strconv.Atoi(r.URL.Path[len("/cart"):])
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
