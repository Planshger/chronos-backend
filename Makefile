include .env 
export 

export PROJECT_ROOT=$(shell pwd)

docker-up:
	@docker compose up -d chronos-postgres

docker-down: 
	@docker compose down chronos-postgres

port-forward-up:
	@docker compose up -d port-forwarder

port-forward-down:
	@docker compose down port-forwarder

cleanup-db: 
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down chronos-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

migrate-create: 
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый праметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm chronos-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
 
migrate-up:
	@make migrate-action action=up

migrate-down: 
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый праметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm chronos-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@chronos-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"