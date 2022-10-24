package main

import (
	"log"
	"rabbitmq-demo/lib"
)

func main() {
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	lib.ErrorHanding(err, "Failed to open a channel")
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
	// 将预取计数器设置为1
	// 在并行处理中将消息分配给不同的工作进程
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
