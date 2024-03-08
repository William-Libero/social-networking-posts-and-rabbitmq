package rabbitMQController

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/William-Libero/social-networking-posts-and-rabbitmq/domain"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/messageQueueCreator"
	"github.com/William-Libero/social-networking-posts-and-rabbitmq/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func HandlesLikePostMessage(queueName string, eventName string, bodyMessage string) {
	ch, q := rabbitmq.HandlesPostsAmqpMessages(queueName)
	publishesAmqpMessageLikePost(ch, q, CreatesAmqpContext(), bodyMessage, eventName)
}

func publishesAmqpMessageLikePost(ch *amqp.Channel, q amqp.Queue, ctx context.Context, bodyMessage string, eventName string) {
	fmt.Println(bodyMessage, eventName)
	// PublishesAmqpMessage
	message := messageQueueCreator.NewMessageQueue(bodyMessage, eventName, "20", bodyMessage)
	body, _ := message.Marshal()

	publishErr := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(publishErr, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func HandlesCreatePostMessage(queueName string, eventName string, bodyMessage domain.Post) {
	ch, q := rabbitmq.HandlesPostsAmqpMessages(queueName)
	publishesAmqpMessageCreatePost(ch, q, CreatesAmqpContext(), bodyMessage, eventName)
}

func publishesAmqpMessageCreatePost(ch *amqp.Channel, q amqp.Queue, ctx context.Context, bodyMessage domain.Post, eventName string) {
	fmt.Println(bodyMessage, eventName)
	// PublishesAmqpMessage
	message := messageQueueCreator.NewMessageQueue(bodyMessage, eventName, "20", bodyMessage)
	body, _ := message.Marshal()

	publishErr := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(publishErr, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func CreatesAmqpContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ctx
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
