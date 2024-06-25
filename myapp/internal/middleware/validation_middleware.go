// middleware/validation_middleware.go

package middleware/validation_middleware

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

// ValidateInputMiddleware realiza la validación de las entradas en las solicitudes
func ValidateInputMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        req := c.Request()

        // Realizar la validación según la estructura de tu entrada
        var input struct {
            Field1 string `json:"field1" validate:"required"`
            Field2 int    `json:"field2" validate:"gte=0"`
            // Define más campos según tus necesidades
        }

        // Decodificar y validar la entrada
        if err := c.Bind(&input); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Entrada no válida")
        }
        if err := validate.Struct(input); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Entrada no válida")
        }

        // Si la validación es exitosa, colocar la entrada validada en el contexto
        c.Set("input", input)

        return next(c)
    }
}
