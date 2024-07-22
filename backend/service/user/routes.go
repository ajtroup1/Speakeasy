package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/ajtroup1/speakeasy/config"
	"github.com/ajtroup1/speakeasy/service/auth"
	"github.com/ajtroup1/speakeasy/service/email"
	"github.com/ajtroup1/speakeasy/types"
	"github.com/ajtroup1/speakeasy/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store       types.UserStore
	friendStore types.FriendStore
	blockStore  types.BlockStore
}

func NewHandler(store types.UserStore, friendStore types.FriendStore, blockStore types.BlockStore) *Handler {
	return &Handler{store: store, friendStore: friendStore, blockStore: blockStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.handleGetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", h.handleGetUserByID).Methods("GET")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/edit", h.handleEdit).Methods("PUT")
	router.HandleFunc("/changepassword", h.handleChangePassword).Methods("PUT")
	router.HandleFunc("/friend", h.handleFriend).Methods("POST")
	router.HandleFunc("/unfriend", h.handleUnfriend).Methods("POST")
	router.HandleFunc("/block", h.handleBlock).Methods("POST")
	router.HandleFunc("/unblock", h.handleUnblock).Methods("POST")
}

func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	us, err := h.store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, us)
}

func (h *Handler) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: %v", err))
		return
	}
	u, err := h.store.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, u)
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

	// Uncomment and use the appropriate JWT secret from your config
	// secret := []byte(config.Envs.JWTSecret)
	// secret := []byte("your_jwt_secret")
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
		Username:           payload.Username,
		Password:           hashedPassword,
		Firstname:          payload.Firstname,
		Lastname:           payload.Lastname,
		Email:              payload.Email,
		PhoneNumber:        payload.PhoneNumber,
		ImgLink:            payload.ImgLink,
		Status:             1,
		CreatedAt:          time.Now(),
		TextNotifications:  payload.TextNotifications,
		EmailNotifications: payload.EmailNotifications,
	}

	// Create user in the database
	err = h.store.CreateUser(u)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if u.EmailNotifications {
		htmlBody := email.GetRegisterHtmlBody(u.Firstname, u.Username)
		err = email.SendEmail(u.Email, "Welcome to Speakeasy!", htmlBody)
		if err != nil {
			log.Fatal(err)
		}
	}

	utils.WriteJSON(w, http.StatusCreated, u)
}

func (h *Handler) handleEdit(w http.ResponseWriter, r *http.Request) {
	var payload types.EditUserPayload
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
		existingUser.Firstname == payload.Firstname &&
		existingUser.Lastname == payload.Lastname &&
		existingUser.Email == payload.Email &&
		existingUser.PhoneNumber == payload.PhoneNumber &&
		existingUser.ImgLink == payload.ImgLink {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("received information is identical to information in database"))
		return
	}

	// Update user details
	existingUser.Username = payload.Username
	existingUser.Firstname = payload.Firstname
	existingUser.Lastname = payload.Lastname
	existingUser.Email = payload.Email
	existingUser.PhoneNumber = payload.PhoneNumber
	existingUser.ImgLink = payload.ImgLink

	err = h.store.EditUser(*existingUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, existingUser)
}

func (h *Handler) handleChangePassword(w http.ResponseWriter, r *http.Request) {
	var payload types.ChangePasswordPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse JSON: %w", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		log.Printf("Invalid payload: %v", errors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetUserByID(int(payload.UserID))
	if err != nil {
		log.Printf("User with id %d doesn't exist: %v", payload.UserID, err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.UserID))
		return
	}

	err = h.store.ChangePassword(uint(payload.UserID), payload.CurrentPassword, payload.NewPassword, payload.ConfirmNewPassword)
	if err != nil {
		log.Printf("User with id %d doesn't exist: %v", payload.UserID, err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error with change password function"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleFriend(w http.ResponseWriter, r *http.Request) {
	var payload types.FriendPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse JSON: %w", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		log.Printf("Invalid payload: %v", errors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetUserByID(int(payload.SendID))
	if err != nil {
		log.Printf("User with id %d doesn't exist: %v", payload.SendID, err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.SendID))
		return
	}

	_, err = h.store.GetUserByID(int(payload.ReceiveID))
	if err != nil {
		log.Printf("User with id %d doesn't exist: %v", payload.ReceiveID, err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.ReceiveID))
		return
	}

	//does a friendship exist? if so, just change the status
	found, err := h.friendStore.GetFriendshipByIDs(payload.SendID, payload.ReceiveID)
	if err != nil {
		log.Printf("Failed to add friend: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add friend: %w", err))
		return
	}
	if found {
		err := h.friendStore.Refriend(payload.SendID, payload.ReceiveID)
		if err != nil {
			log.Printf("Failed to add friend: %v", err)
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add friend: %w", err))
			return
		}
		utils.WriteJSON(w, http.StatusOK, nil)
		return
	}

	err = h.friendStore.FriendUser(payload.SendID, payload.ReceiveID)
	if err != nil {
		log.Printf("Failed to add friend: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add friend: %w", err))
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

	//are they friends?
	friends, err := h.friendStore.GetFriendshipByIDs(payload.SendID, payload.ReceiveID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !friends {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("friendship between user id %d and id %d doesn't exist", payload.SendID, payload.ReceiveID))
	}

	err = h.friendStore.UnfriendUser(payload.SendID, payload.ReceiveID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleBlock(w http.ResponseWriter, r *http.Request) {
	var payload types.FriendPayload

	// Parse the JSON payload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Block the user
	err := h.blockStore.BlockUser(payload.SendID, payload.ReceiveID)
	if err != nil {
		log.Printf("Failed to block user: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to block user: %w", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUnblock(w http.ResponseWriter, r *http.Request) {
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

	//are they blocked?
	friends, err := h.blockStore.GetBlockByIDs(payload.SendID, payload.ReceiveID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !friends {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("block between user id %d and id %d doesn't exist", payload.SendID, payload.ReceiveID))
	}

	err = h.blockStore.UnblockUser(payload.SendID, payload.ReceiveID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
