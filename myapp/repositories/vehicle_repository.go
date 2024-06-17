package repositories

import (
    "parking-management/database"
    "parking-management/models"
)

func FindVehicleByID(id uint) (models.Vehicle, error) {
    var vehicle models.Vehicle
    result := database.DB.First(&vehicle, id)
    return vehicle, result.Error
}

func UpdateVehicle(vehicle models.Vehicle) error {
    result := database.DB.Save(&vehicle)
    return result.Error
}
