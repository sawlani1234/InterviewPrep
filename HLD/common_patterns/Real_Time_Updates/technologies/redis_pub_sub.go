package technologies

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

// Use a background context
var ctx = context.Background()

// publisher sends messages from the command line to a Redis channel.
func Publisher(rdb *redis.Client, channelName string) {
	// Create a reader to read from the command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter messages to publish (type 'exit' to quit):")

	for {
		fmt.Print("> ")
		// Read a line of text from the user
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Println("Publisher exiting.")
			return
		}

		// Publish the message to the Redis channel
		err := rdb.Publish(ctx, channelName, message).Err()
		if err != nil {
			fmt.Printf("Error publishing message: %v\n", err)
		}
	}
}

// subscriber listens to a Redis channel and prints any received messages.
func Subscriber(rdb *redis.Client, channelName string) {
	// Subscribe to the given channel
	pubsub := rdb.Subscribe(ctx, channelName)

	// Wait for confirmation that subscription is created
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Subscribed to channel: %s\n", channelName)
	fmt.Println("Listening for messages...")

	// Get the channel to listen for messages on
	ch := pubsub.Channel()

	// Loop indefinitely and print messages as they arrive
	for msg := range ch {
		fmt.Printf("\n< Received: %s\n> ", msg.Payload)
	}
}
