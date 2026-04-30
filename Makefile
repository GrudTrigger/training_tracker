SHELL := /bin/bash

ROOT_DIR := $(abspath .)
BACKEND_DIR := $(ROOT_DIR)/backend
FRONTEND_DIR := $(ROOT_DIR)/frontend
BIN_DIR := $(ROOT_DIR)/bin

GO ?= go
NPM ?= npm

POSTGRES_DSN ?= postgres://postgres:postgres@localhost:5432/training_tracker?sslmode=disable
GOOSE_DRIVER ?= postgres
GOOSE_MIGRATIONS_DIR ?= $(BACKEND_DIR)/migrations
GOOSE_BIN := $(BIN_DIR)/goose
GOOSE_VERSION ?= v3.26.0
GOOSE_BUILD_TAGS ?= no_clickhouse no_libsql no_mssql no_mysql no_sqlite3 no_turso no_vertica no_ydb

export GOCACHE := $(BACKEND_DIR)/.cache/go-build
export GOMODCACHE := $(BACKEND_DIR)/.cache/go-mod

.PHONY: dev dev-backend dev-frontend \
	backend-run frontend-run \
	backend-test backend-tidy frontend-install frontend-build \
	infra-up infra-down goose-install goose-up goose-down goose-status goose-reset goose-create

dev:
	@trap 'kill 0' INT TERM EXIT; \
	$(MAKE) dev-backend & \
	$(MAKE) dev-frontend & \
	wait

dev-backend:
	cd $(BACKEND_DIR) && $(GO) run ./cmd/training-tracker

dev-frontend:
	cd $(FRONTEND_DIR) && $(NPM) run dev -- --host 0.0.0.0 --port 3000

backend-run: dev-backend

frontend-run: dev-frontend

backend-test:
	cd $(BACKEND_DIR) && $(GO) test ./...

backend-tidy:
	cd $(BACKEND_DIR) && $(GO) mod tidy

frontend-install:
	cd $(FRONTEND_DIR) && $(NPM) install

frontend-build:
	cd $(FRONTEND_DIR) && $(NPM) run build

infra-up:
	docker compose up -d postgres redis

infra-down:
	docker compose down

goose-install:
	@mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) $(GO) install -tags '$(GOOSE_BUILD_TAGS)' github.com/pressly/goose/v3/cmd/goose@$(GOOSE_VERSION)

goose-up: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(GOOSE_MIGRATIONS_DIR) $(GOOSE_DRIVER) $(POSTGRES_DSN) up

goose-down: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(GOOSE_MIGRATIONS_DIR) $(GOOSE_DRIVER) $(POSTGRES_DSN) down

goose-status: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(GOOSE_MIGRATIONS_DIR) $(GOOSE_DRIVER) $(POSTGRES_DSN) status

goose-reset: $(GOOSE_BIN)
	$(GOOSE_BIN) -dir $(GOOSE_MIGRATIONS_DIR) $(GOOSE_DRIVER) $(POSTGRES_DSN) reset

goose-create: $(GOOSE_BIN)
	@if [ -z "$(name)" ]; then \
		echo 'Usage: make goose-create name=create_users_table'; \
		exit 1; \
	fi
	$(GOOSE_BIN) -dir $(GOOSE_MIGRATIONS_DIR) create $(name) sql

$(GOOSE_BIN): goose-install
