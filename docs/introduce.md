# RabbitMQ

## 1、 简介

RabbitMQ是基于Erlang开发的，遵循AMQP(Advanced Message Queuing Protocol)协议的消息中件。其组件包括Connection,Exchange，队列和消息4部分组成。生产者通过Exchange路由规则把消息写入对应队列，消息可持续存储在队列，消息主动推给消费者，已消费的消息从队列中删除。

* Connection: 为生产者和消费者构建安全通道。
* Exchange: 设置路由规则，对应到相应的队列。
* 队列：简单队列和工作队列。
  ![RabbitMQ架构](./img/1.1-0.png "RabbitMQ架构")

## 2、   应用场景

  1. 微服务和无服务器应用程的解藕，消息可靠消费。

     如电商系统：支付，库存，物流都交给MQ做消息可靠消费。
    ![RabbitMQ架构](./img/1.2-0.png "RabbitMQ架构")
  2.异步处理
    ![RabbitMQ架构](./img/1.2-2.png "RabbitMQ架构") 
    服务之间异步调用是异步的。
    例如: A调用B，B需要花费很长时间执行，但是A需要知道B什么时候可以执行完。
    以前一般有两种方式：
    - A过一段时间去调用B的查询api查询。
    - 或者A提供一个callback api， B执行完之后调用API通A服务。
  
    以上两种都不是很优雅，使用消息总线，可以很方便解决这个问题：
    - A调用B服务后，只需要监听B处理完成的消息，当B处理完成后，会发送一条消息给MQ
    - MQ 会将此消息转发给A服务。
    这样A服务既不用循环调用B的查询API，也不用提供callback API。同样B服务也不用做这些操作。A服务还能及时的得到异步处理成功的消息。

   

