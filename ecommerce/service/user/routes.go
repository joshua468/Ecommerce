package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joshua468/ecommerce/cmd/api/service/auth"
	"github.com/joshua468/ecommerce/cmd/api/types"
	"github.com/joshua468/ecommerce/cmd/api/utils"
)

type Handler struct {
	store  types.UserStore
}

func NewHandler(Store types.UserStore) *Handler {
	return &Handler{store:Store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {  
//get JSON payload
var payload types.RegisterUserPayload
if err := utils.ParseJSON(r,&payload);err!= nil {
	utils.WriteError(w,http.StatusBadRequest,err)
	return
}

//validate payload

if err:= utils.Validate.Struct(payload);err!= nil {
	errors := err.(validator.ValidationErrors)
	utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid payload &v",errors))
	return
}

//check if user exist
_,err := h.store.GetUserByEmail(payload.Email) 
if err == nil {
	utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user with email %s already exist",
	payload.Email))
	return 
}
hashedPassword,err:=  auth.HashPassword(payload.Password)
if err!= nil {
	utils.WriteError(w,http.StatusBadRequest, fmt.Errorf("user with email
	%s already",payload.Email))
	return
}
err = h.store.CreateUser(types.User {
	FirstName: payload.FirstName,
	LastName: payload.LastName,
	Email: payload.Email,
	Password : payload.Password,
})
if err != nil {
	utils.WriteError(w, http.StatusInternalServerError,err)
	return
}
utils.WriteJSON(w,http.StatusInternalServerError,err) 
return

}


