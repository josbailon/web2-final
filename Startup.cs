using Microsoft.Owin;
using Microsoft.Owin.Security;
using Microsoft.Owin.Security.Jwt;
using Owin;
using System;
using System.Text;
using System.Web.Http;

[assembly: OwinStartup(typeof(TA2PARQUEADERO.Startup))]

namespace TA2PARQUEADERO
{
    public class Startup
    {
        public void Configuration(IAppBuilder app)
        {
            HttpConfiguration config = new HttpConfiguration();

            config.MapHttpAttributeRoutes();

            ConfigureOAuth(app);

            config.EnableCors();

            app.UseWebApi(config);
        }

        public void ConfigureOAuth(IAppBuilder app)
        {
            var issuer = "Evelyn"; // Reemplaza con tu emisor (issuer)
            var audience = "Evelyn"; // Reemplaza con tu audiencia (audience)
            var secretKey = "Evelyn123"; // Reemplaza con tu clave secreta

            var tokenValidationParameters = new Microsoft.IdentityModel.Tokens.TokenValidationParameters
            {
                ValidateIssuer = true,
                ValidateAudience = true,
                ValidateLifetime = true,
                ValidateIssuerSigningKey = true,
                ValidIssuer = issuer,
                ValidAudience = audience,
                IssuerSigningKey = new Microsoft.IdentityModel.Tokens.SymmetricSecurityKey(Encoding.UTF8.GetBytes(secretKey)),
                ClockSkew = TimeSpan.Zero // Opcional: ajusta la tolerancia de tiempo para la expiración del token
            };

            // Configuración de JWT Bearer Authentication
            var jwtOptions = new JwtBearerAuthenticationOptions
            {
                AuthenticationMode = Microsoft.Owin.Security.AuthenticationMode.Active,
                TokenValidationParameters = tokenValidationParameters
            };

            app.UseJwtBearerAuthentication(jwtOptions);
        }
    }
}













