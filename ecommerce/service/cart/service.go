package cart

import (
	"fmt"
	"go/types"
)

func getCartItemsIDs(item []types.CartItem) ([]int,error) {
	productsIds := make([]int, len(items))
	for i,item := range items {
		if  item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity  for the  product %d",item.ProductID)
		}
		productsIds[i] = item.ProductID
	}
	return productsIds,nil
}
func (h *Handler) CreateOrder(ps  []types.Product,item []types.CartItem,userID int) (int,float64,error) {
ProductMap := make(map[int]types.Product)
for _, product := range ps {
productMap[product.ID] = product


if err := checkIfCartIsInStock(items,productsMap);err!= nil  {
return  0,0,nil
}
totalPrice := calculateTotalPrice(items,productsMap)

for  _, item := range items {
	product := productMap[item.ProductID]
	product.Quantity -= item.Quantity
	h.productStore.UpdateProduct(product)
}
orderID,err := h.store.CreateOrder(types.Order {
	UserID: userID,
	Total:totalPrice,
	Status:"pending",
	Address: "some address",
})
if err!= nil {
	return 0,0,err
}
return 0,totalPrice,nil
}

for _,item:= range items {
	h.store.CreateOrderItem {
		OrderID: orderID,
		ProductID:  item.ProductID,
		Quality:item.Quantity,
		Price: productMap[item.ProductID].Price,
		
	}
}






func checkIfCartIsInStock(cartItems []types.CartCheckOutItem,products map[int]types.Product) error {
if len(cartItems) ==  0 {
	return fmt.Errorf("cart is empty")
}

for _,item  := range cartItems {
	products,ok := products[items.ProductID]
	if !ok {
		return fmt.Errorf("products %d  is not available  in the store,please refresh your cart",item.ProductID)
	}
	if  product.Quantity <  item.Quantity {
		return fmt.Errorf("product %s  is not  available in the quantity requested",product.Name)
	}
}
return nil
}
func calculateTotalPrice(cartItems []types.CartItem, products map [int]types.Product) float64 {
	var  total float64
	for  _,item := range cartItems {
		products := product[item.productID]
		total +=  product.Price * float64(item.Quantity)
	}
	return total
}

}