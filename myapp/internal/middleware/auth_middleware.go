// middleware/auth_middleware.go

package middleware/auth_middleware

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/dgrijalva/jwt-go"
)

// CustomClaims define la estructura de los claims personalizados del token JWT
type CustomClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// JWTAuthMiddleware verifica la validez del token JWT y protege las rutas
func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Token JWT no proporcionado")
        }

        // Validar el token JWT
        jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte("your-secret-key"), nil  // Debes definir tu clave secreta aquí
        })

        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Token JWT inválido")
        }

        // Verificar los claims del token JWT
        claims, ok := jwtToken.Claims.(*CustomClaims)
        if !ok || !jwtToken.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "Token JWT no válido")
        }

        // Colocar los claims en el contexto para usarlos en el controlador
        c.Set("username", claims.Username)

        return next(c)
    }
}
