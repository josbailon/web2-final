package services

import (
    "golang.org/x/crypto/bcrypt"
    "parking-management/models"
    "parking-management/repositories"
    "parking-management/utils"
)

func RegisterUser(username, password string) (models.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return models.User{}, err
    }

    user := models.User{
        Username: username,
        Password: string(hashedPassword),
    }

    err = repositories.CreateUser(user)
    return user, err
}

func AuthenticateUser(username, password string) (string, error) {
    user, err := repositories.FindUserByUsername(username)
    if err != nil {
        return "", err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", err
    }

    token, err := utils.GenerateJWT(user.ID)
    return token, err
}
