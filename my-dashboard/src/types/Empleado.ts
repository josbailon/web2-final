// src/types/Empleado.ts
export interface Empleado {
    id?: number;
    nombre: string;
    apellido: string;
    fechaIngreso: Date;
    email: string;
    salario: number;
    estado: 'ACTIVO' | 'INACTIVO';
  }
  