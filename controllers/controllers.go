package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/William-Libero/social-networking-posts-and-rabbitmq/domain"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/rabbitMQController"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	rabbitMQController.HandlesLikePostMessage("posts", "like_post", id)
	fmt.Fprint(w, "LikePost")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	json.NewDecoder(r.Body).Decode(&post)
	fmt.Println(post)
	rabbitMQController.HandlesCreatePostMessage("posts", "create_post", post)
}
