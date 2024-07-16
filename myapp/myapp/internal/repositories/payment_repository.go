// payment_repository.go

package repositories/payment_repository

import (
    "database/sql"
    "fmt"
    "project/models"  // Importa los modelos necesarios
)

// PaymentRepository representa un repositorio para el modelo Payment
type PaymentRepository struct {
    DB *sql.DB  // Conexión a la base de datos
}

// NewPaymentRepository crea una nueva instancia de PaymentRepository
func NewPaymentRepository(db *sql.DB) *PaymentRepository {
    return &PaymentRepository{
        DB: db,
    }
}

// GetAllPayments obtiene todos los pagos desde la base de datos
func (pr *PaymentRepository) GetAllPayments() ([]models.Payment, error) {
    query := "SELECT id, amount, description, status FROM payments"
    rows, err := pr.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error al consultar pagos: %v", err)
    }
    defer rows.Close()

    var payments []models.Payment
    for rows.Next() {
        var payment models.Payment
        if err := rows.Scan(&payment.ID, &payment.Amount, &payment.Description, &payment.Status); err != nil {
            return nil, fmt.Errorf("error al escanear fila de pagos: %v", err)
        }
        payments = append(payments, payment)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error en las filas de resultados de pagos: %v", err)
    }

    return payments, nil
}

// GetPaymentByID obtiene un pago por su ID desde la base de datos
func (pr *PaymentRepository) GetPaymentByID(id int) (*models.Payment, error) {
    query := "SELECT id, amount, description, status FROM payments WHERE id = $1"
    var payment models.Payment
    row := pr.DB.QueryRow(query, id)
    err := row.Scan(&payment.ID, &payment.Amount, &payment.Description, &payment.Status)
    if err != nil {
        return nil, fmt.Errorf("error al obtener pago: %v", err)
    }

    return &payment, nil
}

// CreatePayment crea un nuevo pago en la base de datos
func (pr *PaymentRepository) CreatePayment(payment *models.Payment) error {
    query := "INSERT INTO payments (amount, description, status) VALUES ($1, $2, $3) RETURNING id"
    err := pr.DB.QueryRow(query, payment.Amount, payment.Description, payment.Status).Scan(&payment.ID)
    if err != nil {
        return fmt.Errorf("error al crear pago: %v", err)
    }

    return nil
}

// UpdatePayment actualiza un pago existente en la base de datos
func (pr *PaymentRepository) UpdatePayment(payment *models.Payment) error {
    query := "UPDATE payments SET amount = $1, description = $2, status = $3 WHERE id = $4"
    _, err := pr.DB.Exec(query, payment.Amount, payment.Description, payment.Status, payment.ID)
    if err != nil {
        return fmt.Errorf("error al actualizar pago: %v", err)
    }

    return nil
}

// DeletePayment elimina un pago por su ID de la base de datos
func (pr *PaymentRepository) DeletePayment(id int) error {
    query := "DELETE FROM payments WHERE id = $1"
    result, err := pr.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error al eliminar pago: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al obtener filas afectadas al eliminar pago: %v", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("pago no encontrado")
    }

    return nil
}
