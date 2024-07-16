// schema.go

package graphql/Schema

import (
    "github.com/graphql-go/graphql"
)

// Definición de tipos GraphQL (Pago y Vehiculo)
var pagoType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Pago",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "monto": &graphql.Field{
                Type: graphql.Float,
            },
            "descripcion": &graphql.Field{
                Type: graphql.String,
            },
            "estado": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

var vehiculoType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Vehiculo",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "marca": &graphql.Field{
                Type: graphql.String,
            },
            "modelo": &graphql.Field{
                Type: graphql.String,
            },
            "año": &graphql.Field{
                Type: graphql.Int,
            },
        },
    },
)

// RootQuery es el punto de entrada para todas las consultas GraphQL
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootQuery",
    Fields: graphql.Fields{
        "pago": &graphql.Field{
            Type: pagoType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.Int),
                },
            },
            Resolve: resolvePago,
        },
        "vehiculo": &graphql.Field{
            Type: vehiculoType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.Int),
                },
            },
            Resolve: resolveVehiculo,
        },
    },
})

// SchemaConfig contiene la configuración del esquema GraphQL
var SchemaConfig = graphql.SchemaConfig{
    Query: RootQuery,
}

// Schema representa el esquema GraphQL de la aplicación
var Schema, _ = graphql.NewSchema(SchemaConfig)
