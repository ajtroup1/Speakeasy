package channel

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

func TestChannel(t *testing.T) {
	channelStore := &mockChannelStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(channelStore, userStore)

	t.Run("should fail if channel payload is invalid", func(t *testing.T) {
		payload := types.CreateChannelPayload{
			Name:        "",
			Description: "",
			CreatedBy:   1,
			ImgLink:     "",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/createchannel", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/createchannel", handler.handleCreate)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("failed with status code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should create a channel", func(t *testing.T) {
		payload := types.CreateChannelPayload{
			Name:        "samplename",
			Description: "sample description",
			CreatedBy:   1,
			ImgLink:     "",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/createchannel", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/createchannel", handler.handleCreate)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockChannelStore struct{}

func (s *mockChannelStore) CreateChannel(channel types.Channel) error {
	return nil
}

type mockUserStore struct{}

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
