package controllers

import (
    "net/http"
    "parking-management/services"
    "github.com/labstack/echo/v4"
)

type LoginRequest struct {
    Username string `xml:"username"`
    Password string `xml:"password"`
}

type RegisterRequest struct {
    Username string `xml:"username"`
    Password string `xml:"password"`
}

func Login(c echo.Context) error {
    var request LoginRequest
    if err := c.Bind(&request); err != nil {
        return c.XML(http.StatusBadRequest, err)
    }

    token, err := services.AuthenticateUser(request.Username, request.Password)
    if err != nil {
        return c.XML(http.StatusUnauthorized, err)
    }

    return c.XML(http.StatusOK, echo.Map{"token": token})
}

func Register(c echo.Context) error {
    var request RegisterRequest
    if err := c.Bind(&request); err != nil {
        return c.XML(http.StatusBadRequest, err)
    }

    user, err := services.RegisterUser(request.Username, request.Password)
    if err != nil {
        return c.XML(http.StatusInternalServerError, err)
    }

    return c.XML(http.StatusOK, user)
}
