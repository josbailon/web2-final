// models/pago.go

package models/pago

// Pago representa la estructura de datos de un pago
type Pago struct {
    ID          int
    Monto       float64
    Descripcion string
    Estado      string
   
}
