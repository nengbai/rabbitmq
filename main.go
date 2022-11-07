package main

import (
	"fmt"
	"rabbitmq-demo/lib"
)

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) MsgContent("131.186.23.190", "5672", "admin" , "Pass@word1" , "/" ) string {
	
	return t.msgContent
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {
	msg := fmt.Sprintf("这是测试任务")

	t := &TestPro{
		msg,
	}
	queueExchange := &lib.QueueExchange{
		"test.rabbit",
		"rabbit.key",
		"test.rabbit.mq",
		"direct",
	}
	mq := lib.New(queueExchange)
	mq.RegisterProducer(t)
	mq.RegisterReceiver(t)
	mq.Start()
}
