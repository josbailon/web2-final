package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"myapp/models"
	"myapp/repository"
)

var validate = validator.New() // Instancia global de validator

// ObtenerTodosVehiculos devuelve todos los vehículos
func ObtenerTodosVehiculos(c echo.Context) error {
	vehiculos, err := repository.ObtenerTodosVehiculos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, vehiculos)
}

// ObtenerVehiculo devuelve un vehículo por su ID
func ObtenerVehiculo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	vehiculo, err := repository.ObtenerVehiculoPorID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, vehiculo)
}

// CrearVehiculo añade un nuevo vehículo
func CrearVehiculo(c echo.Context) error {
	vehiculo := new(models.Vehiculo)
	if err := c.Bind(vehiculo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	// Validar el vehículo usando validator
	if err := validate.Struct(vehiculo); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
	}
	err := repository.CrearVehiculo(*vehiculo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, vehiculo)
}

// ActualizarVehiculo modifica un vehículo existente
func ActualizarVehiculo(c echo.Context) error {
	vehiculo := new(models.Vehiculo)
	if err := c.Bind(vehiculo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	vehiculo.ID = id

	// Validar el vehículo antes de actualizarlo
	if err := validate.Struct(vehiculo); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
	}

	err := repository.ActualizarVehiculo(*vehiculo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, vehiculo)
}

// EliminarVehiculo elimina un vehículo por su ID
func EliminarVehiculo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.EliminarVehiculo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Vehículo eliminado")
}
