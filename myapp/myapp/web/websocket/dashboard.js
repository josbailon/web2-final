// web/websocket/dashboard.js

// Ejemplo básico de configuración de WebSocket para el dashboard
const socket = new WebSocket('ws://localhost:8080/ws/dashboard');

socket.addEventListener('open', function (event) {
    console.log('Conexión establecida con el servidor WebSocket');
});

socket.addEventListener('message', function (event) {
    console.log('Mensaje recibido del servidor WebSocket:', event.data);
});

socket.addEventListener('close', function (event) {
    console.log('Conexión cerrada con el servidor WebSocket');
});
