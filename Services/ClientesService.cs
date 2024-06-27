using System.Collections.Generic;
using System.Data.Entity;
using System.Linq;
using TA2PARQUEADERO.data;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.services
{
    public class ClientesService
    {
        private readonly ApplicationDbContext _context;

        public ClientesService(ApplicationDbContext context)
        {
            _context = context;
        }

        public IEnumerable<Clientes> GetClientes()
        {
            return _context.Clientes.ToList();
        }

        public Clientes GetCliente(int id)
        {
            return _context.Clientes.Find(id);
        }

        public void CreateCliente(Clientes cliente)
        {
            _context.Clientes.Add(cliente);
            _context.SaveChanges();
        }

        public void UpdateCliente(Clientes cliente)
        {
            _context.Entry(cliente).State = EntityState.Modified;
            _context.SaveChanges();
        }

        public void DeleteCliente(int id)
        {
            var cliente = _context.Clientes.Find(id);
            if (cliente != null)
            {
                _context.Clientes.Remove(cliente);
                _context.SaveChanges();
            }
        }

        public bool ClienteExists(int id)
        {
            return _context.Clientes.Any(e => e.Id == id);
        }
    }
}

