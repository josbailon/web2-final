-- Creación de la tabla Clientes
CREATE TABLE Clientes (
    Id SERIAL PRIMARY KEY,
    Nombre VARCHAR(255),
    Apellido VARCHAR(255),
    Telefono VARCHAR(20),
    Direccion VARCHAR(255),
    Email VARCHAR(255)
);

-- Creación de la tabla Carros
CREATE TABLE Carros (
    Id SERIAL PRIMARY KEY,
    Marca VARCHAR(255),
    Modelo VARCHAR(255),
    Color VARCHAR(50),
    Placa VARCHAR(20) UNIQUE,
    PropietarioId INTEGER,
    FOREIGN KEY (PropietarioId) REFERENCES Clientes(Id)
);

-- Creación de la tabla Ingresos
CREATE TABLE Ingresos (
    Id SERIAL PRIMARY KEY,
    FechaEntrada TIMESTAMP,
    VehiculoId INTEGER,
    FOREIGN KEY (VehiculoId) REFERENCES Carros(Id)
);

-- Creación de la tabla Salidas
CREATE TABLE Salidas (
    Id SERIAL PRIMARY KEY,
    FechaSalida TIMESTAMP,
    IngresoId INTEGER,
    FOREIGN KEY (IngresoId) REFERENCES Ingresos(Id)
);

-- Creación de la tabla Pagos
CREATE TABLE Pagos (
    Id SERIAL PRIMARY KEY,
    Monto DECIMAL(10, 2),
    MetodoPago VARCHAR(100),
    FechaPago DATE,
    IngresoId INTEGER,
    FOREIGN KEY (IngresoId) REFERENCES Ingresos(Id)
);

-- Creación de la tabla TarifasEstacionamiento
CREATE TABLE TarifasEstacionamiento (
    Id SERIAL PRIMARY KEY,
    Tarifa DECIMAL(10, 2) NOT NULL,
    Descripcion VARCHAR(255)
);
