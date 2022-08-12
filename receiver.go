package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	// Use a callback to receive messages via subscription1.
	sub := client.Subscription("my-sub")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(string(m.Data))
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Println(err)
	}
}
