// src/apollo.ts

import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client/core';
import { provideApolloClient, DefaultApolloClient } from '@vue/apollo-composable';
import { App } from 'vue';

const httpLink = createHttpLink({
  uri: 'http://localhost:8000/graphql',  // Asegúrate de que esta URL sea correcta
  headers: {
    Authorization: `Bearer 2|Xyw4zcVbSjU453Ap6KhqRamAVhOhMj6r3jllbj8tc1cecff2`,
  },
});

const cache = new InMemoryCache();

const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});

export function provideApollo(app: App) {
  app.provide(DefaultApolloClient, apolloClient);
}
