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

	//http.HandleFunc("/api/userservices", handleUserRegistration)
	http.HandleFunc("/api/userservices/getAllUsers", getUsersFullList)
	http.HandleFunc("/api/userservices/addUser", addPostUsers)
	http.HandleFunc("/api/userservices/getUser/{id}", getUser)
	http.HandleFunc("/api/userservices/updateUser/{id}", updatePutUsers)
	http.HandleFunc("/api/userservices/deleteUser/{id}", removeUser)
	//login verification
	http.HandleFunc("/api/userservice/login", handlePostLogin)

	//User_Routes()

	//start the server
	fmt.Println("Starting server on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

/*
// 2 parameters :- response is interface, request is struct
func handleUserRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle user regestation is getting called")
	//r.method returns which method is the request calling
	switch r.URL.Path {
	case "/api/userservices":
		//● GET /api/userservice: retrieves all user records
		getUsersFullList(w, r)
		//● POST /api/userservice/register: creates a new user record by registering the user, using the UserRegistrationDTO request body
		addPostUsers(w, r)
	case "/api/userservices/{id}":
		//● GET /api/userservice/get/{userId}: retrieves a single user record with the given userId
		getUser(w, r)
		//● PUT /api/userservice/update/{userId}: updates a user record with the given userId, using the UserRegistrationDTO request body
		updatePutUsers(w, r)
		//● DELETE /api/userservice/delete/{userId}: deletes a user record with the given userId
		removeUser(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
*/

/* UserRegistrationController
● GET /api/userservice/getAll/{token}: retrieves all user records using the token passed as
a parameter
● GET /api/userservice/verify/{token}: verifies the given token
*/

// ● GET /api/userservice/getAllUsers: retrieves all user records
func getUsersFullList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Execute the query
	results, err := db.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var ress USERREGIST
		// for each row, scan the result into our tag composite object
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
func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a User details")
	w.Header().Set("Content-Type", "application/json")
	user_Id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/getUser/"):])
	var gettUser USERREGIST
	result, err := db.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration where USERID=?", user_Id)
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
	var post USERREGIST
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&post)
	//read from the request
	_, err := db.Query("INSERT INTO UserRegistration (Id, Username, FirstName, LastName, Password, Email) VALUES (?,?,?,?,?,?)", post.USERID, post.Username, post.FirstName, post.LastName, post.Password, post.Email)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("Data added successfully..."))
}

// ● DELETE /api/userservice/delete/{userId}: deletes a user record with the given userId
func removeUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remove is getting called")
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/deleteUser/"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}

	deleteID, err := db.Exec("delete from UserRegistration where Id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result, err := deleteID.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	print(result)
	json.NewEncoder(w).Encode(USER)
	w.Write([]byte("Data DELETED successfully..."))
}

// ● PUT /api/userservice/update/{userId}: updates a user record with the given userId, using the UserRegistrationDTO request body
func updatePutUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Path[len("/users/{id}"):])
	if err != nil {
		http.Error(w, "Invalid Request ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("Update UserRegistration set Id=?, Username=?, FirstName=?, LastName=?, Password=?, Email=? where id=?)", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	updated, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	print(updated)
	json.NewEncoder(w).Encode(USER)
	w.Write([]byte("Data UPDATED successfully..."))
}

/*
You will get the post request on login api endpoint
then you will check if user exist in database with given username and password
and then get the whole row data with select * from

and using that data create a stuct of user regeration

and then take userstruct.id as a parameter to CreateToken method from token.go
it will return a long string
you have to return it to
postman

username password email etc

you have to map it to the struct of useregration
you will create an ohject from sql data.

with w.Write method
*/

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
