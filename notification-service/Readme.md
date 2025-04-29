# Notification Service

Servicio encargado de enviar notificaciones por email cuando una compra es despachada o entregada.

## Funcionamiento

Escucha eventos de entrega en el bus de eventos (por ejemplo, NATS en memoria) y envía notificaciones correspondientes.

## Comandos útiles

```bash
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