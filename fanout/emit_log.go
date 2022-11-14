package main

import (
	"log"
	"os"
	"rabbitmq/lib"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	lib.ErrorHanding(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"exchange_logs", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	lib.ErrorHanding(err, "Failed to declare an exchange")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	body := bodyFrom(os.Args)

	err = ch.Publish(
		"exchange_logs", // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	lib.ErrorHanding(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s\n", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
