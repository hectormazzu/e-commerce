# E-commerce Distribution Microservices

Este proyecto contiene tres microservicios escritos en Go que simulan el flujo de entrega de compras en una plataforma de ecommerce. Incluye:

- **`orders-service`**: Gestiona el estado e historial de las compras.
- **`delivery-service`**: Administra rutas de distribución y publica eventos.
- **`notification-service`**: Escucha eventos desde NATS y simula el envío de notificaciones.
- **`nats`**: Event bus en memoria para comunicación asíncrona entre microservicios.

## 🚀 Cómo ejecutar

### Requisitos

- Docker y Docker Compose instalados.

### Comando para iniciar

```bash
docker-compose up --build
```

Esto compilará y levantará los servicios en:

- **`orders-service`**: [http://localhost:8081](http://localhost:8081)
- **`delivery-service`**: [http://localhost:8082](http://localhost:8082)
- **`notification-service`**: Escucha eventos desde NATS.
- **`nats`**: Disponible en el puerto `4222`.

---

## 📦 Endpoints REST

### 🛒 `orders-service`

#### **GET /orders/{id}/status**
- Retorna el estado actual y el historial de una orden.
- **Ejemplo con cURL**:

```bash
curl -X GET http://localhost:8081/orders/order123/status
```

---

### 🚚 `delivery-service`

#### **POST /routes**
- Crea una nueva ruta.
- **Body**:

```json
{
  "id": "route1",
  "vehicle_id": "veh123",
  "driver_name": "Alice"
}
```

- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes -H "Content-Type: application/json" -d '{"id": "route1", "vehicle_id": "veh123", "driver_name": "Alice"}'
```

---

#### **POST /routes/{id}/orders**
- Agrega una orden a la ruta y publica un evento si el estado es `DISPATCHED` o `DELIVERED`.
- **Body**:

```json
{
  "order_id": "order123"
}
```

- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes/route1/orders -H "Content-Type: application/json" -d '{"order_id": "order123"}'
```

---

#### **POST /routes/{id}/start**
- Marca todas las órdenes de la ruta como `DISPATCHED` y publica eventos para cada orden.
- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes/route1/start
```

---

#### **POST /routes/{route_id}/orders/{order_id}/deliver**
- Marca una orden específica como `DELIVERED` y publica un evento.
- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes/route1/orders/order123/deliver
```

---

#### **GET /routes/{id}**
- Obtiene los detalles de la ruta y los estados de las órdenes asociadas.
- **Ejemplo con cURL**:

```bash
curl -X GET http://localhost:8082/routes/route1
```

---

## ✉️ Notificaciones

El servicio `notification-service` escucha el tópico `delivery.events` desde NATS y muestra en consola mensajes como:

```
[NOTIFICATION] Order order123 is now DELIVERED
```

---

## 🧪 Tests

Ejecuta los tests unitarios desde cada servicio. Por ejemplo:

```bash
docker-compose run orders-service go test ./...
docker-compose run delivery-service go test ./...
docker-compose run notification-service go test ./...
```

---

## 📂 Estructura de cada servicio

Cada microservicio sigue una estructura modular:

```
<service-name>/
├── cmd/          # Punto de entrada principal (main.go)
├── internal/     # Lógica de negocio, handlers, servicios y repositorios
├── test/         # Tests unitarios
└── Dockerfile    # Configuración para contenedores
```

---

## 🌟 Características principales

1. **Arquitectura basada en eventos**:
   - Los microservicios se comunican de forma asíncrona a través de NATS.
   - `delivery-service` publica eventos en el tópico `delivery.events`.
   - `notification-service` escucha estos eventos y simula el envío de notificaciones.

2. **Modularidad**:
   - Cada servicio tiene capas separadas para handlers, lógica de negocio y acceso a datos.

3. **Pruebas unitarias**:
   - Los tests aseguran la calidad del código y cubren casos de éxito y error.

4. **Escalabilidad**:
   - Los servicios son independientes y pueden escalarse horizontalmente.

---

## 🛠️ Tecnologías utilizadas

- **Go**: Lenguaje principal para los microservicios.
- **Docker**: Contenedores para cada servicio.
- **Docker Compose**: Orquestación de los servicios.
- **NATS**: Event bus para comunicación asíncrona.
- **Gin**: Framework HTTP para manejar rutas y peticiones.

---
