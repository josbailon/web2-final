package models

import "time"

// SalidaVehiculo define la estructura de una salida de vehículo
type SalidaVehiculo struct {
    ID               int       `json:"id"`
    VehiculoID       int       `json:"vehiculo_id" validate:"required"`
    FechaSalida      time.Time `json:"fecha_salida" validate:"required"`
    TiempoEstacionado string   `json:"tiempo_estacionado" validate:"required"`
}
