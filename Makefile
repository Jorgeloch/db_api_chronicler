include .env

build:
	go build -v -o bin/main main.go

create_migration:
	@read -p "Migration name: " NAME; \
	migrate create -ext sql -dir internal/database/migrations -seq $$NAME

migrate_up:
	migrate -path=internal/database/migrations -database "${DATABASE_URL}" -verbose up

migrate_down:
	migrate -path=internal/database/migrations -database "${DATABASE_URL}" -verbose down

.PHONY: create_migration migrate_up migrate_down
