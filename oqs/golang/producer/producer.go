// This is an automatically generated code sample.
// To make this code sample work in your Oracle Cloud tenancy,
// please replace the values for any parameters whose current values do not fit
// your use case (such as resource IDs, strings containing ‘EXAMPLE’ or ‘unique_id’, and
// boolean, number, and enum parameters with values not fitting your use case).

package producer

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/queue"
)

func PutMessages(ids, qid, MessageEndpoint string, content string) (*queue.PutMessagesResponse, error) {

	config := common.DefaultConfigProvider()
	//client, err := queue.NewQueueClientWithConfigurationProvider(config)
	client, err := queue.NewQueueClientWithConfigurationProvider(config)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	client.Host = MessageEndpoint
	// helpers.FatalIfError(err)
	pt := client.Endpoint()
	fmt.Printf("client endpoint is:%s\n", pt)

	// Create a request and dependent object(s).
	req := queue.PutMessagesRequest{
		QueueId: common.String(qid),
		PutMessagesDetails: queue.PutMessagesDetails{
			Messages: []queue.PutMessagesDetailsEntry{
				{
					Content: common.String(content)}},
		},
		OpcRequestId:    common.String("PutMessagesRequest" + ids),
		RequestMetadata: common.RequestMetadata{},
	}

	// Send the request using the service client
	resp, err := client.PutMessages(context.Background(), req)
	if err != nil {
		return nil, err
	}
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	return &resp, nil
}
