FROM golang:1.18.2-alpine3.16 AS base
RUN apk update
WORKDIR /src/posts-and-rabbitmq
COPY go.mod go.sum ./
COPY . .
RUN go build -o posts-and-rabbitmq ./cmd/api

FROM alpine:3.16 as binary
COPY --from=base /src/posts-and-rabbitmq/posts-and-rabbitmq .
EXPOSE 3000
CMD ["./posts-and-rabbitmq"]