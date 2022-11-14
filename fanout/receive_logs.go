package main

import (
	"log"
	"rabbitmq/lib"
)

func main() {
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	lib.ErrorHanding(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"exchange_logs", // exchange name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	lib.ErrorHanding(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	lib.ErrorHanding(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		"exchange_logs", // exchange
		false,
		nil,
	)
	lib.ErrorHanding(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	lib.ErrorHanding(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", string(d.Body))
		}
	}()

	log.Println(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
