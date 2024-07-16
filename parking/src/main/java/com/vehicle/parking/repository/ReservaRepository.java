package com.vehicle.parking.repository;

import com.vehicle.parking.model.Reserva;
import com.vehicle.parking.model.EstadoReserva;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface ReservaRepository extends JpaRepository<Reserva, Long> {
    List<Reserva> findByEstadoReserva(EstadoReserva estadoReserva);
}
