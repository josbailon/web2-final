package models

import "time"

// Pago define la estructura de un pago
type Pago struct {
    ID        int       `json:"id"`
    SalidaID  int       `json:"salida_id" validate:"required"`
    Monto     float64   `json:"monto" validate:"required,gt=0"`
    FechaPago time.Time `json:"fecha_pago" validate:"required"`
}
