package queue

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

type PubSubClient struct{}

func PubSubClientInit() *PubSubClient {
	return &PubSubClient{}
}

func (r *PubSubClient) Connect() (c *pubsub.Client, err error) {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	return client, nil
}

func (r *PubSubClient) PushToQueue(queueName string, queueData []byte) error {
	ctx := context.Background()
	connection, err := r.Connect()
	if err != nil {
		return err
	}
	defer connection.Close()
	fmt.Printf("was here")
	// topic, errTopic := connection.CreateTopic(context.Background(), queueName)
	// fmt.Printf("Published a message; msg ID: %v\n", errTopic)
	// if errTopic != nil {
	// 	return errTopic
	// }

	topic := connection.Topic(queueName)
	// result := topic.Publish(ctx, &pubsub.Message{
	// 	Data: []byte(msg),
	// })

	res := topic.Publish(ctx, &pubsub.Message{Data: queueData})
	id, err := res.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	defer topic.Stop()
	return nil
}
