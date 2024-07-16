package com.vehicle.parking.service;

import com.vehicle.parking.model.Empleado;
import com.vehicle.parking.model.Empleado.EstadoEmpleado;
import com.vehicle.parking.repository.EmpleadoRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class EmpleadoService {

    @Autowired
    private EmpleadoRepository empleadoRepository;

    public List<Empleado> getAllEmpleados() {
        return empleadoRepository.findAll();
    }

    public Empleado createEmpleado(Empleado empleado) {
        empleado.setEstado(EstadoEmpleado.ACTIVO);
        return empleadoRepository.save(empleado);
    }

    public Empleado getEmpleadoById(Long id) {
        Optional<Empleado> empleado = empleadoRepository.findById(id);
        return empleado.orElse(null);
    }

    public List<Empleado> getEmpleadosByEstado(EstadoEmpleado estado) {
        return empleadoRepository.findByEstado(estado);
    }

    public Empleado updateEmpleado(Long id, Empleado empleadoDetails) {
        Optional<Empleado> empleadoOptional = empleadoRepository.findById(id);
        if (empleadoOptional.isPresent()) {
            Empleado empleado = empleadoOptional.get();
            empleado.setNombre(empleadoDetails.getNombre());
            empleado.setApellido(empleadoDetails.getApellido());
            empleado.setFechaIngreso(empleadoDetails.getFechaIngreso());
            empleado.setEmail(empleadoDetails.getEmail());
            empleado.setSalario(empleadoDetails.getSalario());
            return empleadoRepository.save(empleado);
        } else {
            return null;
        }
    }

    public boolean deleteEmpleado(Long id) {
        Optional<Empleado> empleadoOptional = empleadoRepository.findById(id);
        if (empleadoOptional.isPresent()) {
            Empleado empleado = empleadoOptional.get();
            empleado.setEstado(EstadoEmpleado.INACTIVO);
            empleadoRepository.save(empleado);
            return true;
        } else {
            return false;
        }
    }

    public Empleado recoverEmpleado(Long id) {
        Optional<Empleado> empleadoOptional = empleadoRepository.findById(id);
        if (empleadoOptional.isPresent()) {
            Empleado empleado = empleadoOptional.get();
            empleado.setEstado(EstadoEmpleado.ACTIVO);
            return empleadoRepository.save(empleado);
        } else {
            return null;
        }
    }
}
