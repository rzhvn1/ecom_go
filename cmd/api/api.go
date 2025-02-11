package api

import (
	"database/sql"
	"ecom_go/services/shop"
	"ecom_go/services/shopcategory"
	"ecom_go/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRouter := subrouter.PathPrefix("/users").Subrouter()
	shopRouter := subrouter.PathPrefix("/shops").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(userRouter)

	shopCategoryStore := shopcategory.NewStore(s.db)
	shopStore := shop.NewStore(s.db)
	shopHandler := shop.NewHandler(shopStore, shopCategoryStore, userStore)
	shopHandler.RegisterRoutes(shopRouter)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
