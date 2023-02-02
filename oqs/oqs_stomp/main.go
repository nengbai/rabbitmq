package main

import (
	"fmt"
	"net"
	"oqs_stomp/goconfig"
	"os"
	"strconv"
	"time"

	"github.com/go-stomp/stomp"
)

// 读取配置文件
func getConfigFile(filePath string) (configFile *goconfig.ConfigFile) {
	configFile, err := goconfig.LoadConfigFile(filePath)
	if err != nil {
		fmt.Println("load config file error: " + err.Error())
		os.Exit(1)
	}
	return configFile
}

// 使用IP和端口连接到OqsMQ服务器
// 返回OqsMQ连接对象 https://pkg.go.dev/github.com/go-stomp/stomp@v2.1.4+incompatible#example-Connect
func connOqsMq(host, port, sendtimeout, recvtimeout, errortimeout string) (stompConn *stomp.Conn) { // @todo 实现断开重连
	HeartBeatSendTimeOut, _ := strconv.Atoi(sendtimeout)
	HeartBeatRecvTimeOut, _ := strconv.Atoi(recvtimeout)
	HeartBeatErrorTimeOut, _ := strconv.Atoi(errortimeout)
	//fmt.Println("HeartBeatSendTimeOut,HeartBeatRecvTimeOut,HeartBeatErrorTimeOut:", HeartBeatSendTimeOut, HeartBeatRecvTimeOut, HeartBeatErrorTimeOut)
	var options = []func(*stomp.Conn) error{
		//设置读写超时，超时时间为1个小时
		stomp.ConnOpt.Login("haiyouyou/Default/Stomp-User", "jDF6G+]I#sVd+z(sY1ih"),
		stomp.ConnOpt.HeartBeat(time.Duration(HeartBeatSendTimeOut)*time.Second, time.Duration(HeartBeatRecvTimeOut)*time.Second),
		stomp.ConnOpt.HeartBeatError(time.Duration(HeartBeatErrorTimeOut) * time.Second),
		stomp.ConnOpt.AcceptVersion(stomp.V11),
		stomp.ConnOpt.AcceptVersion(stomp.V12),
		stomp.ConnOpt.Header("nonce", "B256B26D320A"),
	}
	x := net.JoinHostPort(host, port)
	// stomp -H cell-1.queue.messaging.ap-tokyo-1.oci.oraclecloud.com -P 61613 -U "haiyouyou/Default/Stomp-User" -W "jDF6G+]I#sVd+z(sY1ih" -V --ssl
	stompConn, err := stomp.Dial("tcp", x, options...)
	if err != nil {
		fmt.Println("connect to oqs_mq server service, error: " + err.Error())
		os.Exit(1)
	}

	return stompConn
}

// 将消息发送到ActiveMQ中
func oqsMqProducer(c chan string, queue string, conn *stomp.Conn) {
	fmt.Println("stompclient:", conn.Server())
	for {
		err := conn.Send(queue, "text/plain", []byte(<-c))
		fmt.Println("send oqs mq..." + queue)
		if err != nil {
			fmt.Println("oqs mq message send erorr: " + err.Error())
		}
	}
}

func main() {
	configPath := "config/config.ini"
	configFile := getConfigFile(configPath)

	host, err := configFile.GetValue("oqs_mq", "host")
	if err != nil {
		fmt.Println("get active_mq host error: " + err.Error())
		os.Exit(1)
	}
	port, err := configFile.GetValue("oqs_mq", "port")
	if err != nil {
		fmt.Println("get active_mq port error: " + err.Error())
		os.Exit(1)
	}

	queue, err := configFile.GetValue("oqs_mq", "queue")
	if err != nil {
		fmt.Println("get active_mq queue error: " + err.Error())
		os.Exit(1)
	}
	sendtimeout, err := configFile.GetValue("oqs_mq", "HeartBeatSendTimeOut")
	if err != nil {
		fmt.Println("get active_mq queue error: " + err.Error())
		os.Exit(1)
	}

	recvtimeout, err := configFile.GetValue("oqs_mq", "HeartBeatRecvTimeOut")
	if err != nil {
		fmt.Println("get oqs_mq queue error: " + err.Error())
		os.Exit(1)
	}
	errortimeout, err := configFile.GetValue("oqs_mq", "HeartBeatErrorTimeOut")
	if err != nil {
		fmt.Println("get oqs_mq queue error: " + err.Error())
		os.Exit(1)
	}

	oqsMq := connOqsMq(host, port, sendtimeout, recvtimeout, errortimeout)
	defer oqsMq.Disconnect()
	c := make(chan string)
	// 启动Go routine发送消息
	go oqsMqProducer(c, queue, oqsMq)

	for i := 0; i < 10000; i++ {
		// 发送1万个消息
		c <- "hello world" + strconv.Itoa(i)
	}

}
