package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jn0x/reddigo/http/controllers"
	"github.com/jn0x/reddigo/http/middlewares"
)

func authRoutes(r *mux.Router) {
	controller := controllers.NewAuthController()
	r.HandleFunc("/api/login", controller.Login).Methods("POST")
	r.HandleFunc("/api/signup", controller.Signup).Methods("POST")
}

func dungeonRoutes(r *mux.Router) {
	controller := controllers.NewDungeonController()
	r.HandleFunc("/api/dungeons/get/{id}", controller.GetDungeon).Methods("GET")
	r.HandleFunc("/api/dungeons/search", controller.SearchDungeon).Methods("GET")
	r.HandleFunc("/api/dungeons/create", controller.CreateDungeon).Methods("POST")
	r.HandleFunc("/api/dungeons/delete/{id}", controller.DeleteDungeon).Methods("DELETE")

	r.Handle("/api/dungeons/join", middlewares.AuthMiddleware(http.HandlerFunc(controller.JoinDungeon))).
		Methods("POST")
}

func postsRoutes(r *mux.Router) {
	r.Use(middlewares.AuthMiddleware)
}

func Setup(r *mux.Router) {
	r.Use(middlewares.LogMiddleware)
	authRoutes(r)
	dungeonRoutes(r)
}
