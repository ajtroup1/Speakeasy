package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajtroup1/speakeasy/types"
	"github.com/gorilla/mux"
)

func TestUser(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Username:    "adamjtroup",
			Password:    "Sample123",
			Firstname:   "adam",
			Lastname:    "troup",
			Email:       "adamjtroup@gmail.com",
			PhoneNumber: "256-746-6217",
			ImgLink:     "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("failed with status code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should register a user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Username:    "adamjtroup",
			Password:    "Sample123!",
			Firstname:   "adam",
			Lastname:    "troup",
			Email:       "adamjtroup@gmail.com",
			PhoneNumber: "256-746-6217",
			ImgLink:     "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("failed with status code %d, received %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}

