package models

// Vehiculo define la estructura de un vehículo
type Vehiculo struct {
    ID     int    `json:"id"`
    Marca  string `json:"marca" validate:"required"`
    Modelo string `json:"modelo" validate:"required"`
    Ano    int    `json:"ano" validate:"required,gte=1886"`
    Placa  string `json:"placa" validate:"required"`
}
