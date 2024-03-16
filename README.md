# Social Network API

[![Go Version](https://img.shields.io/badge/Go-1.17-blue.svg)](https://golang.org/)
[![Next.js Version](https://img.shields.io/badge/Next.js-12.0.7-blue.svg)](https://nextjs.org/)
[![Nest.js Version](https://img.shields.io/badge/Nest.js-8.1.4-blue.svg)](https://nestjs.com/)
[![RabbitMQ Version](https://img.shields.io/badge/RabbitMQ-3.9.11-blue.svg)](https://www.rabbitmq.com/)

This project implements a REST API in Go that serves as a backend for a social networking application developed in Next.js. The API enables functionalities such as searching, creating, and liking posts.

## Main Features

- **Search Posts**: Users can search for posts based on specific criteria.
- **Create Posts**: Users can create new posts.
- **Like Posts**: Users can like posts.
- **Real-time Communication**: Messages are published in a RabbitMQ queue to create and like posts, enabling real-time updates.
- **Microservices Architecture**: Another API, developed in Nest.js, consumes messages from the RabbitMQ queue to handle post creation and liking.
