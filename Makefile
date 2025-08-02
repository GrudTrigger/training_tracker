include .env

MIGRATIONS_DIR = migrations
MIGRATE_BIN = $(HOME)/go/bin/migrate
MIGRATE_CREATE = $(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_DIR) -seq

.PHONY: migrate

migrate:
	@if [ -z "$(NAME)" ]; then \
		echo "❌ Укажи имя миграции: make migrate NAME=add_users_table"; \
		exit 1; \
	fi; \
	echo "🚀 Создание миграции: $(NAME)"; \
	$(MIGRATE_CREATE) $(NAME)

migration_up: 
	migrate -path migrations -database $(DATABASE_URL) -verbose up

migration_down: 
	migrate -path migrations -database $(DATABASE_URL) -verbose down

migration_fix: 
	migrate -path migrations -database $(DATABASE_URL) force VERSION