package product

import (
	"github.com/gorilla/mux"
	"github.com/joshua468/ecommerce/types"
)

type Handler struct {
store types.ProductStore
}

type NewHandler(store types.ProductStore) *Handler {
	return &Handler{store:store}
}

func (h *Handler) RegisterRoutes(mux *Router) {
	router.HandleFunc("/products",h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products",h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
ps,err := h.store.GetProducts()
if err!= nil {
	utils.WriteError(w,http.StatusInternalServerError,err)
	return
}
utils.WriteJSON(w,http.StatusOk,ps)
}