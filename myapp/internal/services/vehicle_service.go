// services/vehicle_service.go

package services/vehicle_service

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "project/models"
    "project/repositories"
    "strconv"
)

// VehicleService representa el servicio para manejar operaciones relacionadas con vehículos
type VehicleService struct {
    VehicleRepo *repositories.VehicleRepository
}

// NewVehicleService crea una nueva instancia de VehicleService
func NewVehicleService(vr *repositories.VehicleRepository) *VehicleService {
    return &VehicleService{
        VehicleRepo: vr,
    }
}

// GetAllVehiclesHandler maneja la solicitud para obtener todos los vehículos
func (vs *VehicleService) GetAllVehiclesHandler(c echo.Context) error {
    vehicles, err := vs.VehicleRepo.GetAllVehicles()
    if err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(200, vehicles)
}

// GetVehicleByIDHandler maneja la solicitud para obtener un vehículo por su ID
func (vs *VehicleService) GetVehicleByIDHandler(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(400, map[string]string{"error": "ID de vehículo inválido"})
    }
    vehicle, err := vs.VehicleRepo.GetVehicleByID(id)
    if err != nil {
        return c.JSON(404, map[string]string{"error": "Vehículo no encontrado"})
    }
    return c.JSON(200, vehicle)
}

// CreateVehicleHandler maneja la solicitud para crear un nuevo vehículo
func (vs *VehicleService) CreateVehicleHandler(c echo.Context) error {
    var vehicle models.Vehicle
    if err := c.Bind(&vehicle); err != nil {
        return c.JSON(400, map[string]string{"error": "Datos de vehículo inválidos"})
    }
    if err := vs.VehicleRepo.CreateVehicle(&vehicle); err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(201, vehicle)
}

// UpdateVehicleHandler maneja la solicitud para actualizar un vehículo existente
func (vs *VehicleService) UpdateVehicleHandler(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(400, map[string]string{"error": "ID de vehículo inválido"})
    }
    var vehicle models.Vehicle
    if err := c.Bind(&vehicle); err != nil {
        return c.JSON(400, map[string]string{"error": "Datos de vehículo inválidos"})
    }
    vehicle.ID = id
    if err := vs.VehicleRepo.UpdateVehicle(&vehicle); err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(200, vehicle)
}

// DeleteVehicleHandler maneja la solicitud para eliminar un vehículo por su ID
func (vs *VehicleService) DeleteVehicleHandler(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(400, map[string]string{"error": "ID de vehículo inválido"})
    }
    if err := vs.VehicleRepo.DeleteVehicle(id); err != nil {
        return c.JSON(500, map[string]string{"error": err.Error()})
    }
    return c.JSON(204, nil)
}
