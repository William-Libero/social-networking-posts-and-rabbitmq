package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueue struct {
	Body    string `json:"body"`
	Pattern string `json:"pattern"`
	Age     string `json:"age"`
	Data    string `json:"data"`
}

func NewMessageQueue(bodyM string, pattern string, age string, data string) *MessageQueue {
	return &MessageQueue{
		bodyM, pattern, age, data,
	}
}

func (m *MessageQueue) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(m)

	if err != nil {
		return nil, err
	}
	return bytes, err
}
func HandlesAmqpMessages(queueName string, bodyMessage string) {
	// ConnectToRabbitMQServer
	conn, errorInConnection := amqp.Dial("amqps://ykytzzhp:7q8Ad42GmGxh_bkckCA75h1o48Or77Xz@jackal.rmq.cloudamqp.com/ykytzzhp")
	failOnError(errorInConnection, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// OpensChannel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// DeclaresQueue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// CreatesAmqpContext
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// PublishesAmqpMessage
	fmt.Println(ctx)
	message := NewMessageQueue(bodyMessage, "hello", "20", bodyMessage)
	// TODO: check the error
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
