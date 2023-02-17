package types

// model of book
// models will specify the attributes of the book record that will be added to the books database
type BOOK_RECORD struct {
	BOOKID int    `json:bookid`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}
