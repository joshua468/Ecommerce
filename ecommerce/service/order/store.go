package cart

import(
	"database/sql"
)

type Store struct {
db *sql.DB
}

type NewStore(db *sql.DB)  *Store{
	return &Store{db:db}
}

func (s *Store) CreateOrder(order  types.Order)  (int,error) {
	res,err := s.db.Exec("INSERT INTO orders (userId,total,status,
		address) VALUES(?,?,?,?)",order.UserID,order.Total,Order.Status,order.Address)
		if err!= nil {
			return 0,nil
		}
		id,err:= res.LastInsertId()
		if err!= nil {
			return 0,nil
		}
		return int(id),nil
}

func  (s *Store) CreateOrderItem(OrderItem types.OrderItem) error {
	_,err:= s.db.Exec("INSERT INTO  order_items(orderId,productId,quantity,price) VALUES(?,?,?,?)",orderItem.orderId,orderItem.productId,orderItem.OrderID,orderItem.ProductID,orderItem.ProductID,orderItem.Quantity,orderItem.Price)
	return err
}