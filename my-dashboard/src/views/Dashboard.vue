<template>
  <div class="dashboard">
    <h1>Dashboard del Parqueadero</h1>
    <div class="buttons-section">
      <button @click="generatePdf('empleados')">Generar PDF Empleados</button>
      <button @click="generatePdf('reservas')">Generar PDF Reservas</button>
      <button @click="generateExcel('empleados')">Generar Excel Empleados</button>
      <button @click="generateExcel('reservas')">Generar Excel Reservas</button>
      <button @click="fetchEmpleados">Ver Empleados</button>
      <button @click="showPostEmpleadoForm">Post Empleado</button>
      <button @click="fetchReservas">Ver Reservas</button>
      <button @click="showPostReservaForm">Post Reserva</button>
    </div>
    <!-- ... Resto del template ... -->
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import apiService from '@/services/apiService';
import pusher from '@/pusher';
import { Client, Vehicle, Payment, Empleado, Reserva } from '@/types';

export default defineComponent({
  data() {
    return {
      clients: [] as Client[],
      payments: [] as Payment[],
      empleados: [] as Empleado[],
      reservas: [] as Reserva[],
      view: '',
      newClient: { id: undefined, name: '', email: '' } as Client,
      isEditingClient: false,
      newVehicle: { id: undefined, client_id: 0, license_plate: '', model: '' } as Vehicle,
      isEditingVehicle: false,
      newEmpleado: { id: undefined, nombre: '', apellido: '', email: '', salario: 0, fechaIngreso: new Date(), estado: 'ACTIVO' } as Empleado,
      isEditingEmpleado: false,
      newReserva: { id: undefined, idCliente: 0, idVehiculo: 0, fechaReserva: new Date(), horaEntrada: '', horaSalida: '', estadoReserva: 'PENDIENTE' } as Reserva,
      isEditingReserva: false
    };
  },
  created() {
    this.fetchEmpleados();

    // Configurar Pusher para actualizaciones en tiempo real
    const channel = pusher.subscribe('my-channel');
    channel.bind('my-event', (data: any) => {
      console.log('WebSocket message received:', data);
      this.fetchEmpleados();
    });
  },
  methods: {
    async fetchEmpleados() {
      try {
        const response = await apiService.getEmpleados();
        this.empleados = response.data;
        this.view = 'empleados';
      } catch (error) {
        console.error('Error fetching empleados:', error);
      }
    },
    async fetchReservas() {
      try {
        const response = await apiService.getReservas();
        this.reservas = response.data;
        this.view = 'reservas';
      } catch (error) {
        console.error('Error fetching reservas:', error);
      }
    },
    async generatePdf(type: string) {
      try {
        let response;
        if (type === 'empleados') {
          response = await apiService.getEmpleadoPdfReport();
        } else if (type === 'reservas') {
          response = await apiService.getReservaPdfReport();
        }
        if (response) {
          const url = window.URL.createObjectURL(new Blob([response.data], { type: 'application/pdf' }));
          const link = document.createElement('a');
          link.href = url;
          link.setAttribute('download', `${type}_report.pdf`);
          document.body.appendChild(link);
          link.click();
        }
      } catch (error) {
        console.error(`Error generating ${type} PDF report:`, error);
      }
    },
    async generateExcel(type: string) {
      try {
        let response;
        if (type === 'empleados') {
          response = await apiService.getEmpleadoExcelReport();
        } else if (type === 'reservas') {
          response = await apiService.getReservaExcelReport();
        }
        if (response) {
          const url = window.URL.createObjectURL(new Blob([response.data], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' }));
          const link = document.createElement('a');
          link.href = url;
          link.setAttribute('download', `${type}_report.xlsx`);
          document.body.appendChild(link);
          link.click();
        }
      } catch (error) {
        console.error(`Error generating ${type} Excel report:`, error);
      }
    },
    editEmpleado(empleado: Empleado) {
      this.newEmpleado = { ...empleado, fechaIngreso: new Date(empleado.fechaIngreso) };
      this.isEditingEmpleado = true;
      this.view = 'postEmpleado';
    },
    cancelEditEmpleado() {
      this.resetEmpleadoForm();
    },
    resetEmpleadoForm() {
      this.newEmpleado = { id: undefined, nombre: '', apellido: '', email: '', salario: 0, fechaIngreso: new Date(), estado: 'ACTIVO' };
      this.isEditingEmpleado = false;
    },
    editReserva(reserva: Reserva) {
      this.newReserva = { ...reserva, fechaReserva: new Date(reserva.fechaReserva) };
      this.isEditingReserva = true;
      this.view = 'postReserva';
    },
    cancelEditReserva() {
      this.resetReservaForm();
    },
    resetReservaForm() {
      this.newReserva = { id: undefined, idCliente: 0, idVehiculo: 0, fechaReserva: new Date(), horaEntrada: '', horaSalida: '', estadoReserva: 'PENDIENTE' };
      this.isEditingReserva = false;
    },
    async deleteEmpleado(id: number) {
      try {
        await apiService.deleteEmpleado(id);
        this.fetchEmpleados();
      } catch (error) {
        console.error('Error deleting empleado:', error);
      }
    },
    async recoverEmpleado(id: number) {
      try {
        await apiService.recoverEmpleado(id);
        this.fetchEmpleados();
      } catch (error) {
        console.error('Error recovering empleado:', error);
      }
    },
    async deleteReserva(id: number) {
      try {
        await apiService.deleteReserva(id);
        this.fetchReservas();
      } catch (error) {
        console.error('Error deleting reserva:', error);
      }
    },
    async recoverReserva(id: number) {
      try {
        await apiService.recoverReserva(id);
        this.fetchReservas();
      } catch (error) {
        console.error('Error recovering reserva:', error);
      }
    },
    showPostEmpleadoForm() {
      this.view = 'postEmpleado';
    },
    showPostReservaForm() {
      this.view = 'postReserva';
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
