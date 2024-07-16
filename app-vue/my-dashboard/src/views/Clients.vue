// src/views/Clients.vue

<template>
  <div>
    <h2>Clients</h2>
    <ul>
      <li v-for="client in clients" :key="client.id">{{ client.name }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useQuery } from '@vue/apollo-composable';
import gql from 'graphql-tag';
import pusher from '../pusher';
import { Client } from '../types/Client';

const GET_CLIENTS = gql`
  query {
    clients {
      id
      name
      email
    }
  }
`;

export default defineComponent({
  name: 'Clients',
  setup() {
    const clients = ref<Client[]>([]);
    const { result } = useQuery(GET_CLIENTS);

    onMounted(() => {
      result.value?.clients && (clients.value = result.value.clients);

      pusher.subscribe('clients').bind('App\\Events\\ClientAdded', (data: { client: Client }) => {
        clients.value.push(data.client);
      });
    });

    return {
      clients,
    };
  },
});
</script>
