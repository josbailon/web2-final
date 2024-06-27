// services/vehicle_service_test.go

package services_test/vehicle_service_test

import (
    "bytes"
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "net/http/httptest"
    "os"
    "strconv"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
    "project/models"
    "project/repositories"
    "project/services"
)

var (
    db     *sql.DB
    router *echo.Echo
    vs     *services.VehicleService
)

func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}

func setup() {
    // Aquí puedes inicializar una conexión de prueba a la base de datos u otras configuraciones necesarias para las pruebas
    db = initDB() // Función ficticia para inicializar una base de datos de prueba
    router = echo.New()

    // Inicializar repositorio y servicio de vehículos
    vehicleRepo := repositories.NewVehicleRepository(db)
    vs = services.NewVehicleService(vehicleRepo)

    // Enlazar rutas del servicio de vehículos a Echo para pruebas de integración
    routes.InitVehicleRoutes(router, db)
}

func teardown() {
    // Aquí puedes realizar limpieza después de las pruebas, como cerrar conexiones de base de datos o liberar recursos
    db.Close()
}

// Función ficticia para inicializar una base de datos de prueba
func initDB() *sql.DB {
    // Inicializar base de datos de prueba
    // Retorna una conexión de base de datos simulada para pruebas
    return nil
}

func TestGetVehicleByIDHandler(t *testing.T) {
    // Preparar prueba de obtención de vehículo por ID
    // Supongamos que el ID del vehículo a probar es 1
    id := 1
    req := httptest.NewRequest(http.MethodGet, "/vehicles/"+strconv.Itoa(id), nil)
    rec := httptest.NewRecorder()
    c := router.NewContext(req, rec)
    c.SetPath("/vehicles/:id")
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(id))

    // Ejecutar el handler
    if assert.NoError(t, vs.GetVehicleByIDHandler(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)

        // Verificar la respuesta (puedes verificar la estructura del JSON devuelto)
        var vehicle models.Vehicle
        if err := json.Unmarshal(rec.Body.Bytes(), &vehicle); err != nil {
            log.Fatalf("Error al decodificar la respuesta JSON: %v", err)
        }

        // Verificar que el vehículo retornado tenga el ID correcto
        assert.Equal(t, id, vehicle.ID, "El ID del vehículo obtenido no coincide")
    }
}

func TestUpdateVehicleHandler(t *testing.T) {
    // Preparar prueba de actualización de vehículo
    // Supongamos que el ID del vehículo a actualizar es 1
    id := 1
    updatedVehicle := models.Vehicle{
        ID:     id,
        Brand:  "Honda",
        Model:  "Civic",
        Type:   "Sedan",
        Year:   2023,
        Status: "available",
    }
    updatedVehicleJSON, err := json.Marshal(updatedVehicle)
    assert.NoError(t, err)

    req := httptest.NewRequest(http.MethodPut, "/vehicles/"+strconv.Itoa(id), bytes.NewReader(updatedVehicleJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := router.NewContext(req, rec)
    c.SetPath("/vehicles/:id")
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(id))

    // Ejecutar el handler
    if assert.NoError(t, vs.UpdateVehicleHandler(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)

        // Verificar la respuesta (puedes verificar la estructura del JSON devuelto)
        var returnedVehicle models.Vehicle
        if err := json.Unmarshal(rec.Body.Bytes(), &returnedVehicle); err != nil {
            log.Fatalf("Error al decodificar la respuesta JSON: %v", err)
        }

        // Verificar que el vehículo retornado tenga los datos actualizados
        assert.Equal(t, updatedVehicle.Brand, returnedVehicle.Brand, "La marca del vehículo actualizado no coincide")
        // Puedes realizar más aserciones según las expectativas de tu aplicación
    }
}

func TestDeleteVehicleHandler(t *testing.T) {
    // Preparar prueba de eliminación de vehículo
    // Supongamos que el ID del vehículo a eliminar es 1
    id := 1
    req := httptest.NewRequest(http.MethodDelete, "/vehicles/"+strconv.Itoa(id), nil)
    rec := httptest.NewRecorder()
    c := router.NewContext(req, rec)
    c.SetPath("/vehicles/:id")
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(id))

    // Ejecutar el handler
    if assert.NoError(t, vs.DeleteVehicleHandler(c)) {
        assert.Equal(t, http.StatusNoContent, rec.Code)

        // Verificar que no haya contenido en la respuesta (código 204)
        assert.Empty(t, rec.Body.String(), "Se esperaba una respuesta vacía")
    }
}



