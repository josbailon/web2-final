using System;
using Microsoft.Owin;
using Microsoft.Owin.Security.Jwt;
using Microsoft.Owin.Security;
using Owin;
using System.IdentityModel.Tokens.Jwt;
using Microsoft.IdentityModel.Tokens;
using System.Text;
using Microsoft.Owin.Security.DataHandler.Encoder;

public class Startup
{  //Esta clase configurará OWIN para usar JWT.
    public void Configuration(IAppBuilder app)
    {
        var issuer = "yourIssuer";
        var audience = "yourAudience";
        var secret = TextEncodings.Base64Url.Decode("yourBase64EncodedSecret");

        app.UseJwtBearerAuthentication(new JwtBearerAuthenticationOptions
        {
            AuthenticationMode = AuthenticationMode.Active,
            TokenValidationParameters = new TokenValidationParameters()
            {
                ValidIssuer = issuer,
                ValidAudience = audience,
                IssuerSigningKey = new SymmetricSecurityKey(secret)
            }
        });
    }
}













