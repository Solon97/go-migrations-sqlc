MIGRATION_PATH := sql/migrations
MIGRATE := podman run -v ./$(MIGRATION_PATH):/migrations migrate/migrate

migration.create: 
	$(MIGRATE) create -ext sql -dir migrations $(NAME)
migration.up:
	$(MIGRATE) -path=migrations -database $(DATABASE) -verbose up
migration.down:
	$(MIGRATE) -path=migrations -database $(DATABASE) -verbose down -all