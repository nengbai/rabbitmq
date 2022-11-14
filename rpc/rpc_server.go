package main

import (
	"log"
	"rabbitmq/lib"
	"strconv"

	"github.com/streadway/amqp"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	lib.ErrorHanding(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	lib.ErrorHanding(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	lib.ErrorHanding(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	lib.ErrorHanding(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer cancel()
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			lib.ErrorHanding(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			lib.ErrorHanding(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
