package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/db"
	"myapp/handlers"
	"myapp/models"
	"myapp/routes"
)

func setupRouter() *echo.Echo {
	e := echo.New()
	routes.InitRoutes(e)
	return e
}

func initDB() {
	// Inicializar la base de datos para pruebas
	db.Init()
}

func TestObtenerTodosVehiculos(t *testing.T) {
	initDB()
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/vehiculos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.ObtenerTodosVehiculos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body.String())
	}
}

func TestCrearVehiculo(t *testing.T) {
	initDB()
	e := setupRouter()

	vehiculo := models.Vehiculo{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Ano:    2020,
		Placa:  "ABC123",
	}
	vehiculoJSON, _ := json.Marshal(vehiculo)

	req := httptest.NewRequest(http.MethodPost, "/vehiculos", bytes.NewBuffer(vehiculoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.CrearVehiculo(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Toyota")
	}
}

func TestObtenerVehiculo(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea un vehículo para luego obtenerlo
	vehiculo := models.Vehiculo{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Ano:    2020,
		Placa:  "ABC123",
	}
	vehiculoJSON, _ := json.Marshal(vehiculo)

	req := httptest.NewRequest(http.MethodPost, "/vehiculos", bytes.NewBuffer(vehiculoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearVehiculo(c)

	// Obtén el vehículo recién creado
	var createdVehiculo models.Vehiculo
	json.Unmarshal(rec.Body.Bytes(), &createdVehiculo)
	id := strconv.Itoa(createdVehiculo.ID)

	req = httptest.NewRequest(http.MethodGet, "/vehiculos/"+id, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.ObtenerVehiculo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Toyota")
	}
}

func TestActualizarVehiculo(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea un vehículo para luego actualizarlo
	vehiculo := models.Vehiculo{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Ano:    2020,
		Placa:  "ABC123",
	}
	vehiculoJSON, _ := json.Marshal(vehiculo)

	req := httptest.NewRequest(http.MethodPost, "/vehiculos", bytes.NewBuffer(vehiculoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearVehiculo(c)

	// Actualiza el vehículo recién creado
	var createdVehiculo models.Vehiculo
	json.Unmarshal(rec.Body.Bytes(), &createdVehiculo)
	id := strconv.Itoa(createdVehiculo.ID)
	vehiculo.Marca = "Honda"
	vehiculoJSON, _ = json.Marshal(vehiculo)
	req = httptest.NewRequest(http.MethodPut, "/vehiculos/"+id, bytes.NewBuffer(vehiculoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.ActualizarVehiculo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Honda")
	}
}

func TestEliminarVehiculo(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea un vehículo para luego eliminarlo
	vehiculo := models.Vehiculo{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Ano:    2020,
		Placa:  "ABC123",
	}
	vehiculoJSON, _ := json.Marshal(vehiculo)

	req := httptest.NewRequest(http.MethodPost, "/vehiculos", bytes.NewBuffer(vehiculoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearVehiculo(c)

	// Elimina el vehículo recién creado
	var createdVehiculo models.Vehiculo
	json.Unmarshal(rec.Body.Bytes(), &createdVehiculo)
	id := strconv.Itoa(createdVehiculo.ID)
	req = httptest.NewRequest(http.MethodDelete, "/vehiculos/"+id, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.EliminarVehiculo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Vehículo eliminado")
	}
}
