<!-- src/views/Reports.vue -->

<template>
  <div>
    <h2>Reports</h2>
    <button @click="generateEmpleadoPDF">Generate Empleado PDF Report</button>
    <button @click="generateReservaPDF">Generate Reserva PDF Report</button>
    <button @click="generateEmpleadoExcel">Generate Empleado Excel Report</button>
    <button @click="generateReservaExcel">Generate Reserva Excel Report</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import apiService from '@/services/apiService';

export default defineComponent({
  name: 'Reports',
  methods: {
    async generateEmpleadoPDF() {
      try {
        const response = await apiService.getEmpleadoPdfReport();
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'empleados_report.pdf');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error generating Empleado PDF report:', error);
      }
    },
    async generateReservaPDF() {
      try {
        const response = await apiService.getReservaPdfReport();
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'reservas_report.pdf');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error generating Reserva PDF report:', error);
      }
    },
    async generateEmpleadoExcel() {
      try {
        const response = await apiService.getEmpleadoExcelReport();
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'empleados_report.xlsx');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error generating Empleado Excel report:', error);
      }
    },
    async generateReservaExcel() {
      try {
        const response = await apiService.getReservaExcelReport();
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'reservas_report.xlsx');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error generating Reserva Excel report:', error);
      }
    }
  }
});
</script>
