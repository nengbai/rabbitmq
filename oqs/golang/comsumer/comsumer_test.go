package comsumer

import (
	"encoding/json"
	"fmt"
	"golang/queue"
	"testing"
	// "log"

	"github.com/gofrs/uuid"
)


type Message struct {
	Typename string `json:"Typename,omitempty"`
	UUID     string `json:"uuid,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func TestGetMessages(t *testing.T) {
	const MessageEndpoint = "https://cell-1.queue.messaging.ap-tokyo-1.oci.oraclecloud.com"
	qid := "ocid1.queue.oc1.ap-tokyo-1.amaaaaaaj37ijuqa7dumi4aueyrmnhfu5s2mawtbecwtxthutgmoapjhbaba"
	id, _ := uuid.NewV4()
	ids := id.String()
	mac := queue.GetLocalMac()
	ids = mac + ids
	//queue.GetQueue(ids, qid)
	fmt.Println("------")
	resp, err := GetMessages(ids, qid, MessageEndpoint)
	if err != nil {
		fmt.Println(err)
	}
	var message Message

	//data := make(map[string]interface{})
	//var forever chan struct{}
	//go func() {
		
	for _, v := range resp.Messages {
		messageReceipt := queue.Decode(*v.Content)
		err := json.Unmarshal([]byte(messageReceipt), &message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("message is: ", message.Typename,string(message.Data) )
	}
	//}()
	//	log.Println(" [*] Waiting for logs. To exit press CTRL+C")
	//	<-forever

}
   