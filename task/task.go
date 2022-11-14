package main

import (
	"log"
	"os"
	"rabbitmq/lib"
	"strings"

	"github.com/streadway/amqp"
)

func bodyForm(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "no Task"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func main() {
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	lib.ErrorHanding(err, "failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"task:queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	lib.ErrorHanding(err, "Failed to declare a queue")
	// 定义一个消费者
	body := bodyForm(os.Args)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:         []byte(body),
		},
	)
	lib.ErrorHanding(err, "Fail to publish a message !")
	log.Println("send message: s%", body)
}
