package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/handlers"
)

// InitRoutes inicializa todas las rutas para la API
func InitRoutes(e *echo.Echo) {
	// Rutas para vehículos
	e.GET("/vehiculos", handlers.ObtenerTodosVehiculos)
	e.GET("/vehiculos/:id", handlers.ObtenerVehiculo)
	e.POST("/vehiculos", handlers.CrearVehiculo)
	e.PUT("/vehiculos/:id", handlers.ActualizarVehiculo)
	e.DELETE("/vehiculos/:id", handlers.EliminarVehiculo)

	// Rutas para pagos
	e.GET("/pagos", handlers.ObtenerTodosPagos)
	e.GET("/pagos/:id", handlers.ObtenerPago)
	e.POST("/pagos", handlers.CrearPago)
	e.PUT("/pagos/:id", handlers.ActualizarPago)
	e.DELETE("/pagos/:id", handlers.EliminarPago)

	// Rutas para salidas de vehículos
	e.GET("/salidas", handlers.ObtenerTodasSalidas)
	e.GET("/salidas/:id", handlers.ObtenerSalida)
	e.POST("/salidas", handlers.CrearSalida)
	e.PUT("/salidas/:id", handlers.ActualizarSalida)
	e.DELETE("/salidas/:id", handlers.EliminarSalida)
}
