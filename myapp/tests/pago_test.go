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

func TestObtenerTodosPagos(t *testing.T) {
	initDB()
	e := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/pagos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.ObtenerTodosPagos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body.String())
	}
}

func TestCrearPago(t *testing.T) {
	initDB()
	e := setupRouter()

	pago := models.Pago{
		SalidaID:  1,
		Monto:     150.00,
		FechaPago: time.Now(),
	}
	pagoJSON, _ := json.Marshal(pago)

	req := httptest.NewRequest(http.MethodPost, "/pagos", bytes.NewBuffer(pagoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.CrearPago(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "150.00")
	}
}

func TestObtenerPago(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea un pago para luego obtenerlo
	pago := models.Pago{
		SalidaID:  1,
		Monto:     150.00,
		FechaPago: time.Now(),
	}
	pagoJSON, _ := json.Marshal(pago)

	req := httptest.NewRequest(http.MethodPost, "/pagos", bytes.NewBuffer(pagoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearPago(c)

	// Obtén el pago recién creado
	id := rec.Body.String() // Asume que la respuesta contiene el ID del pago
	req = httptest.NewRequest(http.MethodGet, "/pagos/"+id, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.ObtenerPago(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "150.00")
	}
}

func TestActualizarPago(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea un pago para luego actualizarlo
	pago := models.Pago{
		SalidaID:  1,
		Monto:     150.00,
		FechaPago: time.Now(),
	}
	pagoJSON, _ := json.Marshal(pago)

	req := httptest.NewRequest(http.MethodPost, "/pagos", bytes.NewBuffer(pagoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearPago(c)

	// Actualiza el pago recién creado
	id := rec.Body.String() // Asume que la respuesta contiene el ID del pago
	pago.Monto = 175.00
	pagoJSON, _ = json.Marshal(pago)
	req = httptest.NewRequest(http.MethodPut, "/pagos/"+id, bytes.NewBuffer(pagoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.ActualizarPago(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "175.00")
	}
}

func TestEliminarPago(t *testing.T) {
	initDB()
	e := setupRouter()

	// Primero, crea un pago para luego eliminarlo
	pago := models.Pago{
		SalidaID:  1,
		Monto:     150.00,
		FechaPago: time.Now(),
	}
	pagoJSON, _ := json.Marshal(pago)

	req := httptest.NewRequest(http.MethodPost, "/pagos", bytes.NewBuffer(pagoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handlers.CrearPago(c)

	// Elimina el pago recién creado
	id := rec.Body.String() // Asume que la respuesta contiene el ID del pago
	req = httptest.NewRequest(http.MethodDelete, "/pagos/"+id, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, handlers.EliminarPago(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Pago eliminado")
	}
}
