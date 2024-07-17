// src/types/Vehicle.ts

import { Client } from './Client';

export interface Vehicle {
  id: number;
  license_plate: string;
  model: string;
  client: Client;
}
