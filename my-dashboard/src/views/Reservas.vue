<template>
  <div>
    <h1>Reservas</h1>
    <button @click="showPostReservaForm">Agregar Reserva</button>
    <table>
      <thead>
        <tr>
          <th>Cliente</th>
          <th>Vehículo</th>
          <th>Fecha Reserva</th>
          <th>Hora Entrada</th>
          <th>Hora Salida</th>
          <th>Estado</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="reserva in reservas" :key="reserva.id">
          <td>{{ reserva.idCliente }}</td>
          <td>{{ reserva.idVehiculo }}</td>
          <td>{{ new Date(reserva.fechaReserva).toLocaleDateString() }}</td>
          <td>{{ reserva.horaEntrada }}</td>
          <td>{{ reserva.horaSalida }}</td>
          <td>{{ reserva.estadoReserva }}</td>
          <td>
            <button @click="editReserva(reserva)">Editar</button>
            <button @click="deleteReserva(reserva.id)">Eliminar</button>
            <button @click="recoverReserva(reserva.id)">Recuperar</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="view === 'postReserva'">
      <h2>{{ isEditingReserva ? 'Editar Reserva' : 'Agregar Reserva' }}</h2>
      <form @submit.prevent="createOrUpdateReserva">
        <input v-model="newReserva.idCliente" placeholder="ID del Cliente" required type="number" />
        <input v-model="newReserva.idVehiculo" placeholder="ID del Vehículo" required type="number" />
        <input v-model="newReserva.fechaReserva" placeholder="Fecha de Reserva" required type="date" />
        <input v-model="newReserva.horaEntrada" placeholder="Hora de Entrada" required type="time" />
        <input v-model="newReserva.horaSalida" placeholder="Hora de Salida" required type="time" />
        <select v-model="newReserva.estadoReserva" required>
          <option value="PENDIENTE">Pendiente</option>
          <option value="CONFIRMADA">Confirmada</option>
          <option value="CANCELADA">Cancelada</option>
        </select>
        <div class="buttons">
          <button type="submit">{{ isEditingReserva ? 'Actualizar' : 'Agregar' }} Reserva</button>
          <button type="button" @click="cancelEditReserva">Cancelar</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import apiService from '@/services/apiService';
import { Reserva } from '@/types';

export default defineComponent({
  data() {
    return {
      reservas: [] as Reserva[],
      newReserva: { id: undefined, idCliente: 0, idVehiculo: 0, fechaReserva: new Date(), horaEntrada: '', horaSalida: '', estadoReserva: 'PENDIENTE' } as Reserva,
      isEditingReserva: false,
      view: ''
    };
  },
  methods: {
    async fetchReservas() {
      try {
        const response = await apiService.getReservas();
        this.reservas = response.data;
      } catch (error) {
        console.error('Error fetching reservas:', error);
      }
    },
    async createOrUpdateReserva() {
      try {
        if (this.isEditingReserva) {
          await apiService.updateReserva(this.newReserva);
        } else {
          await apiService.addReserva(this.newReserva);
        }
        this.resetReservaForm();
        this.fetchReservas();
      } catch (error) {
        console.error('Error creating/updating reserva:', error);
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
      this.view = '';
    },
    showPostReservaForm() {
      this.view = 'postReserva';
    }
  },
  created() {
    this.fetchReservas();
  }
});
</script>

<style scoped>
/* Estilos CSS */
</style>
