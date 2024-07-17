// src/types/Payment.ts

import { Client } from './Client';

export interface Payment {
  id: number;
  amount: number;
  client: Client;
}
