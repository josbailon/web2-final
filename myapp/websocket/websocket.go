package websocket

import (
    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
    "net/http"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleWebSocket(c echo.Context) error {
    conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            return err
        }
        // Aquí puedes procesar el mensaje recibido y responder
        err = conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            return err
        }
    }
}
