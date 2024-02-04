package routes

import (
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
	r.HandleFunc("/api/dungeons/{id}", controller.SearchDungeon)
}

func postsRoutes(r *mux.Router) {
	r.Use(middlewares.AuthMiddleware)
}

func Setup(r *mux.Router) {
	r.Use(middlewares.LogMiddleware)
	authRoutes(r)
}
