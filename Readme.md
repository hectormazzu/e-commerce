E-commerce Distribution Microservices

Este proyecto contiene tres microservicios escritos en Go que simulan el flujo de entrega de compras en una plataforma de ecommerce. Incluye:

orders-service: Gestiona el estado e historial de las compras.

delivery-service: Administra rutas de distribución y publica eventos.

notification-service: Escucha eventos y simula el envío de notificaciones.

nats: Event bus en memoria para comunicación asíncrona.

🚀 Cómo ejecutar

Requisitos

Docker y Docker Compose instalados

Comando

docker-compose up --build

Esto compilará y levantará los servicios en:

orders-service: http://localhost:8081

delivery-service: http://localhost:8082

notification-service: escucha eventos de NATS

nats: puerto 4222

📦 Endpoints REST

🛒 orders-service

GET /orders/:id/status

Retorna el estado actual y el historial de una orden.

🚚 delivery-service

POST /routes

Crea una nueva ruta

Body:

{
  "id": "route1",
  "vehicle_id": "veh123",
  "driver_name": "Alice"
}

POST /routes/:id/orders

Agrega una orden a la ruta y notifica si el estado es DISPATCHED o DELIVERED

Body:

{
  "order_id": "order123"
}

GET /routes/****:id

Obtiene los detalles de la ruta y estados de órdenes asociadas

✉️ Notificaciones

notification-service escucha el tópico delivery.events desde NATS y muestra en consola:

[NOTIFICATION] Order order123 is now DELIVERED

🧪 Tests

Ejecutá tests unitarios desde cada servicio, por ejemplo:

docker-compose run orders-service go test ./...

📂 Estructura de cada servicio

orders-service/
├── cmd/          # main
├── internal/     # lógica de negocio
└── test/         # tests

Este proyecto está preparado para ejecutarse localmente y sirve como ejemplo educativo de microservicios conectados con event-driven architecture en Go.

