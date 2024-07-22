package channel

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ajtroup1/speakeasy/types"
	"github.com/ajtroup1/speakeasy/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.ChannelStore
	userStore types.UserStore
}

func NewHandler(store types.ChannelStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/createchannel", h.handleCreate).Methods("POST")
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateChannelPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Println("1")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		log.Println("2")
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Ensure user exists
	_, err := h.userStore.GetUserByID(payload.CreatedBy)
	if err != nil {
		log.Println("3")
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.CreatedBy))
		return
	}

	c := types.Channel{
		Name:        payload.Name,
		Description: payload.Description,
		ImgLink:     payload.ImgLink,
		CreatedBy:   payload.CreatedBy,
		CreatedAt:   time.Now(),
	}

	err = h.store.CreateChannel(c)
	if err != nil {
		log.Println("4")
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, c)
}
