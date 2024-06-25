// controllers/auth_controller.go

package controllers/auth_controller

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "project/internal/middleware"
)

func ProtectedRoute(c echo.Context) error {
    username := c.Get("username").(string)
    return c.String(http.StatusOK, "Usuario autenticado: "+username)
}

// main.go

package main

import (
    "github.com/labstack/echo/v4"
    "project/internal/controllers"
    "project/internal/middleware"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.JWTAuthMiddleware)

    // Rutas protegidas con JWT
    api := e.Group("/api")
    api.Use(middleware.JWTAuthMiddleware)

    // Ejemplo de ruta protegida
    api.GET("/protected", controllers.ProtectedRoute)

    e.Logger.Fatal(e.Start(":8080"))
}
