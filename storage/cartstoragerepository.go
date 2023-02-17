package repository

import (
	"database/sql"
	"fmt"

	"example.com/BookStoreProj/types"
)

type CartStoreRepository struct {
	DB *sql.DB
}

// funcs to handle queries to database
func (r *CartStoreRepository) GetCartList() ([]*types.CART_DETAILS, error) {
	fmt.Println("Get ALL is calling")
	results, err := r.DB.Query("SELECT * FROM CARTDETAILS")
	if err != nil {
		panic(err.Error())
	}
	//store
	var CART []*types.CART_DETAILS
	for results.Next() {
		//fetch data
		getCart := &types.CART_DETAILS{}
		err := results.Scan(&getCart.CARTID, &getCart.ITEM_NAME, &getCart.Quantity, &getCart.TOTAL_AMOUNT)
		if err != nil {
			return nil, err
		}
		CART = append(CART, getCart)
	}
	return CART, nil
}

func (r *CartStoreRepository) GetCartItemById() ([]*types.CART_DETAILS, error) {
	fmt.Println("Get By ID is calling")
	getCartId := &types.CART_DETAILS{}
	results, err := r.DB.Query("SELECT * FROM CARTDETAILS where CARTID=?", cart_id)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		err := results.Scan(&getCartId.CARTID, &getCartId.ITEM_NAME, &getCartId.Quantity, &getCartId.TOTAL_AMOUNT)
		if err != nil {
			return nil, err
		}
	}
	return getCartId, nil
}
