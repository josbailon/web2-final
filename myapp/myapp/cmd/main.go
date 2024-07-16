package main/main

import (
    "database/sql"
    "log"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    _ "github.com/lib/pq"
    "project/services" // Importa tu paquete de servicios de vehículos
)

func main() {
    // Configuración de la base de datos PostgreSQL
    db, err := sql.Open("postgres", "postgres://user:password@localhost/database_name?sslmode=disable")
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }
    defer db.Close()

    // Inicialización del servicio de vehículos
    vehicleService := services.NewVehicleService(db)

    // Configuración de Echo
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Rutas para gestionar vehículos
    e.GET("/vehicles", func(c echo.Context) error {
        vehicles, err := vehicleService.GetAllVehicles()
        if err != nil {
            return echo.NewHTTPError(500, "Error al obtener vehículos")
        }
        return c.JSON(200, vehicles)
    })

    e.GET("/vehicles/:id", func(c echo.Context) error {
        id := c.Param("id")
        vehicleID, err := strconv.Atoi(id)
        if err != nil {
            return echo.NewHTTPError(400, "ID de vehículo no válido")
        }

        vehicle, err := vehicleService.GetVehicleByID(vehicleID)
        if err != nil {
            return echo.NewHTTPError(404, "Vehículo no encontrado")
        }
        return c.JSON(200, vehicle)
    })

    e.POST("/vehicles", func(c echo.Context) error {
        var vehicle services.Vehicle
        if err := c.Bind(&vehicle); err != nil {
            return echo.NewHTTPError(400, "Datos de vehículo inválidos")
        }

        if err := vehicleService.CreateVehicle(&vehicle); err != nil {
            return echo.NewHTTPError(500, "Error al crear vehículo")
        }
        return c.JSON(201, vehicle)
    })

    e.PUT("/vehicles/:id", func(c echo.Context) error {
        id := c.Param("id")
        vehicleID, err := strconv.Atoi(id)
        if err != nil {
            return echo.NewHTTPError(400, "ID de vehículo no válido")
        }

        var vehicle services.Vehicle
        if err := c.Bind(&vehicle); err != nil {
            return echo.NewHTTPError(400, "Datos de vehículo inválidos")
        }
        vehicle.ID = vehicleID

        if err := vehicleService.UpdateVehicle(&vehicle); err != nil {
            return echo.NewHTTPError(500, "Error al actualizar vehículo")
        }
        return c.JSON(200, vehicle)
    })

    e.DELETE("/vehicles/:id", func(c echo.Context) error {
        id := c.Param("id")
        vehicleID, err := strconv.Atoi(id)
        if err != nil {
            return echo.NewHTTPError(400, "ID de vehículo no válido")
        }

        if err := vehicleService.DeleteVehicle(vehicleID); err != nil {
            return echo.NewHTTPError(500, "Error al eliminar vehículo")
        }
        return c.NoContent(204)
    })

    // Iniciar el servidor Echo
    e.Logger.Fatal(e.Start(":8080"))
}
