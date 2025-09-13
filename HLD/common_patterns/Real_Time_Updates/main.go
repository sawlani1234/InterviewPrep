package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"real_time_updates/technologies"
	"time"

	"github.com/go-redis/redis/v8"
)

func startSSE() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Register the handler for the "/events" endpoint.
	http.HandleFunc("/events", technologies.SseHandler)

	// Start the server on port 8080.
	log.Println("SSE Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}

func startWS() {
	http.HandleFunc("/ws", technologies.WsHandler)
	log.Println("WebSocket server started on ws://localhost:8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}

func redisPubSub() {
		// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Redis address
	})

	// Check if we're running as publisher or subscriber
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run chat.go [subscribe|publish]")
		return
	}

	// Define the channel name for our chat
	channelName := "chat_channel"

	switch os.Args[1] {
	case "subscribe":
		technologies.Subscriber(rdb, channelName)
	case "publish":
		// Give the subscriber a moment to start up
		time.Sleep(1 * time.Second)
		technologies.Publisher(rdb, channelName)
	default:
		fmt.Println("Invalid command. Use 'subscribe' or 'publish'.")
	}
}

func main() {
	 startSSE()
	// startWS()
	//redisPubSub()
}

