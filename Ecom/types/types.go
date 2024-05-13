package types

import (
	"time"
)

type UserStore struct {
	GetUserByEmail(email string) (*User, error)
	
}

type User struct {
	ID       	 int    	`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName  	string 		`json:"lastName"`
	Email     	string 		`json:"email"`
	Password  	string 		`json:"_"`
	CreatedAt 	time.Time   `json:"createdAt"`
}
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
