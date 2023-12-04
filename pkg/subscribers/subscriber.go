package subscriber

import (
	"context"
	"log"
	"sync"
	"time"

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

type ReceivedMessage struct {
	ApplicationProperties      map[string]interface{}
	Body                       []byte
	ContentType                *string
	CorrelationID              *string
	DeadLetterErrorDescription *string
	DeadLetterReason           *string
	DeadLetterSource           *string
	DeliveryCount              uint32
	EnqueuedSequenceNumber     *int64
	EnqueuedTime               *time.Time
	ExpiresAt                  *time.Time
	LockedUntil                *time.Time
	LockToken                  [16]byte
	MessageID                  string
	PartitionKey               *string
	ReplyTo                    *string
	ReplyToSessionID           *string
	ScheduledEnqueueTime       *time.Time
	SequenceNumber             *int64
	SessionID                  *string
	State                      MessageState
	Subject                    *string
	TimeToLive                 *time.Duration
	To                         *string
}

// MessageState represents the current state of a message (Active, Scheduled, Deferred).
type MessageState int32

const (
	MessageStateActive    MessageState = 0
	MessageStateDeferred  MessageState = 1
	MessageStateScheduled MessageState = 2
)
