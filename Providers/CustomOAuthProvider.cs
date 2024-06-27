using Microsoft.Owin.Security;
using Microsoft.Owin.Security.OAuth;
using System.Security.Claims;
using System.Threading.Tasks;

namespace TA2PARQUEADERO.Providers
{
    public class CustomOAuthProvider : OAuthAuthorizationServerProvider
    {
        public override Task ValidateClientAuthentication(OAuthValidateClientAuthenticationContext context)
        {
            // Validación básica del cliente (ejemplo: permitir todos los clientes)
            context.Validated();
            return Task.CompletedTask;
        }

        public override async Task GrantResourceOwnerCredentials(OAuthGrantResourceOwnerCredentialsContext context)
        {
            // Aquí debes implementar la lógica para validar las credenciales del usuario y generar el token JWT
            // Ejemplo básico: validar el nombre de usuario y contraseña
            if (context.UserName != "Evelyn" || context.Password != "Evelyn123")
            {
                context.SetError("invalid_grant", "Nombre de usuario o contraseña incorrectos.");
                return;
            }

            var identity = new ClaimsIdentity(context.Options.AuthenticationType);
            identity.AddClaim(new Claim(ClaimTypes.Name, context.UserName));

            var ticket = new AuthenticationTicket(identity, null);
            context.Validated(ticket);
        }
    }
}

