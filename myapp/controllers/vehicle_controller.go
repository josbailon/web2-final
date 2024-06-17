package controllers

import (
    "net/http"
    "time"
    "parking-management/services"
    "github.com/labstack/echo/v4"
)

type ExitRequest struct {
    VehicleID uint      `xml:"vehicle_id" validate:"required"`
    OutTime   time.Time `xml:"out_time" validate:"required"`
}

func RegisterVehicleExit(c echo.Context) error {
    var request ExitRequest
    if err := c.Bind(&request); err != nil {
        return c.XML(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
    }
    if err := c.Validate(&request); err != nil {
        return c.XML(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    payment, err := services.ProcessVehicleExit(request.VehicleID, request.OutTime)
    if err != nil {
        return c.XML(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.XML(http.StatusOK, payment)
}
