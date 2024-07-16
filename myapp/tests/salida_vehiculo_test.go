package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/handlers"
	"myapp/models"
)

func TestObtenerTodasSalidas(t *testing.T) {
	initDB()
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/salidas", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.ObtenerTodasSalidas(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body.String())
	}
}

func TestCrearSalida(t *testing.T) {
	initDB()
	e := setupRouter()

	salida := models.SalidaVehiculo{
		VehiculoID:       1,
		FechaSalida:      time.Now(),
		TiempoEstacionado: "2 hours",
	}
	salidaJSON, _ := json.Marshal(salida)

	req := httptest.NewRequest(http.MethodPost, "/salidas", bytes.NewBuffer(salidaJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.CrearSalida(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "2 hours")
	}
}

func TestObtenerSalida(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea una salida para luego obtenerla
	salida := models.SalidaVehiculo{
		VehiculoID:       1,
		FechaSalida:      time.Now(),
		TiempoEstacionado: "2 hours",
	}
	salidaJSON, _ := json.Marshal(salida)

	req := httptest.NewRequest(http.MethodPost, "/salidas", bytes.NewBuffer(salidaJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearSalida(c)

	// Obtén la salida recién creada
	id := rec.Body.String() // Asume que la respuesta contiene el ID de la salida
	req = httptest.NewRequest(http.MethodGet, "/salidas/"+id, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.ObtenerSalida(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "2 hours")
	}
}

func TestActualizarSalida(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea una salida para luego actualizarla
	salida := models.SalidaVehiculo{
		VehiculoID:       1,
		FechaSalida:      time.Now(),
		TiempoEstacionado: "2 hours",
	}
	salidaJSON, _ := json.Marshal(salida)

	req := httptest.NewRequest(http.MethodPost, "/salidas", bytes.NewBuffer(salidaJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearSalida(c)

	// Actualiza la salida recién creada
	id := rec.Body.String() // Asume que la respuesta contiene el ID de la salida
	salida.TiempoEstacionado = "3 hours"
	salidaJSON, _ = json.Marshal(salida)
	req = httptest.NewRequest(http.MethodPut, "/salidas/"+id, bytes.NewBuffer(salidaJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.ActualizarSalida(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "3 hours")
	}
}

func TestEliminarSalida(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea una salida para luego eliminarla
	salida := models.SalidaVehiculo{
		VehiculoID:       1,
		FechaSalida:      time.Now(),
		TiempoEstacionado: "2 hours",
	}
	salidaJSON, _ := json.Marshal(salida)

	req := httptest.NewRequest(http.MethodPost, "/salidas", bytes.NewBuffer(salidaJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearSalida(c)

	// Elimina la salida recién creada
	id := rec.Body.String() // Asume que la respuesta contiene el ID de la salida
	req = httptest.NewRequest(http.MethodDelete, "/salidas/"+id, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.EliminarSalida(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Salida eliminada")
	}
}