package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"myapp/models"
	"myapp/repository"
)

// ObtenerTodasSalidas devuelve todas las salidas de vehículos
func ObtenerTodasSalidas(c echo.Context) error {
	salidas, err := repository.ObtenerTodasSalidas()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, salidas)
}

// ObtenerSalida devuelve una salida de vehículo por su ID
func ObtenerSalida(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	salida, err := repository.ObtenerSalidaPorID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, salida)
}

// CrearSalida añade una nueva salida de vehículo
func CrearSalida(c echo.Context) error {
	salida := new(models.SalidaVehiculo)
	if err := c.Bind(salida); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	// Validar la salida usando validator
	if err := validate.Struct(salida); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
	}
	err := repository.CrearSalida(*salida)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, salida)
}

// ActualizarSalida modifica una salida de vehículo existente
func ActualizarSalida(c echo.Context) error {
	salida := new(models.SalidaVehiculo)
	if err := c.Bind(salida); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	salida.ID = id

	// Validar la salida antes de actualizarla
	if err := validate.Struct(salida); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
	}

	err := repository.ActualizarSalida(*salida)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, salida)
}

// EliminarSalida elimina una salida de vehículo por su ID
func EliminarSalida(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.EliminarSalida(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Salida eliminada")
}
