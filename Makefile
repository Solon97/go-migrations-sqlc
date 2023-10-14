include .env
ID_USER := $(shell id -u)
ID_GROUP := $(shell id -g)

db.up:
	docker run --rm  \
	--name sqlc-postgres \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=root \
	-e POSTGRES_DB=courses \
	-v ./db/pg_data:/var/lib/postgresql/data \
	-p 5432:5432 \
	-u $(ID_USER):$(ID_GROUP) \
	--network go-sqlc_default \
	postgres:latest

MIGRATION_PATH := sql/migrations
MIGRATE := docker run --rm -v ./$(MIGRATION_PATH):/migrations -u $(ID_USER):$(ID_GROUP) --network go-sqlc_default migrate/migrate 

# Migrations
migration.create: 
	$(MIGRATE) create -ext sql -dir migrations $(NAME)
migration.up:
	$(MIGRATE) -path=migrations -database $(DATABASE_URL) -verbose up
migration.down:
	$(MIGRATE) -path=migrations -database $(DATABASE_URL) -verbose down -all

# SQLC
SQLC := docker run --rm -u $(ID_USER):$(ID_GROUP) -v .:/src -w /src sqlc/sqlc

gen.sqlc:
	$(SQLC) generate