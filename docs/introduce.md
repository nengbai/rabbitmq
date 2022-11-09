# RabbitMQ

## 1、 简介

RabbitMQ是一个开源的遵循AMQP(Advanced Message Queuing Protocol)协议实现的基于erlang开发的消息中间，由Connection，Exchange，队列和消息4部分组成。生产者通过Exchange路由规则把消息写入对应队列，消息可持续存储在队列，消息主动推给消费者，已消费的消息从队列中删除。

* Connection: 为生产者和消费者构建安全通道。
* Exchange: 设置路由规则，对应到相应的队列。
* 队列：简单队列和工作队列。
  ![RabbitMQ架构](./img/1.1-0.png "RabbitMQ架构")

## 2、   应用场景

  微服务和无服务器应用程的解藕，消息可靠传输。
