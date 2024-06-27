using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TA2PARQUEADERO.models
{
    using System;
    using System.ComponentModel.DataAnnotations;

   
        public class Carros
        {
            public int Id { get; set; }
            [Required]
            public string Marca { get; set; }
            [Required]
           public string Modelo { get; set; }
             [Required]
          public string Color { get; set; }
             [Required]
          [StringLength(20)]
        public string Placa { get; set; }
        [Required]
        public Clientes Propietario { get; set; }

        public int PropietarioId { get; set; }

        }
    }

