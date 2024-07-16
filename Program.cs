using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.OpenApi.Models;
using Microsoft.EntityFrameworkCore;
using WEBFINAL2.IRepositories;
using WEBFINAL2.Models;
using WEBFINAL2.Repositories;
using WEBFINAL2.Services;

var builder = WebApplication.CreateBuilder(args);

// Configurar los puertos
builder.WebHost.UseUrls("http://localhost:5006");

// Configuración del DbContext para PostgreSQL
builder.Services.AddDbContext<YourDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")));

// Agregar servicios al contenedor
builder.Services.AddScoped<IRepository<Pago>, PagoRepository>();
builder.Services.AddScoped<IRepository<TarifaEstacionamiento>, TarifaEstacionamientoRepository>();
builder.Services.AddScoped<IRepository<Cliente>, ClienteRepository>();

builder.Services.AddScoped<IPagoService, PagoService>();
builder.Services.AddScoped<ITarifaEstacionamientoService, TarifaEstacionamientoService>();

// Agregar Swagger/OpenAPI
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen(c =>
{
    c.SwaggerDoc("v1", new OpenApiInfo { Title = "PARQUEADEROEVE", Version = "v1" });
});

var app = builder.Build();

// Configurar el pipeline de solicitudes HTTP
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI(c =>
    {
        c.SwaggerEndpoint("/swagger/v1/swagger.json", "PARQUEADEROEVE v1");
        c.RoutePrefix = string.Empty; // Para servir Swagger UI en la raíz (http://localhost:5006/)
    });
}

app.UseHttpsRedirection();
app.UseAuthorization();

app.MapControllers();

app.Run();


