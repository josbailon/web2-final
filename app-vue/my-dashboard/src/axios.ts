// src/axios.ts

import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://127.0.0.1:8000/api', // Base URL de tu API
  headers: {
    'Content-Type': 'application/json',
    Authorization: `Bearer 2|Xyw4zcVbSjU453Ap6KhqRamAVhOhMj6r3jllbj8tc1cecff2`
  }
});

export default apiClient;
