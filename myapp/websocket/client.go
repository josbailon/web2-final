// websocket/client.go

package websocket

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// Define WebSocket endpoint URL
var wsEndpoint = "ws://localhost:8080/ws"

// Define WebSocket client struct
type Client struct {
	conn *websocket.Conn
}

// NewClient initializes a new WebSocket client
func NewClient() *Client {
	return &Client{}
}

// Connect establishes a WebSocket connection
func (c *Client) Connect() error {
	// Create a new WebSocket dialer
	dialer := websocket.Dialer{
		HandshakeTimeout:  5 * time.Second,
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
	}

	// Attempt to establish WebSocket connection
	conn, _, err := dialer.Dial(wsEndpoint, nil)
	if err != nil {
		return err
	}

	// Assign the WebSocket connection to the client
	c.conn = conn
	log.Println("Connected to WebSocket server")

	// Handle WebSocket close gracefully
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			select {
			case <-interrupt:
				log.Println("Interrupt signal received, closing WebSocket connection...")
				err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("Error sending close message:", err)
					return
				}
				c.conn.Close()
				return
			}
		}
	}()

	return nil
}

// SendMessage sends a message over the WebSocket connection
func (c *Client) SendMessage(message []byte) error {
	err := c.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}
	return nil
}

// ReadMessage reads messages from the WebSocket connection
func (c *Client) ReadMessage() ([]byte, error) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}

// Close closes the WebSocket connection
func (c *Client) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
		log.Println("WebSocket connection closed")
	}
	return nil
}
