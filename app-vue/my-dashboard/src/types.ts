export interface Client {
  id?: number;
  name: string;
  email: string;
  vehicles?: Vehicle[];
}

export interface Vehicle {
  id?: number;
  client_id: number;
  license_plate: string;
  model: string;
}

export interface Payment {
  id: number;
  amount: number;
}
