package subssyncuser

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// SubsSyncTopic represents a subscriber for syncing users.
type SubsSyncUser struct {
	// Add any fields here such as configurations, logger, etc.
}

// InitializeSubsSyncTopic initializes the subscriber for syncing users.
func InitializeSubsSyncUser(client *azservicebus.Client, topicName string, subscriptionName string, shutdownChannel chan struct{}) {
	subs := &SubsSyncUser{
		// Initialize any required fields here.
	}
	subs.start(client, topicName, subscriptionName, shutdownChannel)
}

// start begins the subscription process.
func (subs *SubsSyncUser) start(client *azservicebus.Client, topicName string, subscriptionName string, shutdownChannel chan struct{}) {
	// Create a receiver for the specific topic and subscription.
	receiver, err := client.NewReceiverForSubscription(topicName, subscriptionName, nil)
	if err != nil {
		log.Fatalf("Failed to create receiver: %v", err)
	}

	// Start listening for messages.
	go func() {
		defer receiver.Close(context.Background())
		for {
			select {
			case <-shutdownChannel:
				return
			default:
				// Receive and process messages.
				messages, err := receiver.ReceiveMessages(context.Background(), 10, nil)
				if err != nil {
					log.Printf("Error receiving messages: %v", err)
					continue
				}

				for _, msg := range messages {
					// Process the message.
					processMessage(msg)

					// Complete the message to remove it from the subscription.
					err = receiver.CompleteMessage(context.Background(), msg, nil)
					if err != nil {
						log.Printf("Error completing message: %v", err)
					}
				}
			}
		}
	}()
}

// processMessage handles the processing of each message.
// Boleh dipindahin ke BLL jika diinginkan
func processMessage(msg *azservicebus.ReceivedMessage) {
	fmt.Println("------------------This is from the cloud----------------------------")
	fmt.Println(msg)
	fmt.Println("------------------This is the readable version----------------------------")
	fmt.Println("Message ID:", msg.MessageID)
	fmt.Println("Body:", string(msg.Body))
	fmt.Println("To:", msg.Message().To)
	fmt.Println("Subject:", *msg.Subject)
}
