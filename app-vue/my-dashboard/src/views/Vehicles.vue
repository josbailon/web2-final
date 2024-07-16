// src/views/Vehicles.vue

<template>
  <div>
    <h2>Vehicles</h2>
    <ul>
      <li v-for="vehicle in vehicles" :key="vehicle.id">{{ vehicle.license_plate }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useQuery } from '@vue/apollo-composable';
import gql from 'graphql-tag';
import { Vehicle } from '../types/Vehicle';

const GET_VEHICLES = gql`
  query {
    vehicles {
      id
      license_plate
      model
      client {
        id
        name
      }
    }
  }
`;

export default defineComponent({
  name: 'Vehicles',
  setup() {
    const vehicles = ref<Vehicle[]>([]);
    const { result } = useQuery(GET_VEHICLES);

    result.value?.vehicles && (vehicles.value = result.value.vehicles);

    return {
      vehicles,
    };
  },
});
</script>
