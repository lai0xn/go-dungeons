package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jn0x/reddigo/http/routes"
)

type Server struct {
	ADDR string
}

func NewServer(addr string) *Server {
	return &Server{ADDR: addr}
}

func (s *Server) Serve() {
	r := mux.NewRouter()
	routes.Setup(r)
	fmt.Println("Server Started")
	http.ListenAndServe(s.ADDR, r)
}
