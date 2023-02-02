import oci

config = oci.config.from_file()
#validate_config(config)


# Initialize service client with default config file
queue_client = oci.queue.QueueClient(config)
print(queue_client)


# Send the request to service, some parameters are not required, see API
# doc for more info
put_messages_response = queue_client.put_messages(
    queue_id="ocid1.queue.oc1.ap-tokyo-1.amaaaaaaj37ijuqa7dumi4aueyrmnhfu5s2mawtbecwtxthutgmoapjhbaba",
    put_messages_details=oci.queue.models.PutMessagesDetails(
        messages=[
            oci.queue.models.PutMessagesDetailsEntry(
                content="EXAMPLE-content-Value")]),opc_request_id="UMBWJA4IDH3XEYNKLATW123456")

# Get the data from response
print(put_messages_response.data)

