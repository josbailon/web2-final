using System.Collections.Generic;
using System.Net;
using System.Web.Http;
using System.Web.Http.Description;
using TA2PARQUEADERO.models;
using TA2PARQUEADERO.services;

namespace TA2PARQUEADERO.Controllers
{
    [RoutePrefix("api/Carros")]
    public class CarrosController : ApiController
    {
        private readonly CarrosService _carrosService;

        public CarrosController(CarrosService carrosService)
        {
            _carrosService = carrosService;
        }
        //GET http://localhost:5000/api/carros
        //Obtener todos los carros
        [HttpGet]
        [Route("")]
        public IEnumerable<Carros> GetCarros()
        {
            return _carrosService.GetCarros();
        }
        //GET http://localhost:5000/api/carros/{id}

        [HttpGet]
        [Route("{id:int}")]
        [ResponseType(typeof(Carros))]
        public IHttpActionResult GetCarro(int id)
        {
            var carro = _carrosService.GetCarro(id);
            if (carro == null)
            {
                return NotFound();
            }
            return Ok(carro);
        }
        //POST http://localhost:5000/api/carros

        [HttpPost]
        [Route("")]
        [ResponseType(typeof(Carros))]
        public IHttpActionResult PostCarro(Carros carro)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            _carrosService.CreateCarro(carro);
            return CreatedAtRoute("DefaultApi", new { id = carro.Id }, carro);
        }
        //PUT http://localhost:5000/api/carros/{id}

        [HttpPut]
        [Route("{id:int}")]
        [ResponseType(typeof(void))]
        public IHttpActionResult PutCarro(int id, Carros carro)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            if (id != carro.Id)
            {
                return BadRequest();
            }

            _carrosService.UpdateCarro(carro);
            return StatusCode(HttpStatusCode.NoContent);
        }
        //DELETE http://localhost:5000/api/carros/{id}

        [HttpDelete]
        [Route("{id:int}")]
        [ResponseType(typeof(Carros))]
        public IHttpActionResult DeleteCarro(int id)
        {
            var carro = _carrosService.GetCarro(id);
            if (carro == null)
            {
                return NotFound();
            }

            _carrosService.DeleteCarro(id);
            return Ok(carro);
        }
    }
}








