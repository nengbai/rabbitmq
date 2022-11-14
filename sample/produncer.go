package main

import (
	"encoding/json"
	"log"
	"rabbitmq/lib"

	"github.com/streadway/amqp"
)

type simpleDemo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func main() {
	// 连接RabbitMQ服务器
	conn, err := lib.RabbitMQConn()
	lib.ErrorHanding(err, "Failed to connect to RabbitMQ")
	// 关闭连接
	defer conn.Close()
	// 新建一个通道
	ch, err := conn.Channel()
	lib.ErrorHanding(err, "Failed to open a channel")
	// 关闭通道
	defer ch.Close()
	// 声明或者创建一个队列用来保存消息
	q, err := ch.QueueDeclare(
		// 队列名称
		"simple:queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	lib.ErrorHanding(err, "Failed to declare a queue")
	data := simpleDemo{
		Name: "Tom",
		Addr: "Beijing",
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		lib.ErrorHanding(err, "struct to json failed")
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataBytes,
		})
	log.Printf(" [x] Sent %s", dataBytes)
	lib.ErrorHanding(err, "Failed to publish a message")
}
