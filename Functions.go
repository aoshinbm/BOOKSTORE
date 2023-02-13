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

func User_Routes() {
	http.HandleFunc("/api/userservices/getAllUsers", getUsersFullList)
	http.HandleFunc("/api/userservices/addUser", addPostUsers)
	http.HandleFunc("/api/userservices/getUser/{id}", getUser)
	http.HandleFunc("/api/userservices/updateUser/{id}", updatePutUsers)
	http.HandleFunc("/api/userservices/deleteUser/{id}", removeUser)
	//login verification
	http.HandleFunc("/api/userservice/login", handlePostLogin)
}

func Book_Routes() {
	http.HandleFunc("/getBooks", getBooksFullList)
	http.HandleFunc("/getBookByName/{bookName}", getBookByName)
	http.HandleFunc("/getBook/{bookId}", getBook)
	http.HandleFunc("/addBook", addPostBooks)
	http.HandleFunc("/update/{bookId}", updatePutBooks)
	http.HandleFunc("/delete/{bookId}", removeBook)
}

func Order_Routes() {
	http.HandleFunc("/api/order/retrieveAllOrders", getOrderList)
	http.HandleFunc("/api/order/retrieveOrder/{id}", getOrderById)
	http.HandleFunc("/api/order/cancelOrder/{id}", removeOrder)
	http.HandleFunc("/api/order/insert", addPostOrder)
	http.HandleFunc("/api/order/update", updatePutOrder)
}

func Cart_Routes() {
	http.HandleFunc("/api/cart/getAll", getCartList)
	http.HandleFunc("/api/cart/getById/{cartId}", getCartItemById)
	http.HandleFunc("/api/cart/create", addPostCart)
	http.HandleFunc("/api/cart/updateById/{cartId}", updatePutCart)
	http.HandleFunc("/api/cart/delete/{cartId}", removeCartItemById)
	http.HandleFunc("/api/cart/deleteall", removeCart)
}
