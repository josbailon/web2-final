package routes

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "parking-management/controllers"
    "parking-management/config"
    "parking-management/utils"
)

func InitRoutes(e *echo.Echo) {
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Validator = &utils.CustomValidator{Validator: utils.NewValidator()}
    e.POST("/login", controllers.Login)
    e.POST("/register", controllers.Register)

    // Secured routes
    r := e.Group("/vehicles")
    r.Use(middleware.JWT([]byte(config.JwtSecret)))
    r.POST("/exit", controllers.RegisterVehicleExit)
    r.GET("/:id", controllers.GetVehicleByID)
}
