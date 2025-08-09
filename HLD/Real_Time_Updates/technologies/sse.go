package technologies

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func SseHandler(w http.ResponseWriter, r *http.Request) {
	// Set the essential headers for an SSE connection.
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Use a flusher to send data chunks immediately.
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Log when a client connects.
	log.Println("Client connected.")
	// Send an initial SSE comment line to prime the connection.
	fmt.Fprintf(w, ": connected\n\n")
	flusher.Flush()

	// Use the request's context to detect when the client disconnects.
	ctx := r.Context()

	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected.")
			return
		default:
			message := fmt.Sprintf("The server time is: %v", time.Now().Format("15:04:05"))
			fmt.Println("here", message)

			fmt.Fprintf(w, "data: %s\n\n", message)

			flusher.Flush()

			time.Sleep(2 * time.Second)
		}
	}
}
