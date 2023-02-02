package comsumer

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/queue"
)

func GetMessages(ids, qid, MessageEndpoint string) (*queue.GetMessagesResponse, error) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := queue.NewQueueClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	client.Host = MessageEndpoint
	// Create a request and dependent object(s).
	req := queue.GetMessagesRequest{Limit: common.Int(17),
		OpcRequestId:        common.String("Comsumer" + ids),
		QueueId:             common.String(qid),
		TimeoutInSeconds:    common.Int(28),
		VisibilityInSeconds: common.Int(10502)}

	// Send the request using the service client
	resp, err := client.GetMessages(context.Background(), req)
	if err != nil {
		return nil, err
	}
	// Retrieve value from the response.
	return &resp, nil
}
