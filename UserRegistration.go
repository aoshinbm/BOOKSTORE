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
	USERID    int    `json:"userid"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"emailid"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	http.HandleFunc("/api/userservices/getAllUsers", getUsersFullList)
	http.HandleFunc("/api/userservices/addUser", addPostUsers)
	http.HandleFunc("/api/userservices/getUser/", getUserById)
	http.HandleFunc("/api/userservices/updateUser/", updatePutUsers)
	http.HandleFunc("/api/userservices/deleteUser/", removeUser)
	//login verification
	http.HandleFunc("/api/userservice/login", handlePostLogin)
	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

/* UserRegistrationController
● GET /api/userservice/getAll/{token}: retrieves all user records using the token passed as
a parameter
● GET /api/userservice/verify/{token}: verifies the given token
*/

// ● GET /api/userservice/getAllUsers: retrieves all user records
func getUsersFullList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Users")
	w.Header().Set("Content-Type", "application/json")
	results, err := db.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var ress USERREGIST
		err := results.Scan(&ress.USERID, &ress.Username, &ress.FirstName, &ress.LastName, &ress.Password, &ress.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		USER = append(USER, ress)
	}
	json.NewEncoder(w).Encode(USER)
}

// ● GET /api/userservice/getUser/{userId}: retrieves a single user record with the given userId
func getUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a User details by Id")
	w.Header().Set("Content-Type", "application/json")
	user_Id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/getUser/"):])
	var gettUser USERREGIST
	result, err := db.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration where ID=?", user_Id)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&gettUser.USERID, &gettUser.Username, &gettUser.FirstName, &gettUser.LastName, &gettUser.Password, &gettUser.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(w).Encode(gettUser)
}

// ● POST /api/userservice/register: creates a new user record by registering the user, using the UserRegistrationDTO request body
func addPostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add USER")
	var post USERREGIST
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&post)
	_, err := db.Query("INSERT INTO UserRegistration (Id, Username, FirstName, LastName, Password, Email) VALUES (?,?,?,?,?,?)", post.USERID, post.Username, post.FirstName, post.LastName, post.Password, post.Email)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("Data added successfully..."))
}

// ● DELETE /api/userservice/delete/{userId}: deletes a user record with the given userId
func removeUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remove User is getting called")
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/deleteUser/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}

	deleteID, err := db.Query("delete from UserRegistration where Id=?", id)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(deleteID)
	w.Write([]byte("USER DELETED successfully..."))
}

// ● PUT /api/userservice/update/{userId}: updates a user record with the given userId, using the UserRegistrationDTO request body
func updatePutUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Update User")
	id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/updateUser/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}

	res, err := db.Query("Update UserRegistration set Id=?, Username=?, FirstName=?, LastName=?, Password=?, Email=? where id=?)", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(res)
	w.Write([]byte("USER UPDATED successfully..."))
}

// ● POST /api/userservice/login: logs in the user, using the LoginDTO request body
func handlePostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userlogin Login
	var userregistration USERREGIST
	results, err := db.Query("SELECT * FROM UserRegistration where Username=? and Password=?)", userlogin.Username, userlogin.Password)
	if err != nil {
		http.Error(w, "Error ", http.StatusBadRequest)
		return
	}
	for results.Next() {
		err := results.Scan(&userregistration.USERID, &userregistration.FirstName, &userregistration.LastName, &userregistration.Email, &userregistration.Username, &userregistration.Password)
		if err != nil {
			http.Error(w, "USER doesnt Exist", http.StatusInternalServerError)
			return
		}
	}
}
