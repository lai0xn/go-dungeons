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

var count int64 = 0

func (s *Server) Serve() {
	r := mux.NewRouter()
	routes.Setup(r)

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		count++

		return nil
	})
	fmt.Println("Server Started")
	fmt.Println("number of routes : ", count)
	http.ListenAndServe(s.ADDR, r)
}
