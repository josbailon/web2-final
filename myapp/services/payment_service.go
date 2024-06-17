package services

import (
    "parking-management/models"
    "parking-management/repositories"
    "time"
)

func CalculatePayment(duration time.Duration) float64 {
    hours := duration.Hours()
    rate := 5.0 // Tarifa por hora
    return hours * rate
}

func CreatePayment(vehicleID uint, amount float64) (models.Payment, error) {
    payment := models.Payment{
        VehicleID: vehicleID,
        Amount:    amount,
    }
    err := repositories.CreatePayment(payment)
    return payment, err
}
