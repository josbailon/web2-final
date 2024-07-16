package repository

import (
	"myapp/db"
	"myapp/models"
)

// ObtenerTodosVehiculos devuelve todos los vehículos de la base de datos
func ObtenerTodosVehiculos() ([]models.Vehiculo, error) {
	var vehiculos []models.Vehiculo
	rows, err := db.Conectar().Query("SELECT id, marca, modelo, ano, placa FROM vehiculo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var vehiculo models.Vehiculo
		if err := rows.Scan(&vehiculo.ID, &vehiculo.Marca, &vehiculo.Modelo, &vehiculo.Ano, &vehiculo.Placa); err != nil {
			return nil, err
		}
		vehiculos = append(vehiculos, vehiculo)
	}
	return vehiculos, nil
}

// ObtenerVehiculoPorID devuelve un vehículo específico por ID
func ObtenerVehiculoPorID(id int) (models.Vehiculo, error) {
	var vehiculo models.Vehiculo
	row := db.Conectar().QueryRow("SELECT id, marca, modelo, ano, placa FROM vehiculo WHERE id = $1", id)
	err := row.Scan(&vehiculo.ID, &vehiculo.Marca, &vehiculo.Modelo, &vehiculo.Ano, &vehiculo.Placa)
	return vehiculo, err
}

// CrearVehiculo inserta un nuevo vehículo en la base de datos
func CrearVehiculo(vehiculo models.Vehiculo) error {
	_, err := db.Conectar().Exec("INSERT INTO vehiculo (marca, modelo, ano, placa) VALUES ($1, $2, $3, $4)",
		vehiculo.Marca, vehiculo.Modelo, vehiculo.Ano, vehiculo.Placa)
	return err
}

// ActualizarVehiculo actualiza un vehículo existente en la base de datos
func ActualizarVehiculo(vehiculo models.Vehiculo) error {
	_, err := db.Conectar().Exec("UPDATE vehiculo SET marca=$1, modelo=$2, ano=$3, placa=$4 WHERE id=$5",
		vehiculo.Marca, vehiculo.Modelo, vehiculo.Ano, vehiculo.Placa, vehiculo.ID)
	return err
}

// EliminarVehiculo elimina un vehículo de la base de datos
func EliminarVehiculo(id int) error {
	_, err := db.Conectar().Exec("DELETE FROM vehiculo WHERE id=$1", id)
	return err
}
