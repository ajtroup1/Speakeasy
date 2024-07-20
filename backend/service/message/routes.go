package message

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ajtroup1/speakeasy/types"
	"github.com/ajtroup1/speakeasy/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.MessageStore
	userStore types.UserStore
}

func NewHandler(store types.MessageStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/createmessage", h.handleCreate).Methods("POST")
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateMessagePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	//ensure user exists
	_, err := h.userStore.GetUserByID(payload.CreatedBy)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesnt exists", payload.CreatedBy))
		return
	}
	//ensure channel exists

	m := types.Message{
		Content:    payload.Content,
		CreatedBy:    payload.CreatedBy,
		CreatedAt:   time.Now(),
		ChannelD: payload.ChannelD,
	}

	err = h.store.CreateMessage(m)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, m)
}