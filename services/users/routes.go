package users

import (
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-ecom/services/auth"
	"github.com/Kei-K23/go-ecom/types"
	"github.com/Kei-K23/go-ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")

	router.HandleFunc("/login", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	// get JSON payload
	var payload types.RegisterUserPayload

	// parse payload to json
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationErr := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", validationErr))
	}

	_, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s not found", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("cannot hash password"))
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
