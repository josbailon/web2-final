// src/pusher.ts

import Pusher from 'pusher-js';

const pusher = new Pusher('1047aef918d6aeaac37a', {
  cluster: 'us2',
  // Eliminar la propiedad `encrypted`
  // Pusher.js automáticamente usa SSL si la URL es `https`
});

export default pusher;
