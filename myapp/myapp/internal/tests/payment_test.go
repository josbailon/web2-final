// services/payment_service_test.go

package services_test/payment_service_test

import (
    "database/sql"
    "log"
    "os"
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
    ps     *services.PaymentService
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

    // Inicializar repositorio y servicio de pagos
    paymentRepo := repositories.NewPaymentRepository(db)
    ps = services.NewPaymentService(paymentRepo)

    // Enlazar rutas del servicio de pagos a Echo para pruebas de integración
    routes.InitPaymentRoutes(router, db)
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

func TestGetAllPaymentsHandler(t *testing.T) {
    // Preparar prueba de obtención de todos los pagos
    req := httptest.NewRequest(http.MethodGet, "/payments", nil)
    rec := httptest.NewRecorder()
    c := router.NewContext(req, rec)

    // Ejecutar el handler
    if assert.NoError(t, ps.GetAllPaymentsHandler(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)

        // Verificar la respuesta (puedes verificar la estructura del JSON devuelto)
        var payments []models.Payment
        if err := json.Unmarshal(rec.Body.Bytes(), &payments); err != nil {
            log.Fatalf("Error al decodificar la respuesta JSON: %v", err)
        }

        // Puedes realizar más aserciones según las expectativas de tu aplicación
        assert.NotEmpty(t, payments, "Se esperaban pagos no vacíos")
    }
}

func TestCreatePaymentHandler(t *testing.T) {
    // Preparar prueba de creación de pago
    payment := models.Payment{
        Amount:  100.00,
        Status:  "pending",
        Details: "Pago de prueba",
    }
    paymentJSON, err := json.Marshal(payment)
    assert.NoError(t, err)

    req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(paymentJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := router.NewContext(req, rec)

    // Ejecutar el handler
    if assert.NoError(t, ps.CreatePaymentHandler(c)) {
        assert.Equal(t, http.StatusCreated, rec.Code)

        // Verificar la respuesta (puedes verificar la estructura del JSON devuelto)
        var createdPayment models.Payment
        if err := json.Unmarshal(rec.Body.Bytes(), &createdPayment); err != nil {
            log.Fatalf("Error al decodificar la respuesta JSON: %v", err)
        }

        assert.NotEmpty(t, createdPayment.ID, "Se esperaba un ID de pago no vacío")
        assert.Equal(t, payment.Amount, createdPayment.Amount, "El monto del pago creado no coincide")
        // Puedes realizar más aserciones según las expectativas de tu aplicación
    }
}

// Puedes continuar añadiendo más pruebas para los otros handlers (GetPaymentByID, UpdatePayment, DeletePayment, etc.)

