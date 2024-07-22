package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/ajtroup1/speakeasy/types"
	"github.com/gorilla/mux"
)

func logAllUsers(handler *Handler) {
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users", handler.handleGetUsers)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		log.Printf("Failed to get users, status code: %d", rr.Code)
		log.Printf("Response body: %s", rr.Body.String())
		return
	}

	var users []types.User
	if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
		log.Printf("Failed to unmarshal users: %v", err)
		return
	}

	log.Println("Users in the database:")
	for _, user := range users {
		log.Printf("User: %+v\n", user)
	}
}

func TestUser(t *testing.T) {
	userStore := &mockUserStore{}
	friendStore := &mockFriendStore{}
	blockStore := &mockBlockStore{}
	handler := NewHandler(userStore, friendStore, blockStore)

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
			Username:           "adamjtroup",
			Password:           "Sample123!",
			Firstname:          "adam",
			Lastname:           "troup",
			Email:              "adamjtroup@gmail.com",
			PhoneNumber:        "256-746-6217",
			ImgLink:            "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
			TextNotifications:  false,
			EmailNotifications: false,
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

	// t.Run("should send an email upon registration", func(t *testing.T) {
	// 	payload := types.RegisterUserPayload{
	// 		Username:           "adamjtroup",
	// 		Password:           "Sample123!",
	// 		Firstname:          "adam",
	// 		Lastname:           "troup",
	// 		Email:              "adamjtroup@gmail.com",
	// 		PhoneNumber:        "256-746-6217",
	// 		ImgLink:            "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
	// 		TextNotifications:  false,
	// 		EmailNotifications: true,
	// 	}
	// 	marshal, _ := json.Marshal(payload)

	// 	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	rr := httptest.NewRecorder()
	// 	router := mux.NewRouter()

	// 	router.HandleFunc("/register", handler.handleRegister)

	// 	router.ServeHTTP(rr, req)

	// 	if rr.Code != http.StatusCreated {
	// 		t.Errorf("failed with status code %d, received %d", http.StatusCreated, rr.Code)
	// 	}
	// })

	// t.Run("should send a text upon registration", func(t *testing.T) {
	// 	payload := types.RegisterUserPayload{
	// 		Username:           "adamjtroup",
	// 		Password:           "Sample123!",
	// 		Firstname:          "adam",
	// 		Lastname:           "troup",
	// 		Email:              "adamjtroup@gmail.com",
	// 		PhoneNumber:        "256-746-6217",
	// 		ImgLink:            "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
	// 		TextNotifications:  true,
	// 		EmailNotifications: false,
	// 	}
	// 	marshal, _ := json.Marshal(payload)

	// 	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	rr := httptest.NewRecorder()
	// 	router := mux.NewRouter()

	// 	router.HandleFunc("/register", handler.handleRegister)

	// 	router.ServeHTTP(rr, req)

	// 	if rr.Code != http.StatusCreated {
	// 		t.Errorf("failed with status code %d, received %d", http.StatusCreated, rr.Code)
	// 	}
	// })

	t.Run("should edit a user", func(t *testing.T) {
		payload := types.EditUserPayload{
			ID:          1,
			Username:    "adamjtroup2",
			Firstname:   "Adam",
			Lastname:    "Troup",
			Email:       "adamjtroup@gmail.com",
			PhoneNumber: "256-746-6217",
			ImgLink:     "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/edit", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/edit", handler.handleEdit)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should change a user's password", func(t *testing.T) {
		payload := types.ChangePasswordPayload{
			UserID:             1,
			CurrentPassword:    "Sample123!",
			NewPassword:        "Sample1234!",
			ConfirmNewPassword: "Sample1234!",
		}

		marshal, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPut, "/changepassword", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/changepassword", handler.handleChangePassword)

		router.ServeHTTP(rr, req)

		expectedStatus := http.StatusOK

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", expectedStatus, rr.Code)
		}
	})

	t.Run("should friend a user", func(t *testing.T) {
		payload := types.FriendPayload{
			SendID:    1,
			ReceiveID: 2,
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/friend", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/friend", handler.handleFriend)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
			t.Logf("response body: %s", rr.Body.String())
		}
	})

	t.Run("should accept a friend request from user", func(t *testing.T) {
		payload := types.FriendPayload{
			SendID:    1,
			ReceiveID: 2,
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/acceptfriend", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/acceptfriend", handler.handleFriend)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
			t.Logf("response body: %s", rr.Body.String())
		}
	})

	t.Run("should unfriend a user", func(t *testing.T) {
		payload := types.FriendPayload{
			SendID:    1,
			ReceiveID: 2,
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/unfriend", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/unfriend", handler.handleFriend)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
			t.Logf("response body: %s", rr.Body.String())
		}
	})

	t.Run("should get friendships by user id", func(t *testing.T) {
		userID := uint(1)
		req, err := http.NewRequest(http.MethodGet, "/friendships/"+strconv.Itoa(int(userID)), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/friendships/{id}", handler.handleGetFriendshipsByID)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
			t.Logf("response body: %s", rr.Body.String())
		}
	})

	t.Run("should block a user", func(t *testing.T) {
		payload := types.FriendPayload{
			SendID:    1,
			ReceiveID: 3,
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/block", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/block", handler.handleFriend)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
			t.Logf("response body: %s", rr.Body.String())
		}
	})

	t.Run("should unblock a user", func(t *testing.T) {
		payload := types.FriendPayload{
			SendID:    1,
			ReceiveID: 3,
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/unblock", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/unblock", handler.handleFriend)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
			t.Logf("response body: %s", rr.Body.String())
		}
	})
}

