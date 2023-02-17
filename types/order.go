package types

// model of order
// models will specify the attributes of the order record that will be added to the order database
type ORDER_DETAILS struct {
	ORDERID   int    `json:orderid`
	ITEM_NAME string `json:item_name`
	ADDRESS   string `json:address`
	AMOUNT    string `json:amount`
	STATUS    string `json:status`
}
