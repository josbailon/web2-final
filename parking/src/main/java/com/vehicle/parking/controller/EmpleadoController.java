package com.vehicle.parking.controller;

import com.vehicle.parking.model.Empleado;
import com.vehicle.parking.model.Empleado.EstadoEmpleado;
import com.vehicle.parking.service.EmpleadoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/empleados")
public class EmpleadoController {

    @Autowired
    private EmpleadoService empleadoService;

    @GetMapping
    public List<Empleado> getAllEmpleados() {
        return empleadoService.getAllEmpleados();
    }

    @PostMapping
    public Empleado createEmpleado(@RequestBody Empleado empleado) {
        return empleadoService.createEmpleado(empleado);
    }

    @GetMapping("/{id}")
    public ResponseEntity<Empleado> getEmpleadoById(@PathVariable Long id) {
        Empleado empleado = empleadoService.getEmpleadoById(id);
        if (empleado != null) {
            return ResponseEntity.ok(empleado);
        } else {
            return ResponseEntity.notFound().build();
        }
    }

    @GetMapping("/estado/{estado}")
    public List<Empleado> getEmpleadosByEstado(@PathVariable EstadoEmpleado estado) {
        return empleadoService.getEmpleadosByEstado(estado);
    }

    @PutMapping("/{id}")
    public ResponseEntity<Empleado> updateEmpleado(@PathVariable Long id, @RequestBody Empleado empleadoDetails) {
        Empleado updatedEmpleado = empleadoService.updateEmpleado(id, empleadoDetails);
        if (updatedEmpleado != null) {
            return ResponseEntity.ok(updatedEmpleado);
        } else {
            return ResponseEntity.notFound().build();
        }
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteEmpleado(@PathVariable Long id) {
        boolean deleted = empleadoService.deleteEmpleado(id);
        if (deleted) {
            return ResponseEntity.noContent().build();
        } else {
            return ResponseEntity.notFound().build();
        }
    }

    @PutMapping("/recover/{id}")
    public ResponseEntity<Empleado> recoverEmpleado(@PathVariable Long id) {
        Empleado recoveredEmpleado = empleadoService.recoverEmpleado(id);
        if (recoveredEmpleado != null) {
            return ResponseEntity.ok(recoveredEmpleado);
        } else {
            return ResponseEntity.notFound().build();
        }
    }
}
