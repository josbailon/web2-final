using System.Collections.Generic;
using System.Net;
using System.Web.Http;
using System.Web.Http.Description;
using TA2PARQUEADERO.models;
using TA2PARQUEADERO.services;

namespace TA2PARQUEADERO.Controllers
{
    [RoutePrefix("api/Salidas")]
    public class SalidasController : ApiController
    {
        private readonly SalidasService _salidasService;

        public SalidasController(SalidasService salidasService)
        {
            _salidasService = salidasService;
        }
        //GET http://localhost:5000/api/salidas

        [HttpGet]
        [Route("")]
        public IEnumerable<Salidas> GetSalidas()
        {
            return _salidasService.GetSalidas();
        }
        //GET http://localhost:5000/api/salidas/{id}

        [HttpGet]
        [Route("{id:int}")]
        [ResponseType(typeof(Salidas))]
        public IHttpActionResult GetSalida(int id)
        {
            var salida = _salidasService.GetSalida(id);
            if (salida == null)
            {
                return NotFound();
            }
            return Ok(salida);
        }
        //POST http://localhost:5000/api/salidas

        [HttpPost]
        [Route("")]
        [ResponseType(typeof(Salidas))]
        public IHttpActionResult PostSalida(Salidas salida)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            _salidasService.CreateSalida(salida);
            return CreatedAtRoute("DefaultApi", new { id = salida.Id }, salida);
        }
        //PUT http://localhost:5000/api/salidas/{id}

        [HttpPut]
        [Route("{id:int}")]
        [ResponseType(typeof(void))]
        public IHttpActionResult PutSalida(int id, Salidas salida)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            if (id != salida.Id)
            {
                return BadRequest();
            }

            _salidasService.UpdateSalida(salida);
            return StatusCode(HttpStatusCode.NoContent);
        }
        //DELETE http://localhost:5000/api/salidas/{id}


        [HttpDelete]
        [Route("{id:int}")]
        [ResponseType(typeof(Salidas))]
        public IHttpActionResult DeleteSalida(int id)
        {
            var salida = _salidasService.GetSalida(id);
            if (salida == null)
            {
                return NotFound();
            }

            _salidasService.DeleteSalida(id);
            return Ok(salida);
        }
    }
}

