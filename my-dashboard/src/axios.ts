// src/axios.ts
import axios from 'axios';

// src/services/apiService.ts
const apiClient = axios.create({
  baseURL: 'https://2be5-191-99-252-133.ngrok-free.app/api', // Reemplaza <tu-url-de-ngrok> con la URL proporcionada por ngrok
  headers: {
    'Content-Type': 'application/json',
  }
});


export default apiClient;
