package routes

import (
	"fmt"
	"net/http"

	"github.com/William-Libero/social-networking-posts-and-rabbitmq/controllers"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/middlewares"
	"github.com/gorilla/mux"
)

func HandleRoutes() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleWare)
	r.HandleFunc("/", controllers.Home).Methods("Get")
	http.ListenAndServe(":8000", r)
	fmt.Println("Server running on port 8000")
}
