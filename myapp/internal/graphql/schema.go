// graphql/schema.go

package graphql/schema

import (
    "github.com/graphql-go/graphql"
    "project/internal/services"
)

var (
    // Define los tipos de datos GraphQL
    pagoType = graphql.NewObject(
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

    vehiculoType = graphql.NewObject(
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
                "tipoCombustible": &graphql.Field{
                    Type: graphql.String,
                },
            },
        },
    )

    // Define las consultas disponibles en GraphQL
    rootQuery = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Query",
            Fields: graphql.Fields{
                "pago": &graphql.Field{
                    Type: pagoType,
                    Args: graphql.FieldConfigArgument{
                        "id": &graphql.ArgumentConfig{
                            Type: graphql.NewNonNull(graphql.Int),
                        },
                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        // Aquí se implementa la lógica para resolver la consulta de pago por ID
                        id, ok := p.Args["id"].(int)
                        if !ok {
                            return nil, nil
                        }
                        // Llamar al servicio de pago para obtener el pago por ID
                        payment, err := paymentService.GetPaymentByID(id)
                        if err != nil {
                            return nil, err
                        }
                        return payment, nil
                    },
                },
                "vehiculo": &graphql.Field{
                    Type: vehiculoType,
                    Args: graphql.FieldConfigArgument{
                        "id": &graphql.ArgumentConfig{
                            Type: graphql.NewNonNull(graphql.Int),
                        },
                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        // Aquí se implementa la lógica para resolver la consulta de vehículo por ID
                        id, ok := p.Args["id"].(int)
                        if !ok {
                            return nil, nil
                        }
                        // Llamar al servicio de vehículo para obtener el vehículo por ID
                        vehicle, err := vehicleService.GetVehicleByID(id)
                        if err != nil {
                            return nil, err
                        }
                        return vehicle, nil
                    },
                },
            },
        },
    )

    // Define el esquema GraphQL
    Schema, _ = graphql.NewSchema(
        graphql.SchemaConfig{
            Query: rootQuery,
        },
    )

    // Servicios para interactuar con datos
    var paymentService = services.NewPaymentService()
    var vehicleService = services.NewVehicleService()
)
