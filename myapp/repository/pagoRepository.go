package repository

import (
    "myapp/db"
    "myapp/models"
)

// ObtenerTodosPagos devuelve todos los pagos registrados
func ObtenerTodosPagos() ([]models.Pago, error) {
    var pagos []models.Pago
    rows, err := db.Conectar().Query("SELECT id, salida_id, monto, fecha_pago FROM pago")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var pago models.Pago
        if err := rows.Scan(&pago.ID, &pago.SalidaID, &pago.Monto, &pago.FechaPago); err != nil {
            return nil, err
        }
        pagos = append(pagos, pago)
    }
    return pagos, nil
}

// ObtenerPagoPorID devuelve un pago específico por ID
func ObtenerPagoPorID(id int) (models.Pago, error) {
    var pago models.Pago
    row := db.Conectar().QueryRow("SELECT id, salida_id, monto, fecha_pago FROM pago WHERE id = $1", id)
    err := row.Scan(&pago.ID, &pago.SalidaID, &pago.Monto, &pago.FechaPago)
    return pago, err
}

// CrearPago inserta un nuevo pago en la base de datos
func CrearPago(pago models.Pago) error {
    _, err := db.Conectar().Exec("INSERT INTO pago (salida_id, monto, fecha_pago) VALUES ($1, $2, $3)",
        pago.SalidaID, pago.Monto, pago.FechaPago)
    return err
}

// ActualizarPago actualiza un pago existente
func ActualizarPago(pago models.Pago) error {
    _, err := db.Conectar().Exec("UPDATE pago SET salida_id=$1, monto=$2, fecha_pago=$3 WHERE id=$4",
        pago.SalidaID, pago.Monto, pago.FechaPago, pago.ID)
    return err
}

// EliminarPago elimina un pago de la base de datos
func EliminarPago(id int) error {
    _, err := db.Conectar().Exec("DELETE FROM pago WHERE id=$1", id)
    return err
}
