package repository

import (
	"myapp/db"
	"myapp/models"
)

// ObtenerTodasSalidas devuelve todas las salidas de vehículos
func ObtenerTodasSalidas() ([]models.SalidaVehiculo, error) {
	var salidas []models.SalidaVehiculo
	rows, err := db.Conectar().Query("SELECT id, vehiculo_id, fecha_salida, tiempo_estacionado FROM salida_vehiculo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var salida models.SalidaVehiculo
		if err := rows.Scan(&salida.ID, &salida.VehiculoID, &salida.FechaSalida, &salida.TiempoEstacionado); err != nil {
			return nil, err
		}
		salidas = append(salidas, salida)
	}
	return salidas, nil
}

// ObtenerSalidaPorID devuelve una salida de vehículo específica por ID
func ObtenerSalidaPorID(id int) (models.SalidaVehiculo, error) {
	var salida models.SalidaVehiculo
	row := db.Conectar().QueryRow("SELECT id, vehiculo_id, fecha_salida, tiempo_estacionado FROM salida_vehiculo WHERE id = $1", id)
	err := row.Scan(&salida.ID, &salida.VehiculoID, &salida.FechaSalida, &salida.TiempoEstacionado)
	return salida, err
}

// CrearSalida inserta una nueva salida de vehículo en la base de datos
func CrearSalida(salida models.SalidaVehiculo) error {
	_, err := db.Conectar().Exec("INSERT INTO salida_vehiculo (vehiculo_id, fecha_salida, tiempo_estacionado) VALUES ($1, $2, $3)",
		salida.VehiculoID, salida.FechaSalida, salida.TiempoEstacionado)
	return err
}

// ActualizarSalida actualiza una salida de vehículo existente
func ActualizarSalida(salida models.SalidaVehiculo) error {
	_, err := db.Conectar().Exec("UPDATE salida_vehiculo SET vehiculo_id=$1, fecha_salida=$2, tiempo_estacionado=$3 WHERE id=$4",
		salida.VehiculoID, salida.FechaSalida, salida.TiempoEstacionado, salida.ID)
	return err
}

// EliminarSalida elimina una salida de vehículo de la base de datos
func EliminarSalida(id int) error {
	_, err := db.Conectar().Exec("DELETE FROM salida_vehiculo WHERE id=$1", id)
	return err
}
