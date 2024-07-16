// middleware/validator.go

package middleware/validation

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/go-playground/validator/v10" // Importar el validador de Go Playground
)

// CustomValidator estructura para el validador personalizado
type CustomValidator struct {
    validator *validator.Validate
}

// NewValidator crea una nueva instancia de validador personalizado
func NewValidator() *CustomValidator {
    return &CustomValidator{validator: validator.New()}
}

// Validate realiza la validación de la estructura usando el validador personalizado
func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

// Middleware de validación que utiliza el validador personalizado
func Validate() echo.MiddlewareFunc {
    return middleware.ValidatorWithConfig(middleware.ValidatorConfig{
        Skipper: middleware.DefaultSkipper,
        Validator: &CustomValidator{
            validator: validator.New(),
        },
    })
}

// Ejemplo de estructura que se va a validar (puedes definir tus propias estructuras y reglas de validación)
type Vehicle struct {
    Brand  string `json:"brand" validate:"required"`
    Model  string `json:"model" validate:"required"`
    Type   string `json:"type" validate:"required"`
    Year   int    `json:"year" validate:"required"`
    Status string `json:"status" validate:"required"`
}

// Función de ejemplo para usar el validador personalizado
func ExampleValidatorUsage() {
    // Crear una instancia del validador personalizado
    cv := NewValidator()

    // Ejemplo de uso del validador para validar una estructura de vehículo
    vehicle := Vehicle{
        Brand:  "Toyota",
        Model:  "Corolla",
        Type:   "Sedan",
        Year:   2022,
        Status: "available",
    }

    // Validar la estructura del vehículo usando el validador personalizado
    if err := cv.Validate(vehicle); err != nil {
        // Manejar el error de validación
        // En este ejemplo, podrías devolver un error HTTP 400 Bad Request si la validación falla
        // Por ejemplo: return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        return
    }

    // Si la validación es exitosa, continuar con la lógica de negocio
    // Aquí podrías llamar a un servicio o función que procese la lógica de negocio relacionada con el vehículo
    err := ProcessVehicleBusinessLogic(vehicle)
    if err != nil {
        // Manejar errores de negocio según sea necesario
        // Por ejemplo, loguear el error, devolver un error HTTP 500 Internal Server Error, etc.
        //return echo.NewHTTPError(http.StatusInternalServerError, "Error al procesar la lógica de negocio")
        return
    }

    // Todo salió bien, responder con éxito (por ejemplo, código HTTP 200 OK)
    // Puedes devolver una respuesta JSON, un mensaje de éxito, etc.
    // return c.JSON(http.StatusOK, echo.Map{"message": "Operación exitosa"})
}

// Función ficticia para procesar la lógica de negocio relacionada con el vehículo
func ProcessVehicleBusinessLogic(vehicle Vehicle) error {
    // Aquí podrías implementar la lógica de negocio, como almacenar en la base de datos, realizar cálculos, etc.
    // Por ejemplo, guardar el vehículo en la base de datos
    // vehicleRepo.Save(vehicle)
    return nil
}
