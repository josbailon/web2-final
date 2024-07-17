// src/views/Payments.vue

<template>
  <div>
    <h2>Payments</h2>
    <ul>
      <li v-for="payment in payments" :key="payment.id">{{ payment.amount }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useQuery } from '@vue/apollo-composable';
import gql from 'graphql-tag';
import { Payment } from '../types/Payment';

const GET_PAYMENTS = gql`
  query {
    payments {
      id
      amount
      client {
        id
        name
      }
    }
  }
`;

export default defineComponent({
  name: 'Payments',
  setup() {
    const payments = ref<Payment[]>([]);
    const { result } = useQuery(GET_PAYMENTS);

    result.value?.payments && (payments.value = result.value.payments);

    return {
      payments,
    };
  },
});
</script>
