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

// model of order
// models will specify the attributes of the order record that will be added to the order database
type ORDER_DETAILS struct {
	ORDERID   int    `json:orderid`
	ITEM_NAME string `json:item_name`
	ADDRESS   string `json:address`
	AMOUNT    string `json:amount`
	STATUS    string `json:status`
}

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
	http.HandleFunc("/api/order/retrieveOrder/", getOrderById)
	http.HandleFunc("/api/order/cancelOrder/", removeOrder)
	http.HandleFunc("/api/order/insert", addPostOrder)
	http.HandleFunc("/api/order/update/", updatePutOrder)

	//Order_Routes()

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// ● GET /api/order/retrieveAllOrders - to retrieve all order records
func getOrderList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Orders")
	w.Header().Set("Content-Type", "application/json")
	// Execute the query
	results, err := db.Query("SELECT * FROM OrderDetails")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var getOrder ORDER_DETAILS
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
	fmt.Println("Get a Order by Id")
	w.Header().Set("Content-Type", "application/json")
	//bookId, err := strconv.Atoi(r.URL.Query().Get("bookId"))
	orderID, err := strconv.Atoi(r.URL.Path[len("/api/order/retrieveOrder/"):])
	results, err := db.Query("SELECT * FROM OrderDetails where ORDERID=?", orderID)
	if err != nil {
		panic(err.Error())
	}
	var getOrderId ORDER_DETAILS
	for results.Next() {
		err := results.Scan(&getOrderId.ORDERID, &getOrderId.ITEM_NAME, &getOrderId.ADDRESS, &getOrderId.AMOUNT, &getOrderId.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(w).Encode(getOrderId)
}

// ● POST /api/order/insert - to insert a new order record
func addPostOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Order")
	w.Header().Set("Content-Type", "application/json")
	var postOrder ORDER_DETAILS
	json.NewDecoder(r.Body).Decode(&postOrder)
	_, err := db.Query("INSERT INTO orderdetails (ORDERID ,ITEMNAME ,ADDRESS,AMOUNT,STATUS) VALUES (?,?,?,?,?)", postOrder.ORDERID, postOrder.ITEM_NAME, postOrder.ADDRESS, postOrder.AMOUNT, postOrder.STATUS)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("Data added successfully..."))
}

// ● DELETE /api/order/cancelOrder/{id} - to cancel an order record by its id
func removeOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remove Order")
	w.Header().Set("Content-Type", "application/json")
	order_id, err := strconv.Atoi(r.URL.Path[len("/api/order/cancelOrder/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteOrderID, err := db.Query("delete from ORDERDETAILS where ORDERID=?", order_id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(deleteOrderID)
	w.Write([]byte("Order DELETED successfully..."))
}

// ● PUT /api/order/update/{id} - to cancel an order record by its id
func updatePutOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update order")
	w.Header().Set("Content-Type", "application/json")
	order_id, err := strconv.Atoi(r.URL.Path[len("/api/order/update/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	ress, err := db.Query("Update ORDERDETAILS set ORDERID=?,ITEM_NAME=?,ADDRESS=?,AMOUNT=?,STATUS=? where ORDERID=?)", order_id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(ress)
	w.Write([]byte("Order UPDATED successfully..."))
}
