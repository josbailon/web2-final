package main_test/aplicación

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert" // Librería de aserciones para testing
)

// TestGETVehicles prueba la ruta GET /vehicles
func TestGETVehicles(t *testing.T) {
    // Crear una instancia de Echo
    e := echo.New()

    // Configurar el contexto de prueba
    req := httptest.NewRequest(http.MethodGet, "/vehicles", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Ejecutar el controlador
    if assert.NoError(t, GetVehiclesHandler(c)) {
        // Verificar el código de estado HTTP
        assert.Equal(t, http.StatusOK, rec.Code)

        // Verificar el cuerpo de la respuesta (opcional)
        var response []map[string]interface{}
        err := json.Unmarshal(rec.Body.Bytes(), &response)
        assert.NoError(t, err)


    }
}

// TestPOSTVehicle prueba la ruta POST /vehicles
func TestPOSTVehicle(t *testing.T) {
    // Crear una instancia de Echo
    e := echo.New()

    // Configurar el cuerpo de la solicitud
    requestBody := map[string]interface{}{
        "brand":  "Toyota",
        "model":  "Corolla",
        "type":   "Sedan",
        "year":   2022,
        "status": "available",
    }
    requestBodyBytes, _ := json.Marshal(requestBody)

    // Configurar el contexto de prueba
    req := httptest.NewRequest(http.MethodPost, "/vehicles", bytes.NewBuffer(requestBodyBytes))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Ejecutar el controlador
    if assert.NoError(t, PostVehicleHandler(c)) {
        // Verificar el código de estado HTTP
        assert.Equal(t, http.StatusCreated, rec.Code)

        // Verificar el cuerpo de la respuesta (opcional)
        var response map[string]interface{}
        err := json.Unmarshal(rec.Body.Bytes(), &response)
        assert.NoError(t, err)

    }
}

// Ejemplo de controladores para los endpoints de ejemplo (debes implementarlos en tu aplicación)
func GetVehiclesHandler(c echo.Context) error {
    // Lógica para obtener todos los vehículos (ejemplo)
    vehicles := []map[string]interface{}{
        {"id": 1, "brand": "Toyota", "model": "Corolla"},
        {"id": 2, "brand": "Honda", "model": "Civic"},
    }
    return c.JSON(http.StatusOK, vehicles)
}

func PostVehicleHandler(c echo.Context) error {
    // Lógica para crear un nuevo vehículo (ejemplo)
    var vehicle map[string]interface{}
    if err := c.Bind(&vehicle); err != nil {
        return err
    }
    // Aquí normalmente guardarías el vehículo en la base de datos u otro almacenamiento
    vehicle["id"] = 1 // Simulación de ID generado
    return c.JSON(http.StatusCreated, vehicle)
}
