# Social Network API

<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://go.dev/images/go-logo-white.svg" width="200" alt="Go Logo" /></a>
  <a href="https://www.rabbitmq.com/" target="blank"><img src="https://www.rabbitmq.com/img/rabbitmq-logo-with-name.svg" width="200" alt="Go Logo" /></a>
</p>

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/)

This project implements a REST API in Go that serves as a backend for a social networking application developed in Next.js. The API enables functionalities such as searching, creating, and liking posts.

## Main Features

- **Fetch Posts**: Users can see posts.
- **Create Posts**: Users can create new posts.
- **Like Posts**: Users can like posts.
- **Real-time Communication**: Messages are published in a RabbitMQ queue to create and like posts, enabling real-time updates.
- **Microservices Architecture**: Another API, developed in Nest.js (<a href="https://github.com/William-Libero/social-networking-posts-service" target="blank"><img src="https://nestjs.com/img/logo-small.svg" width="40" alt="Nest Logo" /></a>), consumes messages from the RabbitMQ queue to handle post creation and liking.
