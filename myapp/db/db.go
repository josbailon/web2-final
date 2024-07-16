package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // Importar el driver de PostgreSQL
)

// Configuración de la base de datos
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "1234"
    dbname   = "DB_MyApp"
)

// Db es una variable global que contendrá la conexión a la base de datos
var Db *sql.DB

// Init inicializa la conexión a la base de datos
func Init() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Establecer la conexión a la base de datos
    Db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }

    // Verificar la conexión
    err = Db.Ping()
    if err != nil {
        log.Fatalf("Error al realizar ping a la base de datos: %v", err)
    }

    fmt.Println("Conexión exitosa a la base de datos")
}

// Conectar devuelve una referencia a la conexión a la base de datos
func Conectar() *sql.DB {
    return Db
}
