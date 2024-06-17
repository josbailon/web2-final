// server.go

package main

import (
	"log"
	"net/http"

	"your-app-package/websocket" // Replace with your actual package path
	"github.com/gorilla/websocket"
)

func main() {
	// Initialize WebSocket hub
	hub := websocket.NewHub()
	go hub.Run()

	// WebSocket upgrade handler
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all connections
		},
	}

	// WebSocket endpoint handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}

		// Create a new WebSocket client
		client := &websocket.Client{
			Hub:  hub,
			Conn: conn,
			Send: make(chan []byte, 256),
		}
		client.Hub.RegisterClient(client)
		defer func() {
			client.Hub.UnregisterClient(client)
			client.Conn.Close()
		}()

		// Handle incoming messages from the client (you should handle this in a separate goroutine)
		go client.ReadMessages()

		// Handle outgoing messages to the client
		client.WriteMessages()
	})

	// Start HTTP server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
