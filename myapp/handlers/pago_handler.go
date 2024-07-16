package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"myapp/models"
	"myapp/repository"
)

// ObtenerTodosPagos devuelve todos los pagos
func ObtenerTodosPagos(c echo.Context) error {
	pagos, err := repository.ObtenerTodosPagos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pagos)
}

// ObtenerPago devuelve un pago por su ID
func ObtenerPago(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pago, err := repository.ObtenerPagoPorID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pago)
}

// CrearPago añade un nuevo pago
func CrearPago(c echo.Context) error {
	pago := new(models.Pago)
	if err := c.Bind(pago); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	// Validar el pago usando validator
	if err := validate.Struct(pago); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
	}
	err := repository.CrearPago(*pago)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pago)
}

// ActualizarPago modifica un pago existente
func ActualizarPago(c echo.Context) error {
	pago := new(models.Pago)
	if err := c.Bind(pago); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input: "+err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	pago.ID = id

	// Validar el pago antes de actualizarlo
	if err := validate.Struct(pago); err != nil {
		return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
	}

	err := repository.ActualizarPago(*pago)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pago)
}

// EliminarPago elimina un pago por su ID
func EliminarPago(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.EliminarPago(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Pago eliminado")
}
