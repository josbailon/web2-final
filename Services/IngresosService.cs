using System.Collections.Generic;
using System.Data.Entity;
using System.Linq;
using TA2PARQUEADERO.data;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.services
{
    public class IngresosService
    {
        private readonly ApplicationDbContext _context;

        public IngresosService(ApplicationDbContext context)
        {
            _context = context;
        }

        public IEnumerable<Ingresos> GetIngresos()
        {
            return _context.Ingresos.ToList();
        }

        public Ingresos GetIngreso(int id)
        {
            return _context.Ingresos.Find(id);
        }

        public void CreateIngreso(Ingresos ingreso)
        {
            _context.Ingresos.Add(ingreso);
            _context.SaveChanges();
        }

        public void UpdateIngreso(Ingresos ingreso)
        {
            _context.Entry(ingreso).State = EntityState.Modified;
            _context.SaveChanges();
        }

        public void DeleteIngreso(int id)
        {
            var ingreso = _context.Ingresos.Find(id);
            if (ingreso != null)
            {
                _context.Ingresos.Remove(ingreso);
                _context.SaveChanges();
            }
        }

        public bool IngresoExists(int id)
        {
            return _context.Ingresos.Any(e => e.Id == id);
        }
    }
}

