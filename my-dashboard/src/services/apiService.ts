// src/services/apiService.ts
import axios from 'axios';
import { Client, Vehicle, Payment, Empleado, Reserva } from '@/types';

const apiClient = axios.create({
  baseURL: process.env.VUE_APP_API_URL || 'http://localhost:8080/api', // Asegúrate de que esto apunte a tu URL de Spring Boot
  headers: {
    'Content-Type': 'application/json',
  }
});

export default {
  // Clientes
  getClients() {
    return apiClient.get<Client[]>('/clients');
  },
  getClientVehicles(clientId: number) {
    return apiClient.get<Vehicle[]>(`/clients/${clientId}/vehicles`);
  },
  addClient(client: Client) {
    return apiClient.post('/clients', client);
  },
  updateClient(client: Client) {
    return apiClient.put(`/clients/${client.id}`, client);
  },

  // Vehículos
  addVehicle(vehicle: Vehicle) {
    return apiClient.post('/vehicles', vehicle);
  },
  updateVehicle(vehicle: Vehicle) {
    return apiClient.put(`/vehicles/${vehicle.id}`, vehicle);
  },

  // Pagos
  getPayments() {
    return apiClient.get<Payment[]>('/payments');
  },

  // Empleados
  getEmpleados() {
    return apiClient.get<Empleado[]>('/empleados');
  },
  addEmpleado(empleado: Empleado) {
    return apiClient.post('/empleados', empleado);
  },
  updateEmpleado(empleado: Empleado) {
    return apiClient.put(`/empleados/${empleado.id}`, empleado);
  },
  deleteEmpleado(id: number) {
    return apiClient.delete(`/empleados/${id}`);
  },
  recoverEmpleado(id: number) {
    return apiClient.put(`/empleados/recover/${id}`);
  },

  // Reservas
  getReservas() {
    return apiClient.get<Reserva[]>('/reservas');
  },
  addReserva(reserva: Reserva) {
    return apiClient.post('/reservas', reserva);
  },
  updateReserva(reserva: Reserva) {
    return apiClient.put(`/reservas/${reserva.id}`, reserva);
  },
  deleteReserva(id: number) {
    return apiClient.delete(`/reservas/${id}`);
  },
  recoverReserva(id: number) {
    return apiClient.put(`/reservas/recover/${id}`);
  },

  // Reportes
  getEmpleadoPdfReport() {
    return apiClient.get('/reports/empleados/pdf', { responseType: 'blob' });
  },
  getReservaPdfReport() {
    return apiClient.get('/reports/reservas/pdf', { responseType: 'blob' });
  },
  getEmpleadoExcelReport() {
    return apiClient.get('/reports/empleados/excel', { responseType: 'blob' });
  },
  getReservaExcelReport() {
    return apiClient.get('/reports/reservas/excel', { responseType: 'blob' });
  }
};
