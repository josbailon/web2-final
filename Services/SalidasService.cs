using System.Collections.Generic;
using System.Data.Entity;
using System.Linq;
using TA2PARQUEADERO.data;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.services
{
    public class SalidasService
    {
        private readonly ApplicationDbContext _context;

        public SalidasService(ApplicationDbContext context)
        {
            _context = context;
        }

        public IEnumerable<Salidas> GetSalidas()
        {
            return _context.Salidas.ToList();
        }

        public Salidas GetSalida(int id)
        {
            return _context.Salidas.Find(id);
        }

        public void CreateSalida(Salidas salida)
        {
            _context.Salidas.Add(salida);
            _context.SaveChanges();
        }

        public void UpdateSalida(Salidas salida)
        {
            _context.Entry(salida).State = EntityState.Modified;
            _context.SaveChanges();
        }

        public void DeleteSalida(int id)
        {
            var salida = _context.Salidas.Find(id);
            if (salida != null)
            {
                _context.Salidas.Remove(salida);
                _context.SaveChanges();
            }
        }

        public bool SalidaExists(int id)
        {
            return _context.Salidas.Any(e => e.Id == id);
        }
    }
}

