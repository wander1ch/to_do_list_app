include .env
export


export PROJECT_ROOT = $(shell pwd)

env-up:
	@docker compose up -d todolist-postgres

env-down:
	@docker compose down todolist-postgres
 
env-clean:
	@read -p "Are you sure you want to remove the environment? This action cannot be undone. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todolist-postgres && \
		rm -rf ./out/pgdata && \
		echo "Environment removed."; \
	else \
		echo "Action cancelled."; \
	fi

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Please provide a migration name using 'make migrate-create name=your_migration_name'"; \
		exit 1; \
	fi
	docker compose run --rm todolist-postgres-migrate create \
	-ext sql \
	-dir /migrations \
	-seq "$(name)"

migrate-up:
	@make migrate-action action=up
	
migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Error: Please provide an action (up or down) using 'make migrate-action action=up'"; \
		exit 1; \
	fi	
	docker compose run --rm todolist-postgres-migrate \
	-path=/migrations \
	-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}\
	@todolist-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
	"$(action)"

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder