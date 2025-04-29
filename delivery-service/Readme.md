# Delivery Service

Servicio encargado de gestionar rutas de distribución y publicar eventos relacionados con el estado de las órdenes en una plataforma de e-commerce.

## 🚀 Funcionamiento

El `delivery-service` permite:

- Crear rutas de distribución.
- Agregar órdenes a rutas existentes.
- Iniciar rutas, marcando las órdenes como `DISPATCHED`.
- Marcar órdenes específicas como `DELIVERED`.
- Publicar eventos en el bus de eventos (NATS) para sincronización con otros servicios.

---

## 📦 Comandos útiles

### Ejecutar localmente

```bash
make run
```

### Ejecutar pruebas

```bash
make test
```

### Compilar binario

```bash
make build
```

### Ejecutar en contenedor Docker

```bash
make docker-build
make docker-run
```

---

## 📦 Endpoints REST

### **POST /routes**
- Crea una nueva ruta.
- **Body**:

```json
{
  "vehicle_id": "veh123",
  "driver_id": "driver123"
}
```

- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes -H "Content-Type: application/json" -d '{"vehicle_id": "veh123", "driver_id": "driver123"}'
```

---

### **POST /routes/{id}/orders**
- Agrega una orden a una ruta existente.
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

### **POST /routes/{id}/start**
- Inicia una ruta, marcando todas las órdenes como `DISPATCHED` y publicando eventos.
- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes/route1/start
```

---

### **POST /routes/{route_id}/orders/{order_id}/deliver**
- Marca una orden específica como `DELIVERED` y publica un evento.
- **Ejemplo con cURL**:

```bash
curl -X POST http://localhost:8082/routes/route1/orders/order123/deliver
```

---

### **GET /routes/{id}**
- Obtiene los detalles de una ruta y el estado de las órdenes asociadas.
- **Ejemplo con cURL**:

```bash
curl -X GET http://localhost:8082/routes/route1
```

---

## 🧪 Tests

Ejecuta los tests unitarios para asegurar la calidad del código:

```bash
make test
```

---

## 📂 Estructura del proyecto

```
delivery-service/
├── cmd/          # Punto de entrada principal (main.go)
├── internal/     # Lógica de negocio, handlers, servicios y repositorios
├── test/         # Tests unitarios
├── Dockerfile    # Configuración para contenedores
└── Makefile      # Comandos útiles para desarrollo
```

---

## 🌟 Características principales

1. **Gestión de rutas**:
   - Crear, iniciar y gestionar rutas de distribución.

2. **Publicación de eventos**:
   - Publica eventos en el tópico `delivery.events` para sincronización con otros servicios.

3. **Arquitectura modular**:
   - Separación clara entre lógica de negocio, handlers y acceso a datos.

4. **Pruebas unitarias**:
   - Cobertura de casos de éxito y error para garantizar la calidad del servicio.

5. **Escalabilidad**:
   - Diseñado para manejar grandes volúmenes de órdenes y rutas.

---

## 🛠️ Tecnologías utilizadas

- **Go**: Lenguaje principal para el servicio.
- **Docker**: Contenedor para el despliegue.
- **NATS**: Event bus para comunicación asíncrona.
- **Gin**: Framework HTTP para manejar rutas y peticiones.
- **GORM**: ORM para la interacción con la base de datos.

---

## 🌐 Integración con otros servicios

El `delivery-service` es parte de un sistema de microservicios que incluye:

- **`orders-service`**: Proporciona información sobre el estado de las órdenes.
- **`notification-service`**: Escucha eventos de `delivery.events` y envía notificaciones.

---
