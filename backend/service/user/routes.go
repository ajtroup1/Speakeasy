package user

import (
	"fmt"
	"net/http"
	"time"

	// "github.com/ajtroup1/speakeasy/config"
	"github.com/ajtroup1/speakeasy/service/auth"
	"github.com/ajtroup1/speakeasy/types"
	"github.com/ajtroup1/speakeasy/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
	friendStore types.FriendStore
}

func NewHandler(store types.UserStore, friendStore types.FriendStore) *Handler {
	return &Handler{store: store, friendStore: friendStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/edit", h.handleEdit).Methods("PUT")
	router.HandleFunc("/friend", h.handleFriend).Methods("POST")
	router.HandleFunc("/unfriend", h.handleUnfriend).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// secret := []byte(config.Envs.JWTSecret)
	// token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": ""})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// fmt.Printf("Received payload: %+v\n", payload)

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	u := types.User{
		Username:    payload.Username,
		Password:    hashedPassword,
		Firstname:   payload.Firstname,
		Lastname:    payload.Lastname,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		ImgLink:     payload.ImgLink,
		Status:      1,
		CreatedAt:   time.Now(),
	}

	// Create user in the database
	err = h.store.CreateUser(u)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, u)
}

func (h *Handler) handleEdit(w http.ResponseWriter, r *http.Request) {
	var payload types.User
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Check if user exists
	existingUser, err := h.store.GetUserByID(payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.ID))
		return
	}

	// Check if any data has changed
	if existingUser.Username == payload.Username &&
		existingUser.Password == payload.Password &&
		existingUser.Firstname == payload.Firstname &&
		existingUser.Lastname == payload.Lastname &&
		existingUser.Email == payload.Email &&
		existingUser.PhoneNumber == payload.PhoneNumber &&
		existingUser.ImgLink == payload.ImgLink &&
		existingUser.Status == payload.Status {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("received information is identical to information in database"))
		return
	}

	// Update user details
	existingUser.Username = payload.Username
	existingUser.Password = payload.Password
	existingUser.Firstname = payload.Firstname
	existingUser.Lastname = payload.Lastname
	existingUser.Email = payload.Email
	existingUser.PhoneNumber = payload.PhoneNumber
	existingUser.ImgLink = payload.ImgLink
	existingUser.Status = payload.Status

	err = h.store.EditUser(*existingUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, existingUser)
}

func (h *Handler) handleFriend(w http.ResponseWriter, r *http.Request) {
	var payload types.FriendPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetUserByID(int(payload.SendID))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.SendID))
		return
	}
	_, err = h.store.GetUserByID(int(payload.ReceiveID))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.ReceiveID))
		return
	}

	err = h.friendStore.FriendUser(payload.SendID, payload.ReceiveID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUnfriend(w http.ResponseWriter, r *http.Request) {
	var payload types.FriendPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetUserByID(int(payload.SendID))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.SendID))
		return
	}
	_, err = h.store.GetUserByID(int(payload.ReceiveID))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.ReceiveID))
		return
	}

	err = h.friendStore.UnfriendUser(payload.SendID, payload.ReceiveID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}