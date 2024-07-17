// src/main.ts

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import apiClient from './axios';  // Importar el cliente Axios
import pusher from './pusher';

const app = createApp(App);

app.use(router);

app.config.globalProperties.$axios = apiClient;  // Añadir Axios a las propiedades globales de la aplicación
app.config.globalProperties.$pusher = pusher;

app.mount('#app');
