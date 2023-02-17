package constant

const (
	GetOrders      = "SELECT * FROM OrderDetails"
	GetOrderByIdd  = "SELECT * FROM OrderDetails where ORDERID=?"
	AddPostOrder   = "INSERT INTO orderdetails (ORDERID ,ITEMNAME ,ADDRESS,AMOUNT,STATUS) VALUES (?,?,?,?,?)"
	RemoveOrder    = "delete from ORDERDETAILS where ORDERID=?"
	UpdatePutOrder = "Update ORDERDETAILS set ORDERID=?,ITEM_NAME=?,ADDRESS=?,AMOUNT=?,STATUS=? where ORDERID=?"
)
