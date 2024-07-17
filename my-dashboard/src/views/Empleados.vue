<template>
  <div>
    <h1>Empleados</h1>
    <button @click="showPostEmpleadoForm">Agregar Empleado</button>
    <table>
      <thead>
        <tr>
          <th>Nombre</th>
          <th>Apellido</th>
          <th>Email</th>
          <th>Salario</th>
          <th>Fecha Ingreso</th>
          <th>Estado</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="empleado in empleados" :key="empleado.id">
          <td>{{ empleado.nombre }}</td>
          <td>{{ empleado.apellido }}</td>
          <td>{{ empleado.email }}</td>
          <td>{{ empleado.salario }}</td>
          <td>{{ new Date(empleado.fechaIngreso).toLocaleDateString() }}</td>
          <td>{{ empleado.estado }}</td>
          <td>
            <button @click="editEmpleado(empleado)">Editar</button>
            <button @click="deleteEmpleado(empleado.id)">Eliminar</button>
            <button @click="recoverEmpleado(empleado.id)">Recuperar</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="view === 'postEmpleado'">
      <h2>{{ isEditingEmpleado ? 'Editar Empleado' : 'Agregar Empleado' }}</h2>
      <form @submit.prevent="createOrUpdateEmpleado">
        <input v-model="newEmpleado.nombre" placeholder="Nombre del Empleado" required />
        <input v-model="newEmpleado.apellido" placeholder="Apellido del Empleado" required />
        <input v-model="newEmpleado.email" placeholder="Correo del Empleado" required />
        <input v-model="newEmpleado.salario" placeholder="Salario del Empleado" required type="number" />
        <input v-model="newEmpleado.fechaIngreso" placeholder="Fecha de Ingreso" required type="date" />
        <div class="buttons">
          <button type="submit">{{ isEditingEmpleado ? 'Actualizar' : 'Agregar' }} Empleado</button>
          <button type="button" @click="cancelEditEmpleado">Cancelar</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import apiService from '@/services/apiService';
import { Empleado } from '@/types';

export default defineComponent({
  data() {
    return {
      empleados: [] as Empleado[],
      newEmpleado: { id: undefined, nombre: '', apellido: '', email: '', salario: 0, fechaIngreso: new Date(), estado: 'ACTIVO' } as Empleado,
      isEditingEmpleado: false,
      view: ''
    };
  },
  methods: {
    async fetchEmpleados() {
      try {
        const response = await apiService.getEmpleados();
        this.empleados = response.data;
      } catch (error) {
        console.error('Error fetching empleados:', error);
      }
    },
    async createOrUpdateEmpleado() {
      try {
        if (this.isEditingEmpleado) {
          await apiService.updateEmpleado(this.newEmpleado);
        } else {
          await apiService.addEmpleado(this.newEmpleado);
        }
        this.resetEmpleadoForm();
        this.fetchEmpleados();
      } catch (error) {
        console.error('Error creating/updating empleado:', error);
      }
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
      this.view = '';
    },
    showPostEmpleadoForm() {
      this.view = 'postEmpleado';
    }
  },
  created() {
    this.fetchEmpleados();
  }
});
</script>

<style scoped>
/* Estilos CSS */
</style>
