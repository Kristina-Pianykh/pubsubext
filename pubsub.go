package pubsubext

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub/v2"
	"go.k6.io/k6/js/modules"
	"google.golang.org/api/option"
)

// Register the module with k6
func init() {
	modules.Register("k6/x/pubsub", new(PubSub))
}

// PubSub module type
type PubSub struct{}

type Publisher struct {
	pub *pubsub.Publisher
}

// New creates a new Pub/Sub client and returns a JS-friendly object
func (ps *PubSub) Publisher(projectID string, topic string) *Publisher {
	ctx := context.Background()
	cl, err := pubsub.NewClient(ctx, projectID, option.WithGRPCConnectionPool(4))
	if err != nil {
		log.Printf("failed to create client: %v\n", err)
		fmt.Printf("failed to create client: %v\n", err)
		return nil
	}

	publisher := cl.Publisher(topic)

	return &Publisher{pub: publisher}
}

// Publish publishes a JSON string to the topic
// Returns error string if any, otherwise nil
func (ps *PubSub) Publish(publisher *Publisher, msg string) error {
	ctx := context.Background()
	result := publisher.pub.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("publish failed: %w", err)
		return fmt.Errorf("publish failed: %w", err)
	}

	log.Printf("Published message ID: %s", id)
	fmt.Printf("Published message ID: %s\n", id)
	return nil
}
