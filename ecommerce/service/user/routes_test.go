package user

import (
	"fmt"
	"testing"

	"github.com/joshua468/ecommerce/cmd/api/types"
)



func TestUserServiceHandler(t *testing.T) {
	userStore := &MockUserStore{}
	handler:= NewHandler(userStore)
}

t.Run("Should fail if the user payload is invalid",func(t *testing.T) {
	payload := types.RegisterUserPayload{
		FirstName: "user",
		LastName: "123",
		Email:"invalid",
		Password:"asd",
	}
	marshalled ,err := json.Marshal(payload)
	req,err:= http.NewRequest(http.MethodPost,"/register",bytes.NewBuffer(marshalled))
	if err!= nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecoder()
	router := mux.NewRouter()

	router.HandleFunc("/register",handler.handleRegister)
	router.serveHTTP(rr,req)


	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %d got %d",http.StatusCreated,rr.Code)
	}
})
t.Run("should correctly register the user",func(t *testing.T)) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName: "123",
			Email:"valid@mail.com",
			Password:"asd",
		}
		marshalled ,err := json.Marshal(payload)
		req,err:= http.NewRequest(http.MethodPost,"/register",bytes.NewBuffer(marshalled))
		if err!= nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecoder()
		router := mux.NewRouter()
	
		router.HandleFunc("/register",handler.handleRegister)
		router.serveHTTP(rr,req)
	
	
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d got %d",http.StatusBadRequest,rr.Code)
		}
}
type MockUserStore struct {}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User,error) {
	return nil,fmt.Errorf(user not found)
}

func(m *MockUserStore) GetUserByID(id int) (*types.User,error) {
	return nil,nil
}

func (m *MockUserStore) CreateUser(user *types.User) error {
    return nil
}