FROM golang:1.18.2-alpine3.16 AS base
RUN apk update
WORKDIR /src/social-networking-posts-and-rabbitmq
COPY go.mod go.sum ./
COPY . .
RUN go build -o posts_and_rabbitmq .

FROM alpine:3.16 as binary
COPY --from=base /src/posts_and_rabbitmq/posts_and_rabbitmq .
EXPOSE 3000
CMD ["./posts_and_rabbitmq"]