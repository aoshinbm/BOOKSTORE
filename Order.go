package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// model of order
// models will specify the attributes of the order record that will be added to the order database
type ORDER_DETAILS struct {
	ORDERID   int    `json:id`
	ITEM_NAME string `json:order_name`
	ADDRESS   string `json:address`
	AMOUNT    string `json:amount`
	STATUS    string `json:status`
}

var ORDER []ORDER_DETAILS
var db *sql.DB
var err error

func main() {
	DatabaseConnection()

	or := ORDER_DETAILS{ORDERID: 1, ITEM_NAME: "Harry Potter", ADDRESS: "Goregaon,Mumbai", AMOUNT: "5000", Quantity: 2, STATUS: "Confirmed"}
	ORDER = append(ORDER, or)

	http.HandleFunc("/orders", handleOrderDetails)
	http.HandleFunc("/order/{id}", handleOrderDetails)

	//start the server
	StartServer()
}

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
}

// ● GET /api/order/retrieveAllOrders - to retrieve all order records
func getOrderList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Execute the query
	results, err := db.Query("SELECT ORDERID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS FROM OrderDetails")
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
	results, err := db.Query("SELECT ORDERID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS FROM OrderDetails where ORDERID=?", getOrderId.ORDERID)
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
	_, err := db.Query("INSERT INTO OrderDetails (ORDERID ,ITEM_NAME ,ADDRESS,AMOUNT,STATUS) VALUES (?,?,?,?,?)")
	//.Scan(&postOrder.ORDERID, &postOrder.ORDER_NAME, &postOrder.Quantity, &postOrder.STATUS)
	if err != nil {
		panic(err.Error())
	}

	//ORDER = append(ORDER, postOrder)
	//json.NewEncoder(w).Encode(ORDER) //optional
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
	order_id, err := strconv.Atoi(r.URL.Path[len("/order/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}
	deleteOrderID, err := db.Exec("delete from ORDERDETAILS where BOOKID=?", order_id)
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
