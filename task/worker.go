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
	q, err := ch.QueueDeclare(
		"task:queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	lib.ErrorHanding(err, "Failed to declare a queue")
	err = ch.Qos(
		1,
		0,
		false,
	)
	lib.ErrorHanding(err, "Fail to set Qos!")
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	lib.ErrorHanding(err, "Fail to Register a comsumer!")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Reviced a message:%s\n", string(d.Body))
			log.Println("Done")
			d.Ack(false)
		}

	}()
	log.Println(" [*] Wait for message.To text Press CTRL + C")
	<-forever
}
