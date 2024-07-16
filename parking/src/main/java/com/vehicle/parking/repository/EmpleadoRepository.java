package com.vehicle.parking.repository;

import com.vehicle.parking.model.Empleado;
import com.vehicle.parking.model.Empleado.EstadoEmpleado;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface EmpleadoRepository extends JpaRepository<Empleado, Long> {
    List<Empleado> findByEstado(EstadoEmpleado estado);
}
