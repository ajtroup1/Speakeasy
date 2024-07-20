package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ajtroup1/speakeasy/types"
	"github.com/gorilla/mux"
)

// Test the creation of a message
func TestMessage(t *testing.T) {
	// Initialize mocks
	messageStore := &mockMessageStore{}
	userStore := &mockUserStore{}

	// Initialize the handler with mock stores
	handler := NewHandler(messageStore, userStore)

	// Create a new router and register routes
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	t.Run("should fail if message payload is invalid", func(t *testing.T) {
		payload := types.CreateMessagePayload{
			Content:    "", // Invalid payload
			CreatedBy:  1,
			ChannelD:   1,
		}
		marshal, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/createmessage", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should create a message", func(t *testing.T) {
		payload := types.CreateMessagePayload{
			Content:    "sample message",
			CreatedBy:  1,
			ChannelD:   1,
		}
		marshal, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/createmessage", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}

		// Optionally, check the response body here
		var responseBody types.Message
		if err := json.NewDecoder(rr.Body).Decode(&responseBody); err != nil {
			t.Fatal(err)
		}

		// Check if the response body has the expected content
		if responseBody.Content != payload.Content {
			t.Errorf("expected content %s, got %s", payload.Content, responseBody.Content)
		}
	})
}

// Mock implementation for MessageStore
type mockMessageStore struct{}

func (s *mockMessageStore) CreateMessage(message types.Message) error {
	return nil
}

// Mock implementation for UserStore
type mockUserStore struct{}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	if id == 1 {
		return &types.User{
			ID:          1,
			Username:    "adamjtroup",
			Password:    "hashedpassword",
			Firstname:   "Adam",
			Lastname:    "Troup",
			Email:       "adamjtroup@gmail.com",
			PhoneNumber: "256-746-6217",
			ImgLink:     "https://example.com/image.jpg",
			Status:      1,
			CreatedAt:   time.Now(),
		}, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (s *mockUserStore) EditUser(user types.User) error {
	return nil
}
