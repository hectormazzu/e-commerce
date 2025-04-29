# Orders Service

Servicio de gestión de órdenes de compra para la plataforma de e-commerce.

## 🚀 Cómo ejecutar

### Requisitos

- Docker y Docker Compose instalados.
- Go 1.20 o superior (si deseas ejecutarlo localmente).

### Comandos para iniciar

#### Ejecutar con Docker Compose

```bash
docker-compose up --build
```

Esto levantará el servicio en [http://localhost:8081](http://localhost:8081).

#### Ejecutar localmente

```bash
make run
```

#### Ejecutar pruebas

```bash
make test
```

#### Compilar binario

```bash
make build
```

#### Ejecutar en contenedor Docker

```bash
make docker-build
make docker-run
```

---

## 📦 Endpoints REST

### **GET /orders/{id}/status**
- Obtiene el estado actual y el historial de una orden de compra.
- **Ejemplo con cURL**:

```bash
curl -X GET http://localhost:8081/orders/{order_id}/status
```

---

## 📂 Estructura del proyecto

```
orders-service/
├── cmd/          # Punto de entrada principal (main.go)
├── internal/     # Lógica de negocio, handlers, servicios y repositorios
├── test/         # Tests unitarios
├── Dockerfile    # Configuración para contenedores
└── Makefile      # Comandos útiles para desarrollo
```

---

## 🌟 Características principales

1. **Gestión de órdenes**:
   - Consulta el estado actual y el historial de una orden.

2. **Arquitectura modular**:
   - Separación clara entre handlers, lógica de negocio y acceso a datos.

3. **Pruebas unitarias**:
   - Cobertura de casos de éxito y error para garantizar la calidad del servicio.

4. **Integración con NATS**:
   - Publica y escucha eventos para sincronización con otros servicios.

---

## 🛠️ Tecnologías utilizadas

- **Go**: Lenguaje principal para el servicio.
- **Docker**: Contenedor para el despliegue.
- **Gin**: Framework HTTP para manejar rutas y peticiones.
- **NATS**: Event bus para comunicación asíncrona.
- **GORM**: ORM para la interacción con la base de datos.

---