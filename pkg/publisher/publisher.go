package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/google/uuid"
)

var (
	publisherOnce       sync.Once
	serviceBusPublisher *AzureServiceBusPublisher
	shutdownChannel     = make(chan struct{})
)

type Publisher interface {
	PublishMessage(model interface{}, topic string) error
}

type AzureServiceBusPublisher struct {
	client    *azservicebus.Client
	sender    *azservicebus.Sender
	topicName string
}

func InitializeServiceBusPublisher(connStr string, topicName string) *AzureServiceBusPublisher {
	publisherOnce.Do(func() {
		client, err := azservicebus.NewClientFromConnectionString(connStr, nil)
		if err != nil {
			fmt.Printf("Failed to create service bus client: %s\n", err)
			os.Exit(1)
		}

		sender, err := client.NewSender(topicName, nil)
		if err != nil {
			fmt.Printf("Failed to create service bus sender: %s\n", err)
			os.Exit(1)
		}

		serviceBusPublisher = &AzureServiceBusPublisher{
			client:    client,
			sender:    sender,
			topicName: topicName,
		}

		fmt.Println("Azure Service Bus Publisher initialized")
	})

	return serviceBusPublisher
}

func ShutdownServiceBusPublisher() {
	close(shutdownChannel)

	if serviceBusPublisher != nil {
		if err := serviceBusPublisher.ShutdownServiceBus(); err != nil {
			fmt.Printf("Error shutting down Service Bus Publisher: %v\n", err)
		}
	}
}

func (p *AzureServiceBusPublisher) ShutdownServiceBus() error {
	if p.sender != nil {
		if err := p.sender.Close(context.Background()); err != nil {
			return err
		}
	}
	if p.client != nil {
		if err := p.client.Close(context.Background()); err != nil {
			return err
		}
	}
	return nil
}

func (p *AzureServiceBusPublisher) PublishMessage(model interface{}, topic string) error {
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}

	// Generate a unique MessageId
	messageId := uuid.New().String()
	partitionKey := "partitionKey"

	message := &azservicebus.Message{
		Body:         data,
		MessageID:    &messageId,
		PartitionKey: &partitionKey,
	}
	// Check the shutdown channel before sending the message
	select {
	case <-shutdownChannel:
		return fmt.Errorf("Service Bus Publisher is shutting down")
	default:
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sendErr := p.sender.SendMessage(ctx, message, nil)
	if sendErr != nil {
		return sendErr
	}

	return nil
}
