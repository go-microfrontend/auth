build:
	docker-compose build

up-db:
	docker-compose up -d db

up-app:
	docker-compose up -d app

up: up-db up-app

down:
	docker-compose down

restart: down up

db-shell:
	docker-compose exec db psql -U postgres -d postgres

app-shell:
	docker-compose exec app sh

goose-up:
	goose \
		-dir ./db/migrations/ \
		postgres \
		"postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" \
		up

goose-down:
	goose \
		-dir ./db/migrations/ \
		postgres \
		"postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" \
		down
