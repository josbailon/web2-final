<template>
  <div class="dashboard">
    <h1>Dashboard del Parqueadero</h1>
    <div class="buttons-section">
      <button @click="fetchClients">Ver Clientes</button>
      <button @click="generatePdf">Generar PDF</button>
      <button @click="fetchClientsAndVehicles">Ver Clientes y Vehículos</button>
      <button @click="fetchPayments">Ver Pagos</button>
      <button @click="showPostClientForm">Post Cliente</button>
      <button @click="showPostVehicleForm">Post Vehículo</button>
    </div>
    <div class="content-section">
      <div v-if="view === 'clients'">
        <h2>Clientes</h2>
        <table>
          <thead>
            <tr>
              <th>Cliente</th>
              <th>Correo</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="client in clients" :key="client.id">
              <td>{{ client.name }}</td>
              <td>{{ client.email }}</td>
              <td>
                <button @click="viewClientVehicles(client)">Ver Vehículos</button>
                <button @click="editClient(client)">Editar</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="view === 'clientsAndVehicles'">
        <h2>Clientes y Vehículos</h2>
        <table>
          <thead>
            <tr>
              <th>Cliente</th>
              <th>Correo</th>
              <th>Modelo de Vehículo</th>
              <th>Placa</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="client in clients" :key="client.id">
              <td>{{ client.name }}</td>
              <td>{{ client.email }}</td>
              <td>
                <ul>
                  <li v-for="vehicle in client.vehicles" :key="vehicle.id">
                    {{ vehicle.model }}
                  </li>
                </ul>
              </td>
              <td>
                <ul>
                  <li v-for="vehicle in client.vehicles" :key="vehicle.id">
                    {{ vehicle.license_plate }}
                  </li>
                </ul>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="view === 'payments'">
        <h2>Pagos</h2>
        <table>
          <thead>
            <tr>
              <th>Monto</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="payment in payments" :key="payment.id">
              <td>{{ payment.amount }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="view === 'postClient'">
        <h2>Post Cliente</h2>
        <form @submit.prevent="createOrUpdateClient">
          <input v-model="newClient.name" placeholder="Nombre del Cliente" required />
          <input v-model="newClient.email" placeholder="Correo del Cliente" required />
          <div class="buttons">
            <button type="submit">{{ isEditingClient ? 'Actualizar' : 'Agregar' }} Cliente</button>
            <button type="button" v-if="isEditingClient" @click="cancelEditClient">Cancelar</button>
          </div>
        </form>
      </div>
      <div v-if="view === 'postVehicle'">
        <h2>Post Vehículo</h2>
        <form @submit.prevent="createOrUpdateVehicle">
          <input v-model="newVehicle.client_id" placeholder="ID del Cliente" required />
          <input v-model="newVehicle.license_plate" placeholder="Placa del Vehículo" required />
          <input v-model="newVehicle.model" placeholder="Modelo del Vehículo" required />
          <div class="buttons">
            <button type="submit">{{ isEditingVehicle ? 'Actualizar' : 'Agregar' }} Vehículo</button>
            <button type="button" v-if="isEditingVehicle" @click="cancelEditVehicle">Cancelar</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import apiService from '@/services/apiService';
import pusher from '@/pusher';
import { Client, Vehicle, Payment } from '@/types';

export default defineComponent({
  data() {
    return {
      clients: [] as Client[],
      payments: [] as Payment[],
      view: '',
      newClient: { id: undefined, name: '', email: '' } as Client,
      isEditingClient: false,
      newVehicle: { id: undefined, client_id: 0, license_plate: '', model: '' } as Vehicle,
      isEditingVehicle: false
    };
  },
  created() {
    this.fetchClients();

    // Configurar Pusher para actualizaciones en tiempo real
    const channel = pusher.subscribe('my-channel');
    channel.bind('my-event', (data: any) => {
      console.log('WebSocket message received:', data);
      this.fetchClients();
    });
  },
  methods: {
    async fetchClients() {
      try {
        const response = await apiService.getClients();
        this.clients = response.data;
        this.view = 'clients';
      } catch (error) {
        console.error('Error fetching clients:', error);
      }
    },
    async fetchClientsAndVehicles() {
      await this.fetchClients();
      for (const client of this.clients) {
        await this.viewClientVehicles(client);
      }
      this.view = 'clientsAndVehicles';
    },
    async viewClientVehicles(client: Client) {
      try {
        const response = await apiService.getClientVehicles(client.id!);
        client.vehicles = response.data;
      } catch (error) {
        console.error('Error fetching client vehicles:', error);
      }
    },
    async fetchPayments() {
      try {
        const response = await apiService.getPayments();
        this.payments = response.data;
        this.view = 'payments';
      } catch (error) {
        console.error('Error fetching payments:', error);
      }
    },
    async createOrUpdateClient() {
      try {
        if (this.isEditingClient) {
          await apiService.updateClient(this.newClient);
        } else {
          await apiService.addClient(this.newClient);
        }
        this.resetClientForm();
        this.fetchClients();
      } catch (error) {
        console.error('Error creating/updating client:', error);
      }
    },
    async createOrUpdateVehicle() {
      try {
        if (this.isEditingVehicle) {
          await apiService.updateVehicle(this.newVehicle);
        } else {
          await apiService.addVehicle(this.newVehicle);
        }
        this.resetVehicleForm();
        this.fetchClients(); // Update client list to show new vehicle
      } catch (error) {
        console.error('Error creating/updating vehicle:', error);
      }
    },
    editClient(client: Client) {
      this.newClient = { ...client };
      this.isEditingClient = true;
      this.view = 'postClient';
    },
    cancelEditClient() {
      this.resetClientForm();
    },
    resetClientForm() {
      this.newClient = { id: undefined, name: '', email: '' };
      this.isEditingClient = false;
    },
    editVehicle(vehicle: Vehicle) {
      this.newVehicle = { ...vehicle };
      this.isEditingVehicle = true;
      this.view = 'postVehicle';
    },
    cancelEditVehicle() {
      this.resetVehicleForm();
    },
    resetVehicleForm() {
      this.newVehicle = { id: undefined, client_id: 0, license_plate: '', model: '' };
      this.isEditingVehicle = false;
    },
    async generatePdf() {
      try {
        const response = await apiService.getPdfReport();
        const url = window.URL.createObjectURL(new Blob([response.data], { type: 'application/pdf' }));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'Reporte_Clientes.pdf');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error fetching PDF report:', error);
      }
    },
    showPostClientForm() {
      this.view = 'postClient';
    },
    showPostVehicleForm() {
      this.view = 'postVehicle';
    }
  }
});
</script>

<style scoped>
.dashboard {
  font-family: Arial, sans-serif;
  padding: 20px;
  max-width: 1200px;
  margin: auto;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
h1 {
  font-size: 2.5em;
  text-align: center;
  margin-bottom: 1em;
  color: #333;
}
.buttons-section {
  display: flex;
  justify-content: center;
  gap: 10px;
  margin-bottom: 20px;
}
button {
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 0.5em 1em;
  font-size: 1em;
  cursor: pointer;
  transition: background-color 0.3s ease;
}
button:hover {
  background-color: #0056b3;
}
.content-section {
  margin-top: 20px;
}
h2 {
  font-size: 1.8em;
  margin-bottom: 0.5em;
  color: #333;
}
table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}
th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}
th {
  background-color: #f1f1f1;
}
tbody tr:hover {
  background-color: #f9f9f9;
}
.form-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
input {
  padding: 0.5em;
  margin-bottom: 1em;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1em;
}
.buttons {
  display: flex;
  gap: 10px;
}
</style>
