package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

func MockUserStore() {}

func GetuserByEmail(email string) (*User, error) {
	return nil, nil
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"createdAt"`
}



type ProductStore interface {
	GetProducts() ([]Product,error)
	GetUserByID(ps []int) ([]Product,error)
	UpdateProduct(Product)error

}

type OrderStore interface {
	CreateOrder(Order) (int,error)
	CreateOrderItem(OrderItem) error
}

type Order struct {
	ID		 	 int	    `json:"id"`
	UserID	 	 int	    `json:"userid"`
	Total	     float64	`json:"total"`
	Status		 string		`json:"status"`
	Address		 string	    `json:"address"`
	CreatedAt	 time.Time	`json:"createdat"`

}

type Product struct {
	ID				int 		`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Image			string		`json:"image"`
	Price			float64		`json:"price"`
	Quantity		int			`json:"quantity"`
	CreatedAt		time.Time	`json:"createdat"`

}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"  validate:"required"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required,"`
}

type CartItem struct {
	ProductID int 	`json:"productid"`
	Quantity  int	`json:"quantity"`
}


type CartCheckoutPayload struct {
	Items  []CartItem `json:"items"  validate:"required"` 	
}


