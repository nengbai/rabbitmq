package main

import (
	"fmt"
	"golang/producer"
	"golang/queue"
	"strconv"

	"github.com/gofrs/uuid"
)

type Message struct {
	Typename string `json:"typename,omitempty"`
	UUID     string `json:"uuid,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func main() {
	const MessageEndpoint = "https://cell-1.queue.messaging.ap-tokyo-1.oci.oraclecloud.com"
	qid := "ocid1.queue.oc1.ap-tokyo-1.amaaaaaaj37ijuqa7dumi4aueyrmnhfu5s2mawtbecwtxthutgmoapjhbaba"
	id, _ := uuid.NewV4()
	ids := id.String()
	mac := queue.GetLocalMac()
	ids = mac + ids
	var in string
	var data *Message
	for i := 0; i < 50; i++ {
		in = strconv.Itoa(i)
		fmt.Println(in)
		data = &Message{
			Typename: "order1",
			UUID:     ids,
			Data:     []byte(in),
		}

		content := queue.Encode(data)
		resp, err := producer.PutMessages(ids, qid, MessageEndpoint, content)
		fmt.Println("------")
		//resp, err := consumer.GetMessages(ids, qid, MessageEndpoint)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	}
}
