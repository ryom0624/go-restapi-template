.PHONY: up down build logs ps test install-migrate migrate-up migrate-down migrate-new

up:
	docker-compose up -d

up-build:
	docker-compose up --build

down:
	docker-compose down -v

build:
	docker-compose build

logs:
	docker-compose logs -f

ps:
	docker-compose ps

test:
	go test -v ./...


install-migrate:
	@if ! [ -x "$$(command -v sql-migrate)" ]; then \
		go install github.com/rubenv/sql-migrate/...@latest; \
	fi

migrate-up: install-migrate
	@if [ -z "$(limit)" ]; then \
		sql-migrate up -config=migrations/dbconfig.yml -env=localhost-docker; \
	else \
		sql-migrate up -config=migrations/dbconfig.yml -env=localhost-docker -limit=$(limit); \
	fi

migrate-down: install-migrate
	@if [ -z "$(limit)" ]; then \
		sql-migrate down -config=migrations/dbconfig.yml -env=localhost-docker; \
	else \
		sql-migrate down -config=migrations/dbconfig.yml -env=localhost-docker -limit=$(limit); \
	fi

migrate-new: install-migrate
	sql-migrate new -config=migrations/dbconfig.yml -env=localhost-docker $(name)
