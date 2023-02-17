package constant

const (
	GetUsersFullList = "SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration"
	GetUserById      = "SELECT Id ,Username ,FirstName ,LastName ,Password ,Email FROM userregistration where ID=?"
	AddPostUsers     = "INSERT INTO UserRegistration (Id, Username, FirstName, LastName, Password, Email) VALUES (?,?,?,?,?,?)"
	RemoveUser       = "delete from UserRegistration where Id=?"
	UpdatePutUsers   = "Update UserRegistration set Id=?, Username=?, FirstName=?, LastName=?, Password=?, Email=? where id=?"
	handlePostLogin  = "SELECT * FROM UserRegistration where Username=? and Password=?"
)
