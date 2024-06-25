package controllers/websocket

import (
    "github.com/labstack/echo/v4"
    "github.com/gorilla/websocket"
    "net/http"
)

var (
    upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            // Allow all connections by default
            return true
        },
    }
)

func WebSocketHandler(c echo.Context) error {
    // Upgrade HTTP connection to WebSocket
    ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    // Echo loop
    for {
        // Read message from client
        _, msg, err := ws.ReadMessage()
        if err != nil {
            break
        }

        // Print received message
        c.Logger().Infof("Received message: %s", msg)

        // Write message back to client
        err = ws.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            break
        }
    }

    return nil
}
