package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/William-Libero/social-networking-posts-and-rabbitmq/db"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/domain"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/rabbitMQController"
	"github.com/gorilla/mux"
)

type Post struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Likes       int    `json:"likes"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	supabase := db.CreateConection()

	var results []interface{}
	err := supabase.DB.From("posts").Select("*").Execute(&results)
	if err != nil {
		panic(err)
	}

	output := make([]string, len(results))
	for i := 0; i < len(results); i += 1 {
		jsonBytes, _ := json.Marshal(results[i])
		fmt.Println(string(jsonBytes))
		output = append(output, string(jsonBytes))
	}

	json.NewEncoder(w).Encode(output)
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
