-- Insertar datos en la tabla Clientes
INSERT INTO Clientes (Nombre, Apellido, Telefono, Direccion, Email) VALUES
('lola', 'Loor', '1234567891', 'Altos del porvenir', 'loor123@gmail.com'),
('Ricardo', 'Mero', '9876543211', 'Barrio 4 de Noviembre', 'mero@gmail.com'),
('Jose', 'Pinargote', '1122334456', 'Maria Auxiliadora', 'jose.pina@gmail.com');

-- Seleccionar los IDs generados
SELECT * FROM Clientes;


-- Insertar datos en la tabla Carros
INSERT INTO Carros (Marca, Modelo, Color, Placa, PropietarioId) VALUES
('Nissan', 'Sentra', 'Rojo', 'XYZ127', 1),
('Chevrolet', 'Malibu', 'Azul', 'ABC477', 2),
('BMW', 'X5', 'Blanco', 'LMN798', 3);

-- Seleccionar los IDs generados para los carros
SELECT * FROM Carros;


-- Insertar datos en la tabla Ingresos
INSERT INTO Ingresos (FechaEntrada, VehiculoId) VALUES
('2024-08-01 08:00:00', 1),
('2024-08-02 09:30:00', 2),
('2024-08-03 10:45:00', 3);

-- Seleccionar los IDs generados para los ingresos
SELECT * FROM Ingresos;


-- Insertar datos en la tabla Salidas
INSERT INTO Salidas (FechaSalida, IngresoId) VALUES
('2024-08-01 18:00:00', 1),
('2024-08-02 19:30:00', 2),
('2024-08-03 20:45:00', 3);

SELECT * FROM Salidas;

-- Insertar datos en la tabla Pagos
INSERT INTO Pagos (Monto, MetodoPago, FechaPago, IngresoId) VALUES
(18.00, 'Tarjeta de Crédito', '2024-08-01', 1),
(22.00, 'Efectivo', '2024-08-02', 2),
(27.00, 'Tarjeta de Débito', '2024-08-03', 3);

SELECT * FROM Pagos;

-- Insertar datos en la tabla TarifasEstacionamiento
INSERT INTO TarifasEstacionamiento (Tarifa, Descripcion) VALUES
(12.00, 'Tarifa por hora para autos pequeños'),
(17.00, 'Tarifa por hora para autos medianos'),
(22.00, 'Tarifa por hora para autos grandes');

SELECT * FROM TarifasEstacionamiento;













