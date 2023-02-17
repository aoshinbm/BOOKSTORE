package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"example.com/BookStoreProj/types"
)

type OrderStoreRepository struct {
	DB *sql.DB
}

// funcs to handle queries to database
func (r *OrderStoreRepository) GetOrders() ([]*types.ORDER_DETAILS, error) {
	fmt.Println("Get All Orders")
	results, err := r.DB.Query("SELECT * FROM OrderDetails")
	if err != nil {
		panic(err.Error())
	}
	var ORDER []*types.ORDER_DETAILS
	for results.Next() {
		getOrder := &types.ORDER_DETAILS{}
		err := results.Scan(&getOrder.ORDERID, &getOrder.ITEM_NAME, &getOrder.ADDRESS, &getOrder.AMOUNT, &getOrder.STATUS)
		if err != nil {
			return nil, err
		}
		ORDER = append(ORDER, getOrder)
	}
	return ORDER, nil
}

func (r *OrderStoreRepository) GetOrderByIdd() ([]*types.ORDER_DETAILS, error) {
	fmt.Println("Get a Order by Id")
	results, err := r.DB.Query("SELECT * FROM OrderDetails where ORDERID=?", orderID)
	if err != nil {
		panic(err.Error())
	}
	getOrderId := &types.ORDER_DETAILS{}
	for results.Next() {
		err := results.Scan(&getOrderId.ORDERID, &getOrderId.ITEM_NAME, &getOrderId.ADDRESS, &getOrderId.AMOUNT, &getOrderId.STATUS)
		if err != nil {
			return nil, err
		}
	}
	return getOrderId, nil
}

func (r *OrderStoreRepository) AddPostOrder() ([]*types.ORDER_DETAILS, error) {
	fmt.Println("Add Order")
	_, err := r.DB.Query("INSERT INTO orderdetails (ORDERID ,ITEMNAME ,ADDRESS,AMOUNT,STATUS) VALUES (?,?,?,?,?)", postOrder.ORDERID, postOrder.ITEM_NAME, postOrder.ADDRESS, postOrder.AMOUNT, postOrder.STATUS)
	if err != nil {
		panic(err.Error())
	}
}
func (r *OrderStoreRepository) RemoveOrder() ([]*types.ORDER_DETAILS, error) {
	fmt.Println("Remove Order")
	order_id, err := strconv.Atoi(r.URL.Path[len("/api/order/cancelOrder/"):])
	if err != nil {
		return
	}
	deleteOrderID, err := db.Query("delete from ORDERDETAILS where ORDERID=?", order_id)
	if err != nil {
		panic(err)
	}
}
func (r *OrderStoreRepository) UpdatePutOrder() ([]*types.ORDER_DETAILS, error) {
	fmt.Println("Update order")
	order_id, err := strconv.Atoi(r.URL.Path[len("/api/order/update/"):])
	if err != nil {
		return
	}
	ress, err := db.Query("Update ORDERDETAILS set ORDERID=?,ITEM_NAME=?,ADDRESS=?,AMOUNT=?,STATUS=? where ORDERID=?", order_id)
	if err != nil {
		panic(err)
	}
}
