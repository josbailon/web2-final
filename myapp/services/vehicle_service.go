package services

import (
    "time"
    "parking-management/models"
    "parking-management/repositories"
)

func CalculateParkingTime(inTime, outTime time.Time) time.Duration {
    return outTime.Sub(inTime)
}

func ProcessVehicleExit(vehicleID uint, outTime time.Time) (models.Payment, error) {
    vehicle, err := repositories.FindVehicleByID(vehicleID)
    if err != nil {
        return models.Payment{}, err
    }

    vehicle.OutTime = outTime
    err = repositories.UpdateVehicle(vehicle)
    if err != nil {
        return models.Payment{}, err
    }

    duration := CalculateParkingTime(vehicle.InTime, vehicle.OutTime)
    amount := CalculatePayment(duration)

    payment, err := CreatePayment(vehicle.ID, amount)
    return payment, err
}

func GetVehicleByID(vehicleID uint) (models.Vehicle, error) {
    return repositories.FindVehicleByID(vehicleID)
}
