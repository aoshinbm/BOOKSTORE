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
	ORDERID    int    `json:id`
	ORDER_NAME string `json:order_name`
	Quantity   int    `json:quantity`
	STATUS     string `json:status`
}

var ORDER []ORDER_DETAILS
var db *sql.DB
var err error

func main() {
	// Open up our database connection.
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/BOOKSTORE")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	or := ORDER_DETAILS{ORDERID: 1, ORDER_NAME: "Harry Potter", Quantity: 2, STATUS: "Confirmed"}
	ORDER = append(ORDER, or)

	http.HandleFunc("/orders", handleOrderDetails)
	http.HandleFunc("/order/{id}", handleOrderDetails)

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleOrderDetails(w http.ResponseWriter, r *http.Request) {
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
func getOrderList(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	results, err := db.Query("SELECT ORDERID ,ORDER_NAME ,Quantity ,STATUS FROM OrderDetails")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var getOrder ORDER_DETAILS
		// for each row, scan the result into our tag composite object
		err := results.Scan(&getOrder.ORDERID, &getOrder.ORDER_NAME, &getOrder.Quantity, &getOrder.STATUS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		ORDER = append(ORDER, getOrder)
	}
	json.NewEncoder(w).Encode(ORDER)
}
func addPostOrder(w http.ResponseWriter, r *http.Request) {
	var postOrder ORDER_DETAILS
	json.NewDecoder(r.Body).Decode(&postOrder)
	//read from the request
	// Execute the query
	_, err := db.Query("INSERT INTO OrderDetails (ORDERID ,ORDER_NAME ,Quantity ,STATUS) VALUES (?,?,?,?)")
	//.Scan(&postOrder.ORDERID, &postOrder.ORDER_NAME, &postOrder.Quantity, &postOrder.STATUS)
	if err != nil {
		panic(err.Error())
	}

	//ORDER = append(ORDER, postOrder)
	//json.NewEncoder(w).Encode(ORDER) //optional
	w.Write([]byte("Data added successfully..."))
}
func removeOrder(w http.ResponseWriter, r *http.Request) {
	for i, _ := range ORDER {
		ORDER = append(ORDER[:i], ORDER[i+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(ORDER)
}

func updatePutOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/orders/"):])
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r.URL.Path[len("/orders/"):])
	path := "/order/1"
	orderId := path[8:]
	fmt.Println(orderId)

	for i, data := range ORDER {
		if data.ORDERID == id {
			var newOrder ORDER_DETAILS
			json.NewDecoder(r.Body).Decode(&newOrder)
			newOrder.ORDERID = id
			ORDER[i] = newOrder
			json.NewEncoder(w).Encode(newOrder)
		}
	}
}
