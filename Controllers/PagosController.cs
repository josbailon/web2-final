
using System.Collections.Generic;
using System.Net;
using System.Web.Http;
using System.Web.Http.Description;
using TA2PARQUEADERO.models;
using TA2PARQUEADERO.services;

namespace TA2PARQUEADERO.Controllers
{
    [RoutePrefix("api/Pagos")]
    public class PagosController : ApiController
    {
        private readonly PagosService _pagosService;

        public PagosController(PagosService pagosService)
        {
            _pagosService = pagosService;
        }
        //GET http://localhost:5000/api/pagos

        [HttpGet]
        [Route("")]
        public IEnumerable<Pagos> GetPagos()
        {
            return _pagosService.GetPagos();
        }
        //GET http://localhost:5000/api/pagos/{id}

        [HttpGet]
        [Route("{id:int}")]
        [ResponseType(typeof(Pagos))]
        public IHttpActionResult GetPago(int id)
        {
            var pago = _pagosService.GetPago(id);
            if (pago == null)
            {
                return NotFound();
            }
            return Ok(pago);
        }
        //POST http://localhost:5000/api/pagos

        [HttpPost]
        [Route("")]
        [ResponseType(typeof(Pagos))]
        public IHttpActionResult PostPago(Pagos pago)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            _pagosService.CreatePago(pago);
            return CreatedAtRoute("DefaultApi", new { id = pago.Id }, pago);
        }
        //PUT http://localhost:5000/api/pagos/{id}

        [HttpPut]
        [Route("{id:int}")]
        [ResponseType(typeof(void))]
        public IHttpActionResult PutPago(int id, Pagos pago)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            if (id != pago.Id)
            {
                return BadRequest();
            }

            _pagosService.UpdatePago(pago);
            return StatusCode(HttpStatusCode.NoContent);
        }
        //DELETE http://localhost:5000/api/pagos/{id}

        [HttpDelete]
        [Route("{id:int}")]
        [ResponseType(typeof(Pagos))]
        public IHttpActionResult DeletePago(int id)
        {
            var pago = _pagosService.GetPago(id);
            if (pago == null)
            {
                return NotFound();
            }

            _pagosService.DeletePago(id);
            return Ok(pago);
        }
    }
}
