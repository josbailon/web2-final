import axios from 'axios';
import { Client, Vehicle, Payment } from '@/types';

const apiClient = axios.create({
  baseURL: 'http://127.0.0.1:8000/api',
  headers: {
    'Content-Type': 'application/json',
    Authorization: `Bearer 2|Xyw4zcVbSjU453Ap6KhqRamAVhOhMj6r3jllbj8tc1cecff2`
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

  // Reportes
  getPdfReport() {
    return apiClient.get('/reports/clients/pdf', {
      responseType: 'blob'
    });
  }
};
