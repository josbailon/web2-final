package cmd/main

import (
    "github.com/labstack/echo/v4"
    "project/internal/controllers"
    "project/internal/middleware"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Rutas REST
    // Autenticación y autorización
    e.POST("/login", controllers.Login)
    e.POST("/register", controllers.Register)

    // Rutas protegidas con JWT
    api := e.Group("/api")
    api.Use(middleware.JWTAuth())

    // Ejemplo de ruta protegida
    api.GET("/protected", controllers.ProtectedRoute)

    // Rutas GraphQL
    e.POST("/graphql", controllers.GraphQLHandler())

    // WebSocket
    e.GET("/ws", controllers.WebSocketHandler)

    e.Logger.Fatal(e.Start(":8080"))
}
