using System.Collections.Generic;
using System.Data.Entity;
using System.Linq;
using TA2PARQUEADERO.data;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.services
{
    public class CarrosService
    {
        private readonly ApplicationDbContext _context;

        public CarrosService(ApplicationDbContext context)
        {
            _context = context;
        }

        public IEnumerable<Carros> GetCarros()
        {
            return _context.Carros.ToList();
        }

        public Carros GetCarro(int id)
        {
            return _context.Carros.Find(id);
        }

        public void CreateCarro(Carros carro)
        {
            _context.Carros.Add(carro);
            _context.SaveChanges();
        }

        public void UpdateCarro(Carros carro)
        {
            _context.Entry(carro).State = EntityState.Modified;
            _context.SaveChanges();
        }

        public void DeleteCarro(int id)
        {
            var carro = _context.Carros.Find(id);
            if (carro != null)
            {
                _context.Carros.Remove(carro);
                _context.SaveChanges();
            }
        }

        public bool CarroExists(int id)
        {
            return _context.Carros.Any(e => e.Id == id);
        }
    }
}

