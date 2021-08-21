.PHONY: run
COMPOSE=docker-compose
DOCKER=docker
GO=todo_backend
DB=techdojo_db
DB_URI="mysql://root:my-secret-pw@tcp(127.0.0.1:33061)/techdojo_db"
MIGRATE_FILES=mysql/migrations
v=

run:
	cd go && go run main.go

run/air:
	air

build:
	cd go && go build -o ../tmp/main main.go

docker/up:
	$(COMPOSE) up -d --build

docker/build:
	$(COMPOSE) build

# sh/go:
# 	$(DOCKER) exec -it $(GO) ash

sh/db:
	$(DOCKER) exec -it $(DB) bash

logs:
	$(COMPOSE) logs

ps:
	$(COMPOSE) ps

# restart/go:
# 	$(COMPOSE) restart $(GO)

restart/db:
	$(COMPOSE) restart $(DB)

down:
	$(COMPOSE) down -v

# down/go:
# 	$(COMPOSE) rm -fsv $(GO)

down/db:
	$(COMPOSE) rm -fsv $(DB)

migrate/up:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} up ${v}

migrate/down:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} down ${v}

migrate/force:
	migrate -path ${MIGRATE_FILES} -database ${DB_URI} force ${v}

migrate/updown: migrate/down migrate/up