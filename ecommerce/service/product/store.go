package store

import (
	"database/sql"
	"strings"

	"github.com/joshua468/ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}
func (h *Store) GetProductsByIDs(productsIDs []int) ([]types.Product,error) {
	placeholders := strings.Repeat(",?",len(productsIDs)-1)
	query:= fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)",placeholders)

	args:= make([]interface{},len(productsIDs))
	for i,v := range productsIDs {
		args[i] = v
	}
	rows,err := s.db.Query(query,args...)
	if err!= nil {
		return nil,err
	}
	products  := append(products, *p)

}

func (s *Store) GetProducts() ([]types.Product,error){
rows,err := s.db.Query("SELECT * FROM products")
if err!= nil {
return nil,err
}

products := make([]types.Product,0)
for rows.Next() {
	p,err := ScanRowsIntoProduct(rows)
	if err!= nil {
		return nil,err
	}
	products = append(products, p)
}
return products,nil
}

func (s *Store)  UpdateProduct(product types.Product) error {
_,err := s.db.Exec("UPDATE products SET  name =?,price =?,image = ?,description = ?, quantity = ?
WHERE id  = ?",product.Name,product.Price,product.Image,product.Description,product.Quantity,product.ID)
if err!= nil {
	return  err
}
return nil
}

func ScanRowsIntoProduct(rows *sql.Rows) (*types.Product,error) {
product := new(types.Product)

err:= rows.Scan(
	&product.ID,
	&product.Name,
	&product.Description,
	&product.Image,
	&product.Quantity,
	&product.CreatedAt,
	

)
}

