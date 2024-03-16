# Go API - Social Network Project

<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://go.dev/images/go-logo-white.svg" width="200" alt="Go Logo" /></a>
  <a href="https://www.rabbitmq.com/" target="blank"><img src="https://www.rabbitmq.com/img/rabbitmq-logo-with-name.svg" width="200" alt="Go Logo" /></a>
</p>

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/)
[![RabbitMQ](https://img.shields.io/badge/RabbitMQ-v1.9.0-orange.svg)](https://www.rabbitmq.com/)

This project implements a REST API in Go that serves as a backend for a social networking application developed in Next.js. The API enables functionalities such as searching, creating, and liking posts.

## Main Features

- **Fetch Posts**: Users can see posts.
- **Create Posts**: Users can create new posts.
- **Like Posts**: Users can like posts.
- **Real-time Communication**: Messages are published in a RabbitMQ queue to create and like posts, enabling real-time updates.
- **Microservices Architecture**: Another API, developed in Nest.js (<a href="https://github.com/William-Libero/social-networking-posts-service" target="blank">Project Link</a>), consumes messages from the RabbitMQ queue to handle post creation and liking.
