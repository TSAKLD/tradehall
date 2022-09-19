package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"steamsale/service"
)

type Server struct {
	port   string
	router *mux.Router
}

// NewServer returns http server to work with.
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: mux.NewRouter(),
	}
}

// setRoutes activating handlers and sets routes for http server.
func (s *Server) setRoutes(userServ service.UserService, itemServ service.ItemService) {
	hand := newHandler(userServ, itemServ)

	//Routes for common Users
	s.router.HandleFunc("/users", hand.RegisterUser).Methods(http.MethodPost)
	s.router.HandleFunc("/users", hand.EditUser).Methods(http.MethodPatch)
	s.router.HandleFunc("/users", hand.RemoveUser).Methods(http.MethodDelete)
	s.router.HandleFunc("/users", hand.FindUser).Methods(http.MethodGet)
	s.router.HandleFunc("/users", hand.EditUser).Methods(http.MethodPatch)

	//Routes for Admins
	s.router.HandleFunc("/items", hand.AddItem).Methods(http.MethodPost)
}

// Start is just starting http server.
func (s *Server) Start(userServ service.UserService, itemServ service.ItemService) error {
	s.setRoutes(userServ, itemServ)

	fmt.Println("Server is listening...")

	return http.ListenAndServe(":"+s.port, s.router)
}
