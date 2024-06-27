// websocket_controller.go

package controllers/websocket_controller

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Permitir todas las conexiones WebSocket
    },
}

// WebSocketHandler maneja las solicitudes WebSocket
func WebSocketHandler(c echo.Context) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    // Ejemplo de bucle de escucha y respuesta para WebSocket
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return err
        }
        if err := conn.WriteMessage(messageType, p); err != nil {
            return err
        }
    }
}
