package main

import (
	"fmt"
	"rabbitmq/lib"
)

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) MsgContent() string {

	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
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
	msg := "这是测试任务"
	fmt.Println(msg)

	queueExchange := &lib.QueueExchange{
		"test.rabbit",
		"rabbit.key",
		"test.rabbit.mq",
		"direct",
	}
	mq := lib.New(queueExchange)
	mq.RegisterProducer(conn)
	mq.RegisterReceiver(t)
	mq.Start()
}
