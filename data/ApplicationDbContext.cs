
using System.Data.Entity;
using TA2PARQUEADERO.models;

namespace TA2PARQUEADERO.data
{
    public class ApplicationDbContext : DbContext
    {
        public ApplicationDbContext() : base("Server=postgress;Port=tu_5432;Database=Evelyn;User Id=postgress;Password=tu;")
        {
        }

        public DbSet<Clientes> Clientes { get; set; }
        public DbSet<Carros> Carros { get; set; }
        public DbSet<Ingresos> Ingresos { get; set; }
        public DbSet<Salidas> Salidas { get; set; }
        public DbSet<Pagos> Pagos { get; set; }

        protected override void OnModelCreating(DbModelBuilder modelBuilder)
        {
            // Configuración de la relación entre Carros y Clientes
            modelBuilder.Entity<Carros>()
                .HasKey(c => c.Id);

            modelBuilder.Entity<Carros>()
                .HasRequired(c => c.Propietario)  
                .WithMany(c => c.Carros)         
                .HasForeignKey(c => c.PropietarioId);

            // Configuración de la relación entre Ingresos y Carros
            modelBuilder.Entity<Ingresos>()
                .HasKey(i => i.Id);

            modelBuilder.Entity<Ingresos>()
                .HasRequired(i => i.Vehiculo)    
                .HasForeignKey(i => i.VehiculoId);

            // Configuración de la relación entre Salidas e Ingresos
            modelBuilder.Entity<Salidas>()
                .HasKey(s => s.Id);

            modelBuilder.Entity<Salidas>()
                .HasRequired(s => s.Ingreso)     
                .WithOptional(i => i.Salida);    

            // Configuración de la relación entre Pagos e Ingresos
            modelBuilder.Entity<Pagos>()
                .HasKey(p => p.Id);

            modelBuilder.Entity<Pagos>()
                .HasRequired(p => p.Ingreso)     
                .WithMany(i => i.Pagos)          
                .HasForeignKey(p => p.IngresoId);

            base.OnModelCreating(modelBuilder);
        }
    }
}



