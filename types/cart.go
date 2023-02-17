package types

type CART_DETAILS struct {
	CARTID       int    `json:cartid`
	ITEM_NAME    string `json:item_name`
	Quantity     int    `json:quantity`
	TOTAL_AMOUNT int    `json:total_amount`
}
