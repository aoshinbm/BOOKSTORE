package constant

const (
	GetBooks      = "SELECT BOOKID ,Title ,Author ,Year FROM BookRecord"
	GetBookByName = "SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where Title=?"
	GetBookById   = "SELECT BOOKID ,Title ,Author ,Year FROM BookRecord where BOOKID=? "
	AddBook       = "INSERT INTO BookRecord (BOOKID ,Title ,Author ,Year) VALUES (?,?,?,?)"
	RemoveBook    = "delete from BookRecord where BOOKID=?"
	UpdateBook    = "Update BookRecord set BOOKID=?,Title=? ,Author=? ,Year=? where BOOKID=?"
)
