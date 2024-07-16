// resolvers.go

package graphql/resolvers

import (
    "errors"
    "strconv"

    "github.com/graphql-go/graphql"
)

// Mock data para propósitos de ejemplo
var pagosData = map[int]*Pago{
    1: {ID: 1, Monto: 100.50, Descripcion: "Pago de ejemplo", Estado: "Pagado"},
    2: {ID: 2, Monto: 75.20, Descripcion: "Otro pago", Estado: "Pendiente"},
}

var vehiculosData = map[int]*Vehiculo{
    1: {ID: 1, Marca: "Toyota", Modelo: "Corolla", Año: 2022},
    2: {ID: 2, Marca: "Honda", Modelo: "Civic", Año: 2023},
}

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

// Función resolver para el campo "pago"
func resolvePago(params graphql.ResolveParams) (interface{}, error) {
    id, ok := params.Args["id"].(int)
    if !ok {
        return nil, errors.New("ID de pago no válido")
    }

    pago, found := pagosData[id]
    if !found {
        return nil, errors.New("Pago no encontrado")
    }

    return pago, nil
}

// Función resolver para el campo "vehiculo"
func resolveVehiculo(params graphql.ResolveParams) (interface{}, error) {
    id, ok := params.Args["id"].(int)
    if !ok {
        return nil, errors.New("ID de vehículo no válido")
    }

    vehiculo, found := vehiculosData[id]
    if !found {
        return nil, errors.New("Vehículo no encontrado")
    }

    return vehiculo, nil
}
