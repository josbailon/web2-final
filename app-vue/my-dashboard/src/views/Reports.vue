// src/views/Reports.vue

<template>
  <div>
    <h2>Reports</h2>
    <button @click="generatePDF">Generate PDF Report</button>
    <button @click="generateExcel">Generate Excel Report</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

export default defineComponent({
  name: 'Reports',
  methods: {
    async generatePDF() {
      try {
        const response = await axios.get('http://localhost:8000/api/reports/clients/pdf', {
          responseType: 'blob',
          headers: {
            Authorization: `Bearer 2|Xyw4zcVbSjU453Ap6KhqRamAVhOhMj6r3jllbj8tc1cecff2`,
          },
        });
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'clients_report.pdf');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error generating PDF report:', error);
      }
    },
    async generateExcel() {
      try {
        const response = await axios.get('http://localhost:8000/api/reports/clients/excel', {
          responseType: 'blob',
          headers: {
            Authorization: `Bearer 2|Xyw4zcVbSjU453Ap6KhqRamAVhOhMj6r3jllbj8tc1cecff2`,
          },
        });
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'clients_report.xlsx');
        document.body.appendChild(link);
        link.click();
      } catch (error) {
        console.error('Error generating Excel report:', error);
      }
    },
  },
});
</script>
