package lib

import (
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQ连接函数
func RabbitMQConn() (conn *amqp.Connection, err error) {
	// RabbitMQ分配的用户名称
	var user string = "admin"
	// RabbitMQ用户的密码
	var pwd string = "Pass@word1"
	// RabbitMQ Broker 的ip地址
	var host string = "131.186.23.190"
	// RabbitMQ Broker 监听的端口
	var port string = "5672"
	var vhost string = "/"
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + vhost
	// 新建一个连接
	conn, err = amqp.Dial(url)
	// 返回连接和错误
	return
}

// 错误处理函数
func ErrorHanding(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
