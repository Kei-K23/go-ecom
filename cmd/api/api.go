package api

import (
	"database/sql"
	"fmt"
	"net/http"

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

	userHandler := users.Handler{}
	// register user routes
	userHandler.RegisterRoutes(subRouter)

	fmt.Printf("Listing on %s\n", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
