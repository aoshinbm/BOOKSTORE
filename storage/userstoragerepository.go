package repository

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"example.com/BookStoreProj/types"
)

type UserStoreRepository struct {
	DB *sql.DB
}

// funcs to handle queries to database
func (r *UserStoreRepository) GetUsersFullList() ([]*types.USERREGIST, error) {
	fmt.Println("Get All Users")
	results, err := r.DB.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration")
	if err != nil {
		panic(err.Error())
	}
	var USER []*types.USERREGIST
	for results.Next() {
		ress := &types.USERREGIST{}
		err := results.Scan(&ress.USERID, &ress.Username, &ress.FirstName, &ress.LastName, &ress.Password, &ress.Email)
		if err != nil {
			return nil, err
		}
		USER = append(USER, ress)
	}
	return USER, nil
}
func (r *UserStoreRepository) GetUserById() ([]*types.USERREGIST, error) {
	fmt.Println("Get a User details by Id")
	result, err := r.DB.Query("SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration where ID=?", user_Id)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		gettUser := &types.USERREGIST{}
		err := result.Scan(&gettUser.USERID, &gettUser.Username, &gettUser.FirstName, &gettUser.LastName, &gettUser.Password, &gettUser.Email)
		if err != nil {
			panic(err.Error())
			return nil, err
		}
	}
	return gettUser, nil
}
func (r *UserStoreRepository) AddPostUsers() ([]*types.USERREGIST, error) {
	fmt.Println("Add USER")
	post := &types.USERREGIST{}
	_, err := r.DB.Query("INSERT INTO UserRegistration (Id, Username, FirstName, LastName, Password, Email) VALUES (?,?,?,?,?,?)", post.USERID, post.Username, post.FirstName, post.LastName, post.Password, post.Email)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
}
func (r *UserStoreRepository) RemoveUser() ([]*types.USERREGIST, error) {
	fmt.Println("Remove User is getting called")
	id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/deleteUser/"):])
	if err != nil {
		panic(err.Error())
	}
	deleteID, err := r.DB.Query("delete from UserRegistration where Id=?", id)
	if err != nil {
		return nil, err
	}
	return deleteID, nil
}
func (r *UserStoreRepository) UpdatePutUsers() ([]*types.USERREGIST, error) {
	fmt.Println("Update User")
	id, err := strconv.Atoi(r.URL.Path[len("/api/userservices/updateUser/"):])
	if err != nil {
		panic(err.Error())
	}

	res, err := r.DB.Query("Update UserRegistration set Id=?, Username=?, FirstName=?, LastName=?, Password=?, Email=? where id=?", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *UserStoreRepository) handlePostLogin() ([]*types.USERREGIST, error) {
	userlogin := []*types.Login{}
	userregistration := []*types.USERREGIST{}
	results, err := r.DB.Query("SELECT * FROM UserRegistration where Username=? and Password=?", userlogin.Username, userlogin.Password)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		err := results.Scan(&userregistration.USERID, &userregistration.FirstName, &userregistration.LastName, &userregistration.Email, &userregistration.Username, &userregistration.Password)
		if err != nil {
			http.Error(w, "USER doesnt Exist", http.StatusInternalServerError)
			return nil, err
		}
	}
	return res, nil
}
