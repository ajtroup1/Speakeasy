package user

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

func TestUser(t *testing.T) {
	userStore := &mockUserStore{}
	friendStore := &mockFriendStore{}
	handler := NewHandler(userStore, friendStore)

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

	t.Run("should edit a user", func(t *testing.T) {
		parsedTime, err := ParseTime("2024-07-20 04:27:55")
		if err != nil {
			t.Fatal(err)
		}
		payload := types.User{
			ID: 1,
			Username:    "adamjtroup2",
			Password:    "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname:   "Adam",
			Lastname:    "Troup",
			Email:       "adamjtroup@gmail.com",
			PhoneNumber: "256-746-6217",
			ImgLink:     "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
			Status: 1,
			CreatedAt: parsedTime,
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

	t.Run("should friend a user", func(t *testing.T) {
		payload := types.FriendPayload{
			SendID: 1,
			ReceiveID: 2,
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/friend", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/friend", handler.handleFriend)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusOK, rr.Code)
		}
	})

}

type mockUserStore struct{}
type mockFriendStore struct{}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) GetUserByID(id int) (*types.User, error) {
    if id == 1 {
        return &types.User{
            ID: 1,
            Username:    "adamjtroup",
            Password:    "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
            Firstname:   "Adam",
            Lastname:    "Troup",
            Email:       "adamjtroup@gmail.com",
            PhoneNumber: "256-746-6217",
            ImgLink:     "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
            Status: 1,
            CreatedAt: time.Date(
                2024,                   // year
                time.July,              // month
                20,                     // day
                4,                      // hour
                27,                     // minute
                55,                     // second
                0,                      // nanoseconds
                time.UTC,               // location
            ),
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

func (s *mockFriendStore) FriendUser(sendID, receiveID uint) error {
	return nil
}

func (s *mockFriendStore) UnfriendUser(sendID, receiveID uint) error {
	return nil
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

