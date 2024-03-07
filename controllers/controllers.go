package controllers

import (
	"fmt"
	"net/http"

	"github.com/William-Libero/social-networking-posts-and-rabbitmq/rabbitmq"
)

func Home(w http.ResponseWriter, r *http.Request) {
	connectToRabbitMQServer()
	fmt.Fprint(w, "Hello World!")
}

func connectToRabbitMQServer() {
	rabbitmq.HandlesAmqpMessages("hello", "hello world!")
}
