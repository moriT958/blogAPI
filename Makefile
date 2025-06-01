.PHONY: db-up
db-up:
	docker compose up -d

.PHONY: db-down
db-down:
	docker compose down

.PHONY: db-connect
db-connect:
	mysql --user=username --password --host 127.0.0.1 --port 3306

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migration-create name=[migration name]"; \
		exit 1; \
	fi
	goose -dir ./migrations create $(name) sql

.PHONY: migrate-up
migrate-up:
	goose -dir ./migrations mysql "username:password@tcp(127.0.0.1:3306)/blogdb" up

.PHONY: migrate-down
migrate-down:
	goose -dir ./migrations mysql "username:password@tcp(127.0.0.1:3306)/blogdb" down
