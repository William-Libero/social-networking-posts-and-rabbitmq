package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func HandlesPostsAmqpMessages(queueName string) (ch *amqp.Channel, q amqp.Queue) {
	// ConnectToRabbitMQServer
	conn, errorInConnection := amqp.Dial("amqps://ykytzzhp:7q8Ad42GmGxh_bkckCA75h1o48Or77Xz@jackal.rmq.cloudamqp.com/ykytzzhp")
	failOnError(errorInConnection, "Failed to connect to RabbitMQ")

	// OpensChannel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	// DeclaresQueue
	q, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return ch, q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
