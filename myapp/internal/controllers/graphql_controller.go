package controllers/graphql

import (
    "github.com/labstack/echo/v4"
    "github.com/graphql-go/handler"
    "project/internal/graphql"
)

func GraphQLHandler() echo.HandlerFunc {
    h := handler.New(&handler.Config{
        Schema:   graphql.Schema,
        Pretty:   true,
        GraphiQL: true,
    })

    return func(c echo.Context) error {
        h.ServeHTTP(c.Response(), c.Request())
        return nil
    }
}
