using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.models
{
    public class Ingresos
    {
        public int Id { get; set; }
        [Required]
        public DateTime FechaEntrada { get; set; }
        [Required]
        public int VehiculoId { get; set; }

        // Relación con Carros
        public Carros Vehiculo { get; set; }

        // Relación con Salidas
        public Salidas Salida { get; set; }
    }
}


