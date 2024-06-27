using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TA2PARQUEADERO.models
{
    public class Pagos
    {
        [Required]
        public int Id { get; set; }

        [Range(0, double.MaxValue)]
        public decimal Monto { get; set; }
        public string MetodoPago { get; set; }
        [Required]
        public DateTime FechaPago { get; set; }
        [Required]
        public int IngresoId { get; set; }

        // Relación con Ingresos
        public Ingresos Ingreso { get; set; }
    }
}

