package routes

import (
	"net/http"

	"github.com/William-Libero/social-networking-posts-and-rabbitmq/controllers"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/middlewares"
	"github.com/gorilla/mux"
)

func HandleRoutes() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleWare)
	r.HandleFunc("/", controllers.Home).Methods("Get")
	r.HandleFunc("/likePost/{id}", controllers.LikePost).Methods("Post")
	r.HandleFunc("/createPost", controllers.CreatePost).Methods("Post")
	http.ListenAndServe(":8000", r)
}
