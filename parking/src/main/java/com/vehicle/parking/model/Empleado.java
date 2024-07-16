package com.vehicle.parking.model;

import jakarta.persistence.*;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import java.util.Date;

@Entity
@Table(name = "empleados")
public class Empleado {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @NotNull
    @Size(min = 2, max = 50)
    private String nombre;

    @NotNull
    @Size(min = 2, max = 50)
    private String apellido;

    @Column(name = "fechaIngreso", nullable = false)
    @Temporal(TemporalType.DATE)
    private Date fechaIngreso;

    @Column(name = "email", nullable = false)
    @Email
    private String email;

    @Column(name = "salario", nullable = false)
    private Double salario;

    @Column(columnDefinition = "varchar(255) default 'ACTIVO' NOT NULL")
    @Enumerated(EnumType.STRING)
    private EstadoEmpleado estado;

    // Getters and Setters
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getApellido() {
        return apellido;
    }

    public void setApellido(String apellido) {
        this.apellido = apellido;
    }

    public Date getFechaIngreso() {
        return fechaIngreso;
    }

    public void setFechaIngreso(Date fechaIngreso) {
        this.fechaIngreso = fechaIngreso;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public Double getSalario() {
        return salario;
    }

    public void setSalario(Double salario) {
        this.salario = salario;
    }

    public EstadoEmpleado getEstado() {
        return estado;
    }

    public void setEstado(EstadoEmpleado estado) {
        this.estado = estado;
    }

    public enum EstadoEmpleado {
        ACTIVO,
        INACTIVO
    }
}