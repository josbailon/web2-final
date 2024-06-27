using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TA2PARQUEADERO.models
{
    public class Salidas
    {
        public int Id { get; set; }
        [Required]
        public DateTime FechaSalida { get; set; }
        [Required]
        public int IngresoId { get; set; }

        // Relación con Ingresos
        public Ingresos Ingreso { get; set; }
    }
}

