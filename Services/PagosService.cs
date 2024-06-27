using System.Collections.Generic;
using System.Data.Entity;
using System.Linq;
using TA2PARQUEADERO.data;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.services
{
    public class PagosService
    {
        private readonly ApplicationDbContext _context;

        public PagosService(ApplicationDbContext context)
        {
            _context = context;
        }

        public IEnumerable<Pagos> GetPagos()
        {
            return _context.Pagos.ToList();
        }

        public Pagos GetPago(int id)
        {
            return _context.Pagos.Find(id);
        }

        public void CreatePago(Pagos pago)
        {
            _context.Pagos.Add(pago);
            _context.SaveChanges();
        }

        public void UpdatePago(Pagos pago)
        {
            _context.Entry(pago).State = EntityState.Modified;
            _context.SaveChanges();
        }

        public void DeletePago(int id)
        {
            var pago = _context.Pagos.Find(id);
            if (pago != null)
            {
                _context.Pagos.Remove(pago);
                _context.SaveChanges();
            }
        }

        public bool PagoExists(int id)
        {
            return _context.Pagos.Any(e => e.Id == id);
        }
    }
}

