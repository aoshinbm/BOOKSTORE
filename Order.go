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

var ORDER []ORDER_DETAILS
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

	or := ORDER_DETAILS{ORDERID: 1, ITEM_NAME: "Harry Potter", ADDRESS: "Goregaon,Mumbai", AMOUNT: "5000", STATUS: "Confirmed"}
	ORDER = append(ORDER, or)

	http.HandleFunc("/api/order/retrieveAllOrders", getOrderList)
	http.HandleFunc("/api/order/retrieveOrder/{id}", getOrderById)
	http.HandleFunc("/api/order/cancelOrder/{id}", removeOrder)
	http.HandleFunc("/api/order/insert", addPostOrder)
	http.HandleFunc("/api/order/update", updatePutOrder)

	//Order_Routes()

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

/*
func handleOrderDetails(w http.ResponseWriter, r *http.Request) {
	//OrderController
	//r.method returns which method is the request calling
	switch r.Method {
	case "GET":
		getOrderList(w, r)
	case "PUT":
		updatePutOrder(w, r)
	case "POST":
		addPostOrder(w, r)
	case "DELETE":
		removeOrder(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}*/

// ● GET /api/order/retrieveAllOrders - to retrieve all order records
func getOrderList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Execute the query
	results, err := db.Query("SELECT * FROM OrderDetails")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var getOrder ORDER_DETAILS
		// for each row, scan the result into our tag composite object
		err := results.Scan(&getOrder.ORDERID, &getOrder.ITEM_NAME, &getOrder.ADDRESS, &getOrder.AMOUNT, &getOrder.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ORDER = append(ORDER, getOrder)
	}
	json.NewEncoder(w).Encode(ORDER)
}

// ● GET /api/order/retrieveOrder/{id} - to retrieve an order record by its id
func getOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getOrderId ORDER_DETAILS
	results, err := db.Query("SELECT * FROM OrderDetails where ORDERID=?", getOrderId.ORDERID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err := results.Scan(&getOrderId.ORDERID, &getOrderId.ITEM_NAME, &getOrderId.ADDRESS, &getOrderId.AMOUNT, &getOrderId.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		ORDER = append(ORDER, getOrderId)
	}
	json.NewEncoder(w).Encode(ORDER)
}

// ● POST /api/order/insert - to insert a new order record
func addPostOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postOrder ORDER_DETAILS
	json.NewDecoder(r.Body).Decode(&postOrder)
	//read from the request
	// Execute the query
	_, err := db.Query("INSERT INTO orderdetails (ORDERID ,ITEMNAME ,ADDRESS,AMOUNT,STATUS) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("Data added successfully..."))
}

// ● PUT /api/order/cancelOrder/{id} - to cancel an order record by its id
func updatePutOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	order_id, err := strconv.Atoi(r.URL.Path[len("/orders/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	ress, err := db.Exec("Update ORDERDETAILS set ORDERID=?,ITEM_NAME=?,ADDRESS=?,AMOUNT=?,STATUS=? where ORDERID=?)", order_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	orderUpdated, err := ress.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(orderUpdated)
	w.Write([]byte("Order UPDATED successfully..."))
}

func removeOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	order_id, err := strconv.Atoi(r.URL.Path[len("/api/order/cancelOrder/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteOrderID, err := db.Exec("delete from ORDERDETAILS where ORDERID=?", order_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	del_Order, err := deleteOrderID.RowsAffected()
	if err != nil {
		panic(err)
	}
	print(del_Order)
	w.Write([]byte("Order DELETED successfully..."))
}
