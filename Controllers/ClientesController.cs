using System.Collections.Generic;
using System.Net;
using System.Web.Http;
using System.Web.Http.Description;
using TA2PARQUEADERO.models;
using TA2PARQUEADERO.services;

namespace TA2PARQUEADERO.Controllers
{
    [RoutePrefix("api/Clientes")]
    
    
    public class ClientesController : ApiController
    {
        private readonly ClientesService _clientesService;

        public ClientesController(ClientesService clientesService)
        {
            _clientesService = clientesService;
        }
        //GET http://localhost:5000/api/clientes
        //Obtener todos los clientes
        [HttpGet]
        [Route("")]
        public IEnumerable<Clientes> GetClientes()
        {
            return _clientesService.GetClientes();
        }
        //GET http://localhost:5000/api/clientes/{id}
        // Obtener un cliente por ID

        [HttpGet]
        [Route("{id:int}")]
        [ResponseType(typeof(Clientes))]
        public IHttpActionResult GetCliente(int id)
        {
            var cliente = _clientesService.GetCliente(id);
            if (cliente == null)
            {
                return NotFound();
            }
            return Ok(cliente);
        }
        //POST http://localhost:5000/api/clientes
        //Crear un nuevo cliente
        [HttpPost]
        [Route("")]
        [ResponseType(typeof(Clientes))]
        public IHttpActionResult PostCliente(Clientes cliente)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            _clientesService.CreateCliente(cliente);
            return CreatedAtRoute("DefaultApi", new { id = cliente.Id }, cliente);
        }
        //PUT http://localhost:5000/api/clientes/{id}
        // Actualizar un cliente por ID

        [HttpPut]
        [Route("{id:int}")]
        [ResponseType(typeof(void))]
        public IHttpActionResult PutCliente(int id, Clientes cliente)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            if (id != cliente.Id)
            {
                return BadRequest();
            }

            _clientesService.UpdateCliente(cliente);
            return StatusCode(HttpStatusCode.NoContent);
        }
        //DELETE http://localhost:5000/api/clientes/{id}
        //Borrar un cliente por ID

        [HttpDelete]
        [Route("{id:int}")]
        [ResponseType(typeof(Clientes))]
        public IHttpActionResult DeleteCliente(int id)
        {
            var cliente = _clientesService.GetCliente(id);
            if (cliente == null)
            {
                return NotFound();
            }

            _clientesService.DeleteCliente(id);
            return Ok(cliente);
        }
    }
}

