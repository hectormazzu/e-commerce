# Makefile

# Nombre del stack de docker-compose
COMPOSE=docker-compose

# Comandos principales
up:
	$(COMPOSE) up --build

down:
	$(COMPOSE) down

restart:
	$(MAKE) down
	$(MAKE) up

logs:
	$(COMPOSE) logs -f

build:
	$(COMPOSE) build

ps:
	$(COMPOSE) ps

# Comando para eliminar todos los volúmenes (opcional)
clean:
	$(COMPOSE) down -v --remove-orphans

# Atajo para entrar a un servicio, ejemplo: make sh service=delivery-service
sh:
	docker exec -it $$(docker ps -qf "name=$$service") /bin/sh
