include .env

MIGRATE := docker run --rm -v ./$(MIGRATION_PATH):/migrations -u $(ID_USER):$(ID_GROUP) --network go-sqlc_default migrate/migrate 

# Migrations
migration.create: 
	$(MIGRATE) create -ext sql -dir migrations $(NAME)
migration.down:
	$(MIGRATE) -path=migrations -database $(DB_URL) -verbose down -all

# SQLC
SQLC := docker run --rm -u $(ID_USER):$(ID_GROUP) -v .:/src -w /src sqlc/sqlc

gen.sqlc:
	$(SQLC) generate