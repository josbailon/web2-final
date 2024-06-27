// routes/routes.go

package routes/routes

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "project/repositories"
    "project/services"
)

// InitRoutes inicializa todas las rutas de la aplicación
func InitRoutes(e *echo.Echo, db *sql.DB) {
    // Inicializar repositorios
    paymentRepo := repositories.NewPaymentRepository(db)
    vehicleRepo := repositories.NewVehicleRepository(db)

    // Inicializar servicios
    paymentService := services.NewPaymentService(paymentRepo)
    vehicleService := services.NewVehicleService(vehicleRepo)

    // Rutas para pagos
    e.GET("/payments", paymentService.GetAllPaymentsHandler)
    e.GET("/payments/:id", paymentService.GetPaymentByIDHandler)
    e.POST("/payments", paymentService.CreatePaymentHandler)
    e.PUT("/payments/:id", paymentService.UpdatePaymentHandler)
    e.DELETE("/payments/:id", paymentService.DeletePaymentHandler)

    // Rutas para vehículos
    e.GET("/vehicles", vehicleService.GetAllVehiclesHandler)
    e.GET("/vehicles/:id", vehicleService.GetVehicleByIDHandler)
    e.POST("/vehicles", vehicleService.CreateVehicleHandler)
    e.PUT("/vehicles/:id", vehicleService.UpdateVehicleHandler)
    e.DELETE("/vehicles/:id", vehicleService.DeleteVehicleHandler)
}
