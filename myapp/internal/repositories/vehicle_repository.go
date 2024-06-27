// vehicle_repository.go

package repositories/vehicle_repository

import (
    "database/sql"
    "fmt"
    "project/models"  // Importa los modelos necesarios
)

// VehicleRepository representa un repositorio para el modelo Vehicle
type VehicleRepository struct {
    DB *sql.DB  // Conexión a la base de datos
}

// NewVehicleRepository crea una nueva instancia de VehicleRepository
func NewVehicleRepository(db *sql.DB) *VehicleRepository {
    return &VehicleRepository{
        DB: db,
    }
}

// GetAllVehicles obtiene todos los vehículos desde la base de datos
func (vr *VehicleRepository) GetAllVehicles() ([]models.Vehicle, error) {
    query := "SELECT id, make, model, year FROM vehicles"
    rows, err := vr.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error al consultar vehículos: %v", err)
    }
    defer rows.Close()

    var vehicles []models.Vehicle
    for rows.Next() {
        var vehicle models.Vehicle
        if err := rows.Scan(&vehicle.ID, &vehicle.Make, &vehicle.Model, &vehicle.Year); err != nil {
            return nil, fmt.Errorf("error al escanear fila de vehículos: %v", err)
        }
        vehicles = append(vehicles, vehicle)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error en las filas de resultados de vehículos: %v", err)
    }

    return vehicles, nil
}

// GetVehicleByID obtiene un vehículo por su ID desde la base de datos
func (vr *VehicleRepository) GetVehicleByID(id int) (*models.Vehicle, error) {
    query := "SELECT id, make, model, year FROM vehicles WHERE id = $1"
    var vehicle models.Vehicle
    row := vr.DB.QueryRow(query, id)
    err := row.Scan(&vehicle.ID, &vehicle.Make, &vehicle.Model, &vehicle.Year)
    if err != nil {
        return nil, fmt.Errorf("error al obtener vehículo: %v", err)
    }

    return &vehicle, nil
}

// CreateVehicle crea un nuevo vehículo en la base de datos
func (vr *VehicleRepository) CreateVehicle(vehicle *models.Vehicle) error {
    query := "INSERT INTO vehicles (make, model, year) VALUES ($1, $2, $3) RETURNING id"
    err := vr.DB.QueryRow(query, vehicle.Make, vehicle.Model, vehicle.Year).Scan(&vehicle.ID)
    if err != nil {
        return fmt.Errorf("error al crear vehículo: %v", err)
    }

    return nil
}

// UpdateVehicle actualiza un vehículo existente en la base de datos
func (vr *VehicleRepository) UpdateVehicle(vehicle *models.Vehicle) error {
    query := "UPDATE vehicles SET make = $1, model = $2, year = $3 WHERE id = $4"
    _, err := vr.DB.Exec(query, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.ID)
    if err != nil {
        return fmt.Errorf("error al actualizar vehículo: %v", err)
    }

    return nil
}

// DeleteVehicle elimina un vehículo por su ID de la base de datos
func (vr *VehicleRepository) DeleteVehicle(id int) error {
    query := "DELETE FROM vehicles WHERE id = $1"
    result, err := vr.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error al eliminar vehículo: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al obtener filas afectadas al eliminar vehículo: %v", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("vehículo no encontrado")
    }

    return nil
}
