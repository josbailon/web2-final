package schema

import (
    "time"

    "github.com/graphql-go/graphql"
)

// Vehiculo representa un vehículo
type Vehiculo struct {
    ID     int    `json:"id"`
    Marca  string `json:"marca"`
    Modelo string `json:"modelo"`
    Ano    int    `json:"ano"`
    Placa  string `json:"placa"`
}

// SalidaVehiculo representa una salida de vehículo
type SalidaVehiculo struct {
    ID               int       `json:"id"`
    VehiculoID       int       `json:"vehiculo_id"`
    FechaSalida      time.Time `json:"fecha_salida"`
    TiempoEstacionado string    `json:"tiempo_estacionado"`
}

// Pago representa un pago
type Pago struct {
    ID        int       `json:"id"`
    SalidaID  int       `json:"salida_id"`
    Monto     float64   `json:"monto"`
    FechaPago time.Time `json:"fecha_pago"`
}

// Datos simulados
var Vehiculos = []Vehiculo{
    {ID: 1, Marca: "Toyota", Modelo: "Corolla", Ano: 2020, Placa: "ABC123"},
    {ID: 2, Marca: "Honda", Modelo: "Civic", Ano: 2019, Placa: "XYZ789"},
}

var Salidas = []SalidaVehiculo{
    {ID: 1, VehiculoID: 1, FechaSalida: time.Now(), TiempoEstacionado: "2 hours"},
    {ID: 2, VehiculoID: 2, FechaSalida: time.Now(), TiempoEstacionado: "3 hours"},
}

var Pagos = []Pago{
    {ID: 1, SalidaID: 1, Monto: 50.00, FechaPago: time.Now()},
    {ID: 2, SalidaID: 2, Monto: 75.00, FechaPago: time.Now()},
}

// Definir el objeto Vehiculo para GraphQL
var vehiculoType = graphql.NewObject(graphql.ObjectConfig{
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
        "ano": &graphql.Field{
            Type: graphql.Int,
        },
        "placa": &graphql.Field{
            Type: graphql.String,
        },
    },
})

// Definir el objeto SalidaVehiculo para GraphQL
var salidaVehiculoType = graphql.NewObject(graphql.ObjectConfig{
    Name: "SalidaVehiculo",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "vehiculo_id": &graphql.Field{
            Type: graphql.Int,
        },
        "fecha_salida": &graphql.Field{
            Type: graphql.DateTime,
        },
        "tiempo_estacionado": &graphql.Field{
            Type: graphql.String,
        },
    },
})

// Definir el objeto Pago para GraphQL
var pagoType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Pago",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "salida_id": &graphql.Field{
            Type: graphql.Int,
        },
        "monto": &graphql.Field{
            Type: graphql.Float,
        },
        "fecha_pago": &graphql.Field{
            Type: graphql.DateTime,
        },
    },
})

// Definir el root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "vehiculo": &graphql.Field{
            Type:        vehiculoType,
            Description: "Obtener un vehículo por ID",
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                id, ok := params.Args["id"].(int)
                if ok {
                    for _, vehiculo := range Vehiculos {
                        if vehiculo.ID == id {
                            return vehiculo, nil
                        }
                    }
                }
                return nil, nil
            },
        },
        "vehiculos": &graphql.Field{
            Type:        graphql.NewList(vehiculoType),
            Description: "Obtener todos los vehículos",
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return Vehiculos, nil
            },
        },
        "salidaVehiculo": &graphql.Field{
            Type:        salidaVehiculoType,
            Description: "Obtener una salida de vehículo por ID",
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                id, ok := params.Args["id"].(int)
                if ok {
                    for _, salida := range Salidas {
                        if salida.ID == id {
                            return salida, nil
                        }
                    }
                }
                return nil, nil
            },
        },
        "salidasVehiculo": &graphql.Field{
            Type:        graphql.NewList(salidaVehiculoType),
            Description: "Obtener todas las salidas de vehículos",
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return Salidas, nil
            },
        },
        "pago": &graphql.Field{
            Type:        pagoType,
            Description: "Obtener un pago por ID",
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.Int,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                id, ok := params.Args["id"].(int)
                if ok {
                    for _, pago := range Pagos {
                        if pago.ID == id {
                            return pago, nil
                        }
                    }
                }
                return nil, nil
            },
        },
        "pagos": &graphql.Field{
            Type:        graphql.NewList(pagoType),
            Description: "Obtener todos los pagos",
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                return Pagos, nil
            },
        },
    },
})

// Definir el esquema de GraphQL
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: rootQuery,
})
