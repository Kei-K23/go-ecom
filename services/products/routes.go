package products

import (
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-ecom/middleware"
	"github.com/Kei-K23/go-ecom/types"
	"github.com/Kei-K23/go-ecom/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	productsRouter := router.PathPrefix("/products").Subrouter()
	productsRouter.Use(middleware.CheckAuthMiddleware)

	productsRouter.HandleFunc("", h.createProduct).Methods(http.MethodPost)
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateProduct

	// parse payload to json
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		validationErr := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", validationErr))
	}

	p, err := h.store.CreateProduct(payload)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, p)
}