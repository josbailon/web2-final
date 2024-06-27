// graphql_controller.go

package controllers/graphql_controller

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/graphql-go/graphql"
)

// GraphQLHandler maneja las solicitudes GraphQL
func GraphQLHandler(schema graphql.Schema) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Parsear el cuerpo JSON de la solicitud
        req := map[string]interface{}{}
        if err := c.Bind(&req); err != nil {
            return err
        }

        // Ejecutar la consulta GraphQL
        result := graphql.Do(graphql.Params{
            Schema:        schema,
            RequestString: req["query"].(string),
            VariableValues: req["variables"].(map[string]interface{}),
        })

        // Verificar errores en la ejecución de la consulta
        if len(result.Errors) > 0 {
            return c.JSON(http.StatusBadRequest, result.Errors)
        }

        // Devolver la respuesta exitosa de la consulta GraphQL
        return c.JSON(http.StatusOK, result.Data)
    }
}
