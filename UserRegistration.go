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

type USERREGIST struct {
	USERID    int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"emailid"`
}

var db *sql.DB
var USER []USERREGIST
var err error

func main() {
	// Open up our database connection.
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstore")
	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	u := USERREGIST{USERID: 1, Username: "aoshub", FirstName: "Aoshin", LastName: "Manjuran", Password: "#7Aoshinb", Email: "aoshuthanatos@gmail.com"}
	USER = append(USER, u)

	u = USERREGIST{USERID: 2, Username: "bobby", FirstName: "Jenny", LastName: "Lawson", Password: "Jen@Law", Email: "jenL@gmail.com"}
	USER = append(USER, u)

	u = USERREGIST{USERID: 3, Username: "mandrake", FirstName: "Subin", LastName: "Manjuran", Password: "Khaldrogo!663", Email: "subin@gmail.com"}
	USER = append(USER, u)

	//http.HandleFunc("/", homePage)
	http.HandleFunc("/users", handleUserRegistration)
	http.HandleFunc("/users/{id}", handleUserRegistration)

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// 2 parameters :- response is interface, request is struct
func handleUserRegistration(w http.ResponseWriter, r *http.Request) {
	//r.method returns which method is the request calling
	/*
			UserRegistrationController
		● GET /api/userservice/get/{userId}: retrieves a single user record with the given userId
		● GET /api/userservice/getAll/{token}: retrieves all user records using the token passed as
		a parameter
		● GET /api/userservice/verify/{token}: verifies the given token
		● POST /api/userservice/login: logs in the user, using the LoginDTO request body
	*/
	switch r.Method {
	case "GET":
		//● GET /api/userservice: retrieves all user records
		getUsersFullList(w, r)
	case "POST":
		//● POST /api/userservice/register: creates a new user record by registering the user,
		//using the UserRegistrationDTO request body
		addPostUsers(w, r)
	//case "PUT":
	//● PUT /api/userservice/update/{userId}: updates a user record with the given userId,
	// using the UserRegistrationDTO request body
	//updatePostUsers(w, r)
	case "DELETE":
		//● DELETE /api/userservice/delete/{userId}: deletes a user record with the given userId
		removeUser(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getUsersFullList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Execute the query
	results, err := db.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM UserRegistration")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var ress USERREGIST
		// for each row, scan the result into our tag composite object
		err := results.Scan(&ress.USERID, &ress.Username, &ress.FirstName, &ress.LastName, &ress.Password, &ress.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // proper error handling instead of panic in your app
			return
		}
		USER = append(USER, ress)
	}
	json.NewEncoder(w).Encode(USER)
}
func addPostUsers(w http.ResponseWriter, r *http.Request) {
	var post USERREGIST
	json.NewDecoder(r.Body).Decode(&post)
	//read from the request
	// Execute the query
	_, err := db.Query("INSERT INTO UserRegistration (Id, Username, FirstName, LastName, Password, Email) VALUES (?,?,?,?,?,?)", post.USERID, post.Username, post.FirstName, post.LastName, post.Password, post.Email)
	//.Scan(&post.USERID,&post.Username, &post.FirstName, &post.LastName, &post.Password, &post.Email)
	if err != nil {
		panic(err.Error())
	}

	//USER = append(USER, post)
	//json.NewEncoder(w).Encode(USER) //optional
	w.Write([]byte("Data added successfully..."))
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	for i, _ := range USER {
		USER = append(USER[:i], USER[i+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}
}

func updatePutUsers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r.URL.Path[len("/users/"):])
	path := "/users/1"
	userId := path[8:]
	fmt.Println(userId)

	for i, userdata := range USER {
		if userdata.USERID == id {
			var newUserr USERREGIST
			json.NewDecoder(r.Body).Decode(&newUserr)
			newUserr.USERID = id
			USER[i] = newUserr
			json.NewEncoder(w).Encode(newUserr)
		}
	}
}
