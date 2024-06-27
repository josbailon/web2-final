// middleware/jwt.go

package middleware/jwt

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
)

var (
    jwtSecret = []byte("your_jwt_secret_key") // Clave secreta para firmar JWT (debería ser almacenada de forma segura)
)

// CustomClaims estructura para los claims personalizados en JWT
type CustomClaims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

// GenerarTokenJWT genera un nuevo token JWT con los claims especificados
func GenerarTokenJWT(userID, username, role string) (string, error) {
    // Crear claims personalizados
    claims := CustomClaims{
        userID,
        username,
        role,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
            Issuer:    "your_app_name",                        // Emisor del token (nombre de tu aplicación)
        },
    }

    // Crear token JWT con los claims y la clave secreta
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Middleware de autenticación JWT
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Token JWT no proporcionado")
        }

        // Validar el token JWT
        claims := &CustomClaims{}
        tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                return echo.NewHTTPError(http.StatusUnauthorized, "Token JWT no válido")
            }
            return echo.NewHTTPError(http.StatusBadRequest, "Token JWT no válido")
        }
        if !tkn.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "Token JWT no válido")
        }

        // Agregar los claims al contexto
        c.Set("user", claims)

        return next(c)
    }
}

// Ejemplo de uso del middleware en una ruta
func ExampleJwtUsage() {
    e := echo.New()

    // Middleware JWT protegido
    e.GET("/ruta-protegida", func(c echo.Context) error {
        user := c.Get("user").(*CustomClaims)
        return c.String(http.StatusOK, "Hola "+user.Username)
    }, JwtMiddleware)

    e.Logger.Fatal(e.Start(":8080"))
}
