package cart

import(
	"github.com/gorilla/mux"
	"github.com/joshua468/ecommerce/types"
)

type Handler struct {
store types.OrderStore
productStore types.ProductStore
userStore types.UserStore

}

type NewHandler(store types.OrderStore,productStore types.ProductStore,userStore types.UserStore)  *Handler {
	return &Handler{store:store,productStore:productStore,userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout",auth.WithJWTAuth,(h.handleCheckout,h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r &cart);err!= nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}
	if err := utils.Validate.Struct(cart);err!= nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("Invalid payload : %v",errors))
		return
	}
	productsIDs,err := getCartItemsIDs(cart.Items)
	if err!= nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}
	ps,err:= h.store.ProductStore.GetProducts(productIDs)
	if err!= nil {
		utils.WriteError(w, http.StatusInternalServerErrror,err)
		return
	}
	orderID,totalPrice,err := h.createOrder(ps,cart.Items,userID)
	if err != nil  {
		utils.WriteJSON(w,http.StatusBadRequest,err)
		return
	}
	utils.WriteJSON(w,http.StatusOK,map[string]any {
		"title_price" : totalPrice,
		"order_id": orderID,


	})

}