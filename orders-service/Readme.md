# Orders Service

Servicio de gestión de órdenes de compra para la plataforma de e-commerce.

## Endpoints

- `GET /orders/{id}/status`  
  Obtiene el estado actual y el historial de una orden de compra.

### Ejemplos de uso (cURL)

```bash
# Obtener estado de una orden
curl -X GET http://localhost:8081/orders/{order_id}/status

# Correr localmente
make run

# Correr tests
make test

# Compilar binario
make build

# Correr en contenedor
make docker-build
make docker-run

# Hot reload
make air