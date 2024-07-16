// services/payment_service.go

package services/payment_service

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "project/models"
    "project/repositories"
    "strconv"
)

// PaymentService representa el servicio para manejar operaciones relacionadas con pagos
type PaymentService struct {
    PaymentRepo *repositories.PaymentRepository
}

// NewPaymentService crea una nueva instancia de PaymentService
func NewPaymentService(pr *repositories.PaymentRepository) *PaymentService {
    return &PaymentService{
        PaymentRepo: pr,
    }
}

// GetAllPaymentsHandler maneja la solicitud para obtener todos los pagos
func (ps *PaymentService) GetAllPaymentsHandler(c echo.Context) error {
    payments, err := ps.PaymentRepo.GetAllPayments()
    if err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(200, payments)
}

// GetPaymentByIDHandler maneja la solicitud para obtener un pago por su ID
func (ps *PaymentService) GetPaymentByIDHandler(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(400, map[string]string{"error": "ID de pago inválido"})
    }
    payment, err := ps.PaymentRepo.GetPaymentByID(id)
    if err != nil {
        return c.JSON(404, map[string]string{"error": "Pago no encontrado"})
    }
    return c.JSON(200, payment)
}

// CreatePaymentHandler maneja la solicitud para crear un nuevo pago
func (ps *PaymentService) CreatePaymentHandler(c echo.Context) error {
    var payment models.Payment
    if err := c.Bind(&payment); err != nil {
        return c.JSON(400, map[string]string{"error": "Datos de pago inválidos"})
    }
    if err := ps.PaymentRepo.CreatePayment(&payment); err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(201, payment)
}

// UpdatePaymentHandler maneja la solicitud para actualizar un pago existente
func (ps *PaymentService) UpdatePaymentHandler(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(400, map[string]string{"error": "ID de pago inválido"})
    }
    var payment models.Payment
    if err := c.Bind(&payment); err != nil {
        return c.JSON(400, map[string]string{"error": "Datos de pago inválidos"})
    }
    payment.ID = id
    if err := ps.PaymentRepo.UpdatePayment(&payment); err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(200, payment)
}

// DeletePaymentHandler maneja la solicitud para eliminar un pago por su ID
func (ps *PaymentService) DeletePaymentHandler(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(400, map[string]string{"error": "ID de pago inválido"})
    }
    if err := ps.PaymentRepo.DeletePayment(id); err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(204, nil)
}