type mockUserStore struct{}
type mockFriendStore struct{}
type mockBlockStore struct{}

func (s *mockUserStore) GetAllUsers() ([]*types.User, error) {
	return nil, fmt.Errorf("users not found")
}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
	if id == 1 {
		return &types.User{
			ID:          1,
			Username:    "adamjtroup",
			Password:    "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname:   "Adam",
			Lastname:    "Troup",
			Email:       "adamjtroup@gmail.com",
			PhoneNumber: "256-746-6217",
			ImgLink:     "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
			Status:      1,
			CreatedAt: time.Date(
				2024,      // year
				time.July, // month
				20,        // day
				4,         // hour
				27,        // minute
				55,        // second
				0,         // nanoseconds
				time.UTC,  // location
			),
			TextNotifications:  true,
			EmailNotifications: true,
		}, nil
	} else if id == 2 {
		return &types.User{
			ID:          2,
			Username:    "friend",
			Password:    "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname:   "Sample",
			Lastname:    "Name",
			Email:       "sample@mail.com",
			PhoneNumber: "111-222-3333",
			ImgLink:     "",
			Status:      1,
			CreatedAt: time.Date(
				2024,      // year
				time.July, // month
				20,        // day
				4,         // hour
				27,        // minute
				55,        // second
				0,         // nanoseconds
				time.UTC,  // location
			),
			TextNotifications:  false,
			EmailNotifications: false,
		}, nil
	} else if id == 3 {
		return &types.User{
			ID:          3,
			Username:    "enemy",
			Password:    "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname:   "Sample",
			Lastname:    "Name",
			Email:       "sample2@mail.com",
			PhoneNumber: "111-222-3333",
			ImgLink:     "",
			Status:      1,
			CreatedAt: time.Date(
				2024,      // year
				time.July, // month
				20,        // day
				4,         // hour
				27,        // minute
				55,        // second
				0,         // nanoseconds
				time.UTC,  // location
			),
			TextNotifications:  false,
			EmailNotifications: false,
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

func (s *mockUserStore) ChangePassword(id uint, currentPassword, newPassword, confirmNewPassword string) error {
	return nil
}

func (s *mockFriendStore) FriendUser(sendID, receiveID uint) error {
	return nil
}

func (s *mockFriendStore) Accept(sendID, receiveID uint) error {
	return nil
}

func (s *mockFriendStore) UnfriendUser(sendID, receiveID uint) error {
	return nil
}

func (s *mockFriendStore) Refriend(sendID, receiveID uint) error {
	return nil
}

func (s *mockFriendStore) GetFriendshipByIDs(sendID, receiveID uint) (bool, error) {
	return false, nil
}

func (s *mockFriendStore) GetFriendshipsByID(userID uint) ([]*types.User, error) {
	return nil, nil
}

func (s *mockBlockStore) BlockUser(sendID, receiveID uint) error {
	return nil
}

func (s *mockBlockStore) UnblockUser(sendID, receiveID uint) error {
	return nil
}

func (s *mockBlockStore) GetBlockByIDs(sendID, receiveID uint) (bool, error) {
	return false, nil
}

// Correct usage of time.Date
func ParseTime(timestamp string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
