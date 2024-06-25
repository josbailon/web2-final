// web/graphql/graphql.js

// Ejemplo básico de consulta GraphQL desde el cliente
const query = `
    query {
        pago(id: 1) {
            id
            monto
            descripcion
            estado
        }
    }
`;

fetch('/graphql', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
})
.then(response => response.json())
.then(data => {
    console.log('Respuesta GraphQL:', data);
})
.catch(error => {
    console.error('Error al realizar la consulta GraphQL:', error);
});
