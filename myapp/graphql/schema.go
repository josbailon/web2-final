package graphql

import (
    "github.com/graphql-go/graphql"
)

var VehicleType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Vehicle",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.ID,
        },
        "plate": &graphql.Field{
            Type: graphql.String,
        },
        "in_time": &graphql.Field{
            Type: graphql.DateTime,
        },
        "out_time": &graphql.Field{
            Type: graphql.DateTime,
        },
	)