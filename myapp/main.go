package main

import (
    "net/http"
    "myapp/schema"

    "github.com/graphql-go/graphql"
    "github.com/labstack/echo/v4"
    "myapp/db"
)

func GraphQLHandler(c echo.Context) error {
    var params struct {
        Query string `json:"query"`
    }
    if err := c.Bind(&params); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    result := graphql.Do(graphql.Params{
        Schema:        schema.Schema,
        RequestString: params.Query,
    })
    if len(result.Errors) > 0 {
        return c.JSON(http.StatusBadRequest, result.Errors)
    }
    return c.JSON(http.StatusOK, result)
}

func main() {
    // Inicializar la base de datos
    db.Init()

    e := echo.New()

    e.POST("/graphql", GraphQLHandler)

    e.Logger.Fatal(e.Start(":8080"))
}
