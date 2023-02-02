package queue

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/queue"
)

func GetQueue(ids string, qid string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueAdminClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).
	req := queue.GetQueueRequest{OpcRequestId: common.String(ids),
		QueueId: common.String(qid)}

	// Send the request using the service client
	resp, err := client.GetQueue(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func UpdateMessages(ids string, qid string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := queue.UpdateMessagesRequest{OpcRequestId: common.String("6RH5SFHSF7SH0WDWMOBU" + ids),
		QueueId: common.String(qid),
		UpdateMessagesDetails: queue.UpdateMessagesDetails{Entries: []queue.UpdateMessagesDetailsEntry{
			{
				VisibilityInSeconds: common.Int(42690),
				Receipt:             common.String("EXAMPLE-receipt-Value")}}}}

	// Send the request using the service client
	resp, err := client.UpdateMessages(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func GetMessages(ids string, qid string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := queue.GetMessagesRequest{Limit: common.Int(17),
		OpcRequestId:        common.String("MessagesRequest" + ids),
		QueueId:             common.String(qid),
		TimeoutInSeconds:    common.Int(28),
		VisibilityInSeconds: common.Int(10502)}

	// Send the request using the service client
	resp, err := client.GetMessages(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func DeleteMessage(ids string, qid string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := queue.DeleteMessageRequest{QueueId: common.String(qid),
		MessageReceipt: common.String("EXAMPLE-messageReceipt-Value"),
		OpcRequestId:   common.String("DeleteMessage" + ids)}

	// Send the request using the service client
	resp, err := client.DeleteMessage(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func DeleteMessages(ids string, qid string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := queue.DeleteMessagesRequest{DeleteMessagesDetails: queue.DeleteMessagesDetails{Entries: []queue.DeleteMessagesDetailsEntry{{Receipt: common.String("EXAMPLE-receipt-Value")}}},
		OpcRequestId: common.String("queue-demo" + ids),
		QueueId:      common.String(qid)}
	// Send the request using the service client
	resp, err := client.DeleteMessages(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func ChangeQueueCompartment(comid, ids, qid string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueAdminClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := queue.ChangeQueueCompartmentRequest{ChangeQueueCompartmentDetails: queue.ChangeQueueCompartmentDetails{CompartmentId: common.String("ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value")},
		IfMatch:      common.String("EXAMPLE-ifMatch-Value"),
		OpcRequestId: common.String("ChangeQueueCompartmentRequest" + ids),
		QueueId:      common.String(qid)}

	// Send the request using the service client
	resp, err := client.ChangeQueueCompartment(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	fmt.Println(resp)
}

func Encode(data interface{}) string {
	b, _ := json.Marshal(&data)
	// Base64 Standard Encoding
	// sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	sEnc := base64.StdEncoding.EncodeToString(b)
	return sEnc // aGVsbG8gd29ybGQxMjM0NSE/JComKCknLUB+
}

func Decode(data string) []byte {
	// Base64 Standard Decoding
	sDec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return nil
	}

	//fmt.Println("Decode message is:",string(sDec))
	return sDec
}

// GetLocalMac() 获取本机的MAC地址
func GetLocalMac() (mac string) {

	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		fmt.Println(inter.Name)
		mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println("MAC ===== ", mac)
	}
	fmt.Println("MAC = ", mac)
	return mac
}

// GetIps() 获取本机ip地址
func GetIps() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interfaces ipAddress: %v\n", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isVailIpNet := address.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	fmt.Println("ips = ", ips)
	return ips
}
