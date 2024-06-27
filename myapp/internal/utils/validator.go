// utils/validation_utils.go

package utils/validation_utils

import (
    "github.com/go-playground/validator/v10"
)

// Validator estructura para el validador
type Validator struct {
    validator *validator.Validate
}

// NewValidator crea una nueva instancia del validador
func NewValidator() *Validator {
    return &Validator{validator: validator.New()}
}

// Validate realiza la validación de la estructura utilizando el validador
func (v *Validator) Validate(i interface{}) error {
    return v.validator.Struct(i)
}

// Función de ejemplo para usar el validador
func ExampleValidationUsage() {
    // Crear una instancia del validador
    validator := NewValidator()

    // Ejemplo de uso del validador para validar una estructura
    type User struct {
        FirstName string `validate:"required"`
        LastName  string `validate:"required"`
    }

    user := User{
        FirstName: "John",
        LastName:  "Doe",
    }

    // Validar la estructura del usuario utilizando el validador
    if err := validator.Validate(user); err != nil {
        // Manejar el error de validación
        // Por ejemplo: log.Printf("Error de validación: %v", err)
        return
    }

    // Continuar con la lógica de negocio si la validación es exitosa
    // Por ejemplo: userService.Create(user)
}
