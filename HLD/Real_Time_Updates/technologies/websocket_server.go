package technologies

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections for demo purposes
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket client connected.")

	for {
		// Read message from client (optional, for echo demo)
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", msg)

		// Echo message back to client
		response := fmt.Sprintf("Echo: %s", msg)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(response)); err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	log.Println("WebSocket client disconnected.")
}
