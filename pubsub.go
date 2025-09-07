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

type Config struct {
	ProjectID string
	Topic     string
}

// New creates a new Pub/Sub client and returns a JS-friendly object
func (ps *PubSub) Publisher(config *Config) *Publisher {
	ctx := context.Background()
	cl, err := pubsub.NewClient(ctx, config.ProjectID, option.WithGRPCConnectionPool(4))
	if err != nil {
		log.Printf("failed to create client: %v\n", err)
		return nil
	}

	publisher := cl.Publisher(config.Topic)

	return &Publisher{pub: publisher}
}

// PublishBatch publishes multiple messages
func (ps *PubSub) PublishBatch(p *Publisher, msgs []string) []string {
	ctx := context.Background()
	var ids []string

	for _, msg := range msgs {
		result := p.pub.Publish(ctx, &pubsub.Message{Data: []byte(msg)})
		id, err := result.Get(ctx)
		if err != nil {
			log.Printf("publish failed: %v", err)
			return ids
		}
		ids = append(ids, id)
		fmt.Printf("Published message ID: %s\n", id)
	}

	return ids
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
		return fmt.Errorf("publish failed: %w", err)
	}

	log.Printf("Published message ID: %s", id)
	return nil
}
