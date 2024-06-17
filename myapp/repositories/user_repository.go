package repositories

import (
    "parking-management/database"
    "parking-management/models"
)

func FindUserByUsername(username string) (models.User, error) {
    var user models.User
    result := database.DB.Where("username = ?", username).First(&user)
    return user, result.Error
}

func CreateUser(user models.User) error {
    result := database.DB.Create(&user)
    return result.Error
}
