package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kei-K23/go-ecom/config"
	"github.com/Kei-K23/go-ecom/middleware"
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
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Create a subrouter for /users route
	usersRouter := router.PathPrefix("/users").Subrouter()

	// Apply middleware to the subrouter
	usersRouter.Use(middleware.CheckAuthMiddleware)

	// Define routes for /users subrouter
	usersRouter.HandleFunc("", h.handleGetUser).Methods(http.MethodGet)

	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {

	// check request has auth token
	claims := r.Context().Value(middleware.ClaimsContextKey).(*auth.JWTClaim)
	u, err := h.store.GetUserByID(claims.UserId)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user credential"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.CreatedUserRes{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	})

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	// get JSON payload
	var payload types.LoginUserPayload

	// parse payload to json
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationErr := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", validationErr))
	}

	u, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user credential"))
		return
	}

	if !auth.ComparePassword(u.Password, payload.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid user credential"))
	}

	token, err := auth.CreateJWT([]byte(config.Env.Secret), u.ID)

	if err != nil {
		log.Fatal(err)
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
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

	if err == nil {
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

	utils.WriteJSON(w, http.StatusCreated, types.CreatedUserRes{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	})
}
