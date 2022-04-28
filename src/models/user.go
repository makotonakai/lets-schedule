package models

type User struct {	
	Id 						int			`json:"id" param:"id"`	
	UserName 			string 	`json:"username"`
	EmailAddress  string 	`json:"email_address"`
	Password 			string 	`json:"password"`
	IsAdmin 			bool 		`json:"is_admin"`
	CanLogin 			bool 		`json:"can_login"`				
}