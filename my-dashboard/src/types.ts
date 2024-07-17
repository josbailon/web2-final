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
  id?: number;
  amount: number;
}

export interface Empleado {
  id?: number;
  nombre: string;
  apellido: string;
  email: string;
  salario: number;
  fechaIngreso: Date;
  estado: 'ACTIVO' | 'INACTIVO';
}

export interface Reserva {
  id?: number;
  idCliente: number;
  idVehiculo: number;
  fechaReserva: Date;
  horaEntrada: string;
  horaSalida: string;
  estadoReserva: 'PENDIENTE' | 'CONFIRMADA' | 'CANCELADA';
}
