package subscriber

import (
	"context"
	"log"
	"sync"

	"gopubsub/pkg/subscribers/subssyncuser"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

var (
	client          *azservicebus.Client
	connectionOnce  sync.Once
	shutdownChannel = make(chan struct{})
)

// InitializeServiceBus initializes the Service Bus client and subscribers.
func InitializeServiceBus(connStr string) {
	connectionOnce.Do(func() {
		var err error
		client, err = azservicebus.NewClientFromConnectionString(connStr, nil)
		if err != nil {
			log.Fatalf("Error creating Service Bus client: %v", err)
		}
	})

	// Initialize your topic subscribers here.
	subssyncuser.InitializeSubsSyncUser(client, "ep-identity-dev", "ep-sub-user-syncuser", shutdownChannel)
	//subssyncprofile.InitializeSubsSyncProfile(client, "ep-identity-dev", "ep-sub-profile-syncprofile", shutdownChannel)
}

// ShutdownServiceBus gracefully closes the Service Bus client.
func ShutdownServiceBusSubscriber() {
	close(shutdownChannel)

	if client != nil {
		client.Close(context.Background())
	}
}
