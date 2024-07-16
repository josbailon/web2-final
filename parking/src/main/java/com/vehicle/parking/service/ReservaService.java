package com.vehicle.parking.service;

import com.vehicle.parking.model.Reserva;
import com.vehicle.parking.model.EstadoReserva;
import com.vehicle.parking.repository.ReservaRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ReservaService {

    @Autowired
    private ReservaRepository reservaRepository;

    public List<Reserva> getAllReservas() {
        return reservaRepository.findAll();
    }

    public Optional<Reserva> getReservaById(Long id) {
        return reservaRepository.findById(id);
    }

    public Reserva saveReserva(Reserva reserva) {
        return reservaRepository.save(reserva);
    }

    public void deleteReserva(Long id) {
        reservaRepository.deleteById(id);
    }

    public List<Reserva> getReservasByEstado(EstadoReserva estadoReserva) {
        return reservaRepository.findByEstadoReserva(estadoReserva);
    }

    public Reserva updateReserva(Long id, Reserva updatedReserva) {
        return reservaRepository.findById(id).map(reserva -> {
            reserva.setIdCliente(updatedReserva.getIdCliente());
            reserva.setIdTarifa(updatedReserva.getIdTarifa());
            reserva.setIdVehiculo(updatedReserva.getIdVehiculo());
            reserva.setFechaReserva(updatedReserva.getFechaReserva());
            reserva.setHoraEntrada(updatedReserva.getHoraEntrada());
            reserva.setHoraSalida(updatedReserva.getHoraSalida());
            reserva.setEstadoReserva(updatedReserva.getEstadoReserva());
            return reservaRepository.save(reserva);
        }).orElseGet(() -> {
            updatedReserva.setId(id);
            return reservaRepository.save(updatedReserva);
        });
    }
}
