package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func main() {
	err := os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	if err != nil {
		// TODO: Handle error.
	}
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "test")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// Publish "hello world" on topic1.
	topic, err := client.CreateTopic(ctx, "topic1")
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.CreateSubscription(ctx, "my-sub", pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		log.Fatal(err)
	}
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	msgID, err := res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("msg id :", msgID)
}
