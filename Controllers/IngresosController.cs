using System.Collections.Generic;
using System.Net;
using System.Web.Http;
using System.Web.Http.Description;
using TA2PARQUEADERO.models;
using TA2PARQUEADERO.services;

namespace TA2PARQUEADERO.Controllers
{
    [RoutePrefix("api/Ingresos")]
    public class IngresosController : ApiController
    {
        private readonly IngresosService _ingresosService;

        public IngresosController(IngresosService ingresosService)
        {
            _ingresosService = ingresosService;
        }
        //GET http://localhost:5000/api/ingresos
        [HttpGet]
        [Route("")]
        public IEnumerable<Ingresos> GetIngresos()
        {
            return _ingresosService.GetIngresos();
        }
       
        //GET http://localhost:5000/api/ingresos/{id}
        [HttpGet]
        [Route("{id:int}")]
        [ResponseType(typeof(Ingresos))]
        public IHttpActionResult GetIngreso(int id)
        {
            var ingreso = _ingresosService.GetIngreso(id);
            if (ingreso == null)
            {
                return NotFound();
            }
            return Ok(ingreso);
        }
        
        //POST http://localhost:5000/api/ingresos

        [HttpPost]
        [Route("")]
        [ResponseType(typeof(Ingresos))]
        public IHttpActionResult PostIngreso(Ingresos ingreso)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            _ingresosService.CreateIngreso(ingreso);
            return CreatedAtRoute("DefaultApi", new { id = ingreso.Id }, ingreso);
        }
        //PUT http://localhost:5000/api/ingresos/{id}

        [HttpPut]
        [Route("{id:int}")]
        [ResponseType(typeof(void))]
        public IHttpActionResult PutIngreso(int id, Ingresos ingreso)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            if (id != ingreso.Id)
            {
                return BadRequest();
            }

            _ingresosService.UpdateIngreso(ingreso);
            return StatusCode(HttpStatusCode.NoContent);
        }
        //DELETE http://localhost:5000/api/ingresos/{id}
        [HttpDelete]
        [Route("{id:int}")]
        [ResponseType(typeof(Ingresos))]
        public IHttpActionResult DeleteIngreso(int id)
        {
            var ingreso = _ingresosService.GetIngreso(id);
            if (ingreso == null)
            {
                return NotFound();
            }

            _ingresosService.DeleteIngreso(id);
            return Ok(ingreso);
        }
    }
}

