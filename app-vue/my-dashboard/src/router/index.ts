// src/router/index.ts

import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../views/Dashboard.vue';
import Clients from '../views/Clients.vue';
import Vehicles from '../views/Vehicles.vue';
import Payments from '../views/Payments.vue';
import Reports from '../views/Reports.vue';

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/clients',
    name: 'Clients',
    component: Clients
  },
  {
    path: '/vehicles',
    name: 'Vehicles',
    component: Vehicles
  },
  {
    path: '/payments',
    name: 'Payments',
    component: Payments
  },
  {
    path: '/reports',
    name: 'Reports',
    component: Reports
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
