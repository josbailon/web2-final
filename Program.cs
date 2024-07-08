using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.OpenApi.Models;
using PARQUEADERO2.Repositories;
using PARQUEADERO2.Services;
using PARQUEADERO2.ServicesInterfaces;
using PARQUEADERO2.Interfaces;
using Microsoft.AspNetCore.Mvc.NewtonsoftJson;
using PARQUEADERO2.Interfaces;
using PARQUEADERO2.Repositories;
using PARQUEADERO2.Services;
using PARQUEADERO2.ServicesInterfaces;

var builder = WebApplication.CreateBuilder(args);

// Configurar los puertos
builder.WebHost.UseUrls("http://localhost:5006");

// Configurar servicios
builder.Services.AddControllers()
    .AddNewtonsoftJson(options =>
        options.SerializerSettings.ReferenceLoopHandling = Newtonsoft.Json.ReferenceLoopHandling.Ignore);

// Agregar servicios de Swagger
builder.Services.AddSwaggerGen(c =>
{
    c.SwaggerDoc("v1", new OpenApiInfo { Title = "PARQUEADERO2", Version = "EVELYN" });
});

// Registro de servicios
builder.Services.AddScoped<ICarroService, CarroService>();
builder.Services.AddScoped<IClienteService, ClienteService>();
builder.Services.AddScoped<IIngresoService, IngresoService>();
builder.Services.AddScoped<ISalidaService, SalidaService>();
builder.Services.AddScoped<IPagoService, PagoService>();
builder.Services.AddScoped<ITarifaEstacionamientoService, TarifaEstacionamientoService>();

// Registro de repositorios
builder.Services.AddScoped<ICarro, CarroRepository>();
builder.Services.AddScoped<ICliente, ClienteRepository>();
builder.Services.AddScoped<IIngreso, IngresoRepository>();
builder.Services.AddScoped<ISalida, SalidaRepository>();
builder.Services.AddScoped<IPago, PagoRepository>();
builder.Services.AddScoped<ITarifaEstacionamiento, TarifaEstacionamientoRepository>();

try
{
    var app = builder.Build();

    if (app.Environment.IsDevelopment())
    {
        app.UseDeveloperExceptionPage();
    }
    else
    {
        app.UseExceptionHandler("/Home/Error");
        app.UseHsts();
    }

    app.UseHttpsRedirection();
    app.UseStaticFiles();

    app.UseRouting();

    app.UseAuthorization();

    app.UseSwagger();

    app.UseSwaggerUI(c =>
    {
        c.SwaggerEndpoint("/swagger/v1/swagger.json", "My API V1");
        c.RoutePrefix = string.Empty; // Para servir Swagger UI en la raíz (http://localhost:5006/)
    });

    app.MapControllers();

    app.Run();
}
catch (Exception ex)
{
    // Loggear la excepción
    Console.WriteLine($"Excepción no controlada: {ex}");
}



