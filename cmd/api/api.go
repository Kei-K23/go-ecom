package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-ecom/middleware"
	"github.com/Kei-K23/go-ecom/services/products"
	"github.com/Kei-K23/go-ecom/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.Use(middleware.LoggingMiddleware)

	userStore := users.NewStore(s.db)
	productsStore := products.NewStore(s.db)

	userHandler := users.NewHandler(userStore)
	// register user routes
	userHandler.RegisterRoutes(subRouter)

	productHandler := products.NewHandler(productsStore)
	productHandler.RegisterRoutes(subRouter)

	fmt.Printf("Listing on %s\n", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
